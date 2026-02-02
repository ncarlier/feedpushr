package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

// BuildInitialIndex create initial index (only if empty)
func (store *SQLiteStore) BuildInitialIndex(ctx context.Context) error {
	// FTS5 index is automatically maintained via triggers
	// Check if we need to rebuild the index
	var count int
	err := store.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM feeds_fts").Scan(&count)
	if err != nil {
		return fmt.Errorf("unable to check FTS index: %w", err)
	}

	if count == 0 {
		// Rebuild FTS index from existing feeds
		_, err := store.db.ExecContext(ctx, "INSERT INTO feeds_fts(feeds_fts) VALUES('rebuild')")
		if err != nil {
			return fmt.Errorf("unable to rebuild FTS index: %w", err)
		}
	}

	return nil
}

// SearchFeeds search feeds using SQLite FTS5
func (store *SQLiteStore) SearchFeeds(ctx context.Context, query string, page, size int) (*model.FeedDefPage, error) {
	offset := (page - 1) * size

	// Count total matches
	var total int
	err := store.db.QueryRowContext(ctx, `
		SELECT COUNT(*)
		FROM feeds_fts
		WHERE feeds_fts MATCH ?
	`, query).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("unable to count search results: %w", err)
	}

	// Get paginated results
	rows, err := store.db.QueryContext(ctx, `
		SELECT f.id, f.xml_url, f.html_url, f.hub_url, f.title, f.status, f.tags, f.cdate, f.mdate
		FROM feeds_fts fts
		JOIN feeds f ON f.id = fts.id
		WHERE feeds_fts MATCH ?
		ORDER BY rank
		LIMIT ? OFFSET ?
	`, query, size, offset)
	if err != nil {
		return nil, fmt.Errorf("unable to search feeds: %w", err)
	}
	defer rows.Close()

	result := model.FeedDefPage{
		Total: total,
		Page:  page,
		Size:  size,
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
