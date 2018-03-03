package opml

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/ncarlier/feedpushr/pkg/builder"
	"github.com/ncarlier/feedpushr/pkg/store"
	"github.com/rs/zerolog/log"
)

func newItemError(index int, err error) error {
	return fmt.Errorf("line_%d:%s", index, err.Error())
}

// ImportOPMLToDB imports OPML object into the database.
func ImportOPMLToDB(opml *OPML, db store.DB) error {
	var result *multierror.Error
	logger := log.With().Str("component", "import").Str("title", opml.Head.Title).Logger()
	logger.Debug().Msg("importing OPML")
	for idx, outline := range opml.Body.Outlines {
		feed, err := builder.NewFeed(outline.XMLURL)
		if err != nil {
			logger.Warn().Err(err).Msg("unable to create feed")
			result = multierror.Append(result, newItemError(idx, err))
			continue
		}
		// TODO register new feed aggregators
		err = db.SaveFeed(feed)
		if err != nil {
			logger.Warn().Err(err).Msg("unable to save feed")
			result = multierror.Append(result, newItemError(idx, err))
		}
	}
	if result.ErrorOrNil() != nil {
		logger.Debug().Int("errors", len(result.Errors)).Msg("OPML imported with errors")
	} else {
		logger.Debug().Msg("OPML imported")
	}
	return result.ErrorOrNil()
}
