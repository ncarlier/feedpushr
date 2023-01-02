package opml

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/feed"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var jobID uint32

// ImportJobItemResult is the OPML outline import result
type ImportJobItemResult struct {
	url string
	err error
}

// ImportJob is a job that import OPML
type ImportJob struct {
	ID          uint32
	opml        *OPML
	outputFile  *os.File
	wOutputFile *bufio.Writer
	workload    sync.WaitGroup
	db          store.DB
	logger      zerolog.Logger
}

func newOPMLImportJob(db store.DB, _opml *OPML) (*ImportJob, error) {
	id := atomic.AddUint32(&jobID, 1)
	outputFilename := path.Join(os.TempDir(), fmt.Sprintf("feedpushr_import_%d_%d.txt", id, time.Now().Unix()))
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return nil, err
	}

	return &ImportJob{
		ID:          id,
		opml:        _opml,
		outputFile:  outputFile,
		wOutputFile: bufio.NewWriter(outputFile),
		db:          db,
		logger:      log.With().Uint32("import-job", id).Logger(),
	}, nil
}

// start import job
func (job *ImportJob) start() {
	job.logger.Debug().Str("title", job.opml.Head.Title).Msg("importing OPML")
	job.workload.Add(1)
	go job.importOutlines(job.opml.Body.Outlines, "")
}

// Wait for job complete
// Returns true if the wait completed without timing out, false otherwise.
func (job *ImportJob) Wait(timeout time.Duration) bool {
	ch := make(chan struct{})
	go func() {
		job.workload.Wait()
		close(ch)
	}()
	select {
	case <-ch:
		return true
	case <-time.After(timeout):
		return false
	}
}

func (job *ImportJob) writeResult(XMLURL string, err error) {
	status := "ok"
	if err != nil {
		status = err.Error()
	}
	job.wOutputFile.WriteString(fmt.Sprintf("%s|%s\n", XMLURL, status))
	job.wOutputFile.Flush()
}

func (job *ImportJob) importOutlines(outlines []Outline, category string) {
	for _, outline := range outlines {
		if len(outline.Outlines) > 0 {
			job.workload.Add(1)
			job.importOutlines(outline.Outlines, feed.JoinTags(category, outline.Title))
			continue
		}
		logger := job.logger.With().Str("url", outline.XMLURL).Logger()
		if job.db.ExistsFeed(outline.XMLURL) {
			logger.Debug().Msg("feed already exists: skipped")
			job.writeResult(outline.XMLURL, common.ErrFeedAlreadyExists)
			continue
		}
		cat := feed.JoinTags(category, outline.Category)
		logger.Debug().Msg("importing feed")
		_feed, err := feed.NewFeed(outline.XMLURL, &cat)
		if err != nil {
			logger.Warn().Err(err).Msg("unable to create feed: skipped")
			job.writeResult(outline.XMLURL, err)
			continue
		}
		if len(strings.TrimSpace(outline.Title)) > 0 {
			_feed.Title = outline.Title
		}
		// TODO register new feed aggregators
		err = job.db.SaveFeed(_feed)
		if err != nil {
			logger.Warn().Err(err).Msg("unable to save feed: skipped")
			job.writeResult(outline.XMLURL, err)
			continue
		}
		job.writeResult(outline.XMLURL, nil)
		logger.Info().Str("title", _feed.Title).Msg("feed imported")
	}
	if category == "" {
		job.wOutputFile.WriteString("done\n")
		job.wOutputFile.Flush()
		job.outputFile.Close()
	}
	job.workload.Done()
}
