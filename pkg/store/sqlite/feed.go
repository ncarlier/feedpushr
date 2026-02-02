package store

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// ExistsFeed returns true if a feed exists for this url.
func (store *SQLiteStore) ExistsFeed(ctx context.Context, url string) bool {
	hasher := md5.New()
	hasher.Write([]byte(url))
	id := hex.EncodeToString(hasher.Sum(nil))

	var exists bool
	err := store.db.QueryRowContext(ctx, "SELECT COUNT(*) > 0 FROM feeds WHERE id = ?", id).Scan(&exists)
	return err == nil && exists
}

// GetFeed returns a stored Feed.
func (store *SQLiteStore) GetFeed(ctx context.Context, id string) (*model.FeedDef, error) {
	var feed model.FeedDef
	var tagsJSON sql.NullString

	err := store.db.QueryRowContext(ctx, `
		SELECT id, xml_url, html_url, hub_url, title, status, tags, cdate, mdate
		FROM feeds WHERE id = ?
	`, id).Scan(
		&feed.ID,
		&feed.XMLURL,
		&feed.HTMLURL,
		&feed.HubURL,
		&feed.Title,
		&feed.Status,
		&tagsJSON,
		&feed.Cdate,
		&feed.Mdate,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrFeedNotFound
		}
		return nil, err
	}

	// Unmarshal tags
	if tagsJSON.Valid && tagsJSON.String != "" {
		if err := json.Unmarshal([]byte(tagsJSON.String), &feed.Tags); err != nil {
			return nil, err
		}
	}

	return &feed, nil
}

// DeleteFeed removes a feed.
func (store *SQLiteStore) DeleteFeed(ctx context.Context, id string) (*model.FeedDef, error) {
	feed, err := store.GetFeed(ctx, id)
	if err != nil {
		return nil, err
	}

	// FTS index is automatically updated via trigger
	_, err = store.db.ExecContext(ctx, "DELETE FROM feeds WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func (store *SQLiteStore) assertFeedQuota(ctx context.Context, feed *model.FeedDef) error {
	if store.quota.MaxNbFeeds > 0 {
		var exists bool
		err := store.db.QueryRowContext(ctx, "SELECT COUNT(*) > 0 FROM feeds WHERE id = ?", feed.ID).Scan(&exists)
		if err != nil {
			return err
		}

		if !exists {
			total, err := store.CountFeeds(ctx)
			if err != nil {
				return err
			}
			if total >= store.quota.MaxNbFeeds {
				return common.ErrFeedQuotaExceeded
			}
		}
	}
	return nil
}

// SaveFeed stores a feed.
func (store *SQLiteStore) SaveFeed(ctx context.Context, feed *model.FeedDef) error {
	if err := store.assertFeedQuota(ctx, feed); err != nil {
		return err
	}

	// Marshal tags to JSON
	tagsJSON, err := json.Marshal(feed.Tags)
	if err != nil {
		return err
	}

	_, err = store.db.ExecContext(ctx, `
		INSERT INTO feeds (id, xml_url, html_url, hub_url, title, status, tags, cdate, mdate)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			xml_url = excluded.xml_url,
			html_url = excluded.html_url,
			hub_url = excluded.hub_url,
			title = excluded.title,
			status = excluded.status,
			tags = excluded.tags,
			mdate = excluded.mdate
	`,
		feed.ID,
		feed.XMLURL,
		feed.HTMLURL,
		feed.HubURL,
		feed.Title,
		feed.Status,
		string(tagsJSON),
		feed.Cdate,
		feed.Mdate,
	)

	if err != nil {
		return err
	}

	// FTS index is automatically updated via trigger
	return nil
}

// ListFeeds returns a paginated list of feeds.
func (store *SQLiteStore) ListFeeds(ctx context.Context, page, size int) (*model.FeedDefPage, error) {
	offset := (page - 1) * size

	rows, err := store.db.QueryContext(ctx, `
		SELECT id, xml_url, html_url, hub_url, title, status, tags, cdate, mdate
		FROM feeds
		ORDER BY mdate DESC
		LIMIT ? OFFSET ?
	`, size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	total, err := store.CountFeeds(ctx)
	if err != nil {
		return nil, err
	}

	result := model.FeedDefPage{
		Page:  page,
		Size:  size,
		Total: total,
	}

	for rows.Next() {
		var feed model.FeedDef
		var tagsJSON sql.NullString

		err := rows.Scan(
			&feed.ID,
			&feed.XMLURL,
			&feed.HTMLURL,
			&feed.HubURL,
			&feed.Title,
			&feed.Status,
			&tagsJSON,
			&feed.Cdate,
			&feed.Mdate,
		)
		if err != nil {
			return nil, err
		}

		// Unmarshal tags
		if tagsJSON.Valid && tagsJSON.String != "" {
			if err := json.Unmarshal([]byte(tagsJSON.String), &feed.Tags); err != nil {
				return nil, err
			}
		}

		result.Feeds = append(result.Feeds, feed)
	}

	return &result, rows.Err()
}

// CountFeeds returns total number of feeds.
func (store *SQLiteStore) CountFeeds(ctx context.Context) (int, error) {
	var count int
	err := store.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM feeds").Scan(&count)
	return count, err
}

// ForEachFeed iterates over all feeds
func (store *SQLiteStore) ForEachFeed(ctx context.Context, cb func(*model.FeedDef) error) error {
	rows, err := store.db.QueryContext(ctx, `
		SELECT id, xml_url, html_url, hub_url, title, status, tags, cdate, mdate
		FROM feeds
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var feed model.FeedDef
		var tagsJSON sql.NullString

		err := rows.Scan(
			&feed.ID,
			&feed.XMLURL,
			&feed.HTMLURL,
			&feed.HubURL,
			&feed.Title,
			&feed.Status,
			&tagsJSON,
			&feed.Cdate,
			&feed.Mdate,
		)
		if err != nil {
			return err
		}

		// Unmarshal tags
		if tagsJSON.Valid && tagsJSON.String != "" {
			if err := json.Unmarshal([]byte(tagsJSON.String), &feed.Tags); err != nil {
				return err
			}
		}

		if err := cb(&feed); err != nil {
			return err
		}
	}

	return rows.Err()
}
