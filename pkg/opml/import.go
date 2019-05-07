package opml

import (
	"fmt"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func newItemError(index int, err error) error {
	return fmt.Errorf("line_%d:%s", index, err.Error())
}

// ImportOPMLToDB imports OPML object into the database.
func ImportOPMLToDB(opml *OPML, db store.DB) error {
	logger := log.With().Str("component", "import").Str("title", opml.Head.Title).Logger()
	logger.Debug().Msg("importing OPML")
	result := importOutlines(db, logger, opml.Body.Outlines, "")
	if result.ErrorOrNil() != nil {
		logger.Info().Int("errors", len(result.Errors)).Msg("OPML file imported with errors")
	} else {
		logger.Info().Msg("OPML file imported")
	}
	return result.ErrorOrNil()
}

func importOutlines(db store.DB, logger zerolog.Logger, outlines []Outline, category string) *multierror.Error {
	var result *multierror.Error
	for idx, outline := range outlines {
		if len(outline.Outlines) > 0 {
			cat := outline.Title
			if category != "" {
				cat += "," + category
			}
			importOutlines(db, logger, outline.Outlines, cat)
			continue
		}
		if db.ExistsFeed(outline.XMLURL) {
			logger.Debug().Str("url", outline.XMLURL).Msg("feed already exists: skipped")
			continue
		}
		if category == "" {
			category = outline.Category
		}
		logger.Debug().Str("url", outline.XMLURL).Msg("importing feed")
		feed, err := builder.NewFeed(outline.XMLURL, &category)
		if err != nil {
			logger.Warn().Err(err).Str("url", outline.XMLURL).Msg("unable to create feed: skipped")
			result = multierror.Append(result, newItemError(idx, err))
			continue
		}
		if len(strings.TrimSpace(outline.Title)) > 0 {
			feed.Title = outline.Title
		}
		// TODO register new feed aggregators
		err = db.SaveFeed(feed)
		if err != nil {
			logger.Warn().Err(err).Msg("unable to save feed: skipped")
			result = multierror.Append(result, newItemError(idx, err))
		}
		logger.Info().Str("title", feed.Title).Str("url", feed.XMLURL).Msg("feed imported")
	}
	return result
}
