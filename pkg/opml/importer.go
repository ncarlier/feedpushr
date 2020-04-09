package opml

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/ncarlier/feedpushr/v2/pkg/helper"
	"github.com/ncarlier/feedpushr/v2/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Importer is used to trace OPML import jobs in background
type Importer struct {
	db     store.DB
	logger zerolog.Logger
}

// NewOPMLImporter creates new OPML importer
func NewOPMLImporter(db store.DB) *Importer {
	return &Importer{
		db:     db,
		logger: log.With().Str("component", "importer").Logger(),
	}
}

// ImportOPMLFile imports OPML file into the DB
func (importer *Importer) ImportOPMLFile(filename string) (*ImportJob, error) {
	importer.logger.Debug().Str("filename", filename).Msg("importing OPML file...")
	o, err := NewOPMLFromFile(filename)
	if err != nil {
		return nil, err
	}
	return importer.ImportOPML(o)
}

// ImportOPML imports OPML object into the DB
func (importer *Importer) ImportOPML(opml *OPML) (*ImportJob, error) {
	job, err := newOPMLImportJob(importer.db, opml)
	if err != nil {
		return nil, err
	}
	job.start()
	return job, nil
}

// Get import job
func (importer *Importer) Get(jobID uint64) (chan string, error) {
	logPattern := path.Join(os.TempDir(), fmt.Sprintf("feedpushr_import_%d_*.txt", jobID))
	files, err := filepath.Glob(logPattern)
	if err != nil {
		return nil, err
	}
	if len(files) > 0 {
		filename := files[len(files)-1]
		return helper.TailFile(filename, "done")
	}
	return nil, errors.New("job not found")
}
