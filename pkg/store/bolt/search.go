package store

import (
	"fmt"
	"strings"

	"github.com/blevesearch/bleve"
	bolt "github.com/coreos/bbolt"
	"github.com/getlantern/errors"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

func openSearchIndex(dbPath string) (bleve.Index, error) {
	idxPath := strings.TrimSuffix(dbPath, "db") + "idx"
	index, err := bleve.Open(idxPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		// TODO create tailor made mapping
		mapping := bleve.NewIndexMapping()
		return bleve.New(idxPath, mapping)
	} else if err != nil {
		return nil, err
	}
	return index, nil
}

func initSearchIndex(db *bolt.DB) (bleve.Index, error) {
	// Open search index
	index, err := openSearchIndex(db.Path())
	if err != nil {
		return nil, fmt.Errorf("unable to open search index, %v", err)
	}
	return index, nil
}

// BuildInitialIndex create initial index (only if empty)
func (store *BoltStore) BuildInitialIndex() error {
	nb, err := store.index.DocCount()
	if err != nil {
		return fmt.Errorf("unable to initialize index: %w", err)
	}
	if nb == 0 {
		return store.ForEachFeed(func(f *model.FeedDef) error {
			if f == nil {
				return errors.New("unable to index feed: feed is null")
			}
			// TODO batch indexing: index.NewBatch()
			return store.index.Index(f.ID, f)
		})
	}
	return nil
}

// SearchFeeds search feeds using search index
func (store *BoltStore) SearchFeeds(query string, page, size int) (*model.FeedDefPage, error) {
	matchQuery := bleve.NewMatchQuery(query)
	searchRequest := bleve.NewSearchRequest(matchQuery)
	searchRequest.Size = size
	searchRequest.From = (page - 1) * size
	searchResults, err := store.index.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("unable to search feed: %w", err)
	}

	result := model.FeedDefPage{
		Total: int(searchResults.Total),
		Page:  page,
		Size:  size,
	}
	hits := searchResults.Hits
	for _, doc := range hits {
		feed, err := store.GetFeed(doc.ID)
		if err != nil {
			return nil, fmt.Errorf("unable to get feed %s: %w", doc.ID, err)
		}
		result.Feeds = append(result.Feeds, *feed)
	}

	return &result, nil
}
