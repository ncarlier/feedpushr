package store

import (
	"database/sql"
	_ "embed"
	"fmt"
	"net/url"
	"strings"

	"github.com/ncarlier/feedpushr/v3/pkg/model"
	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schemaSQL string

// SQLiteStore is a data store backed by SQLite
type SQLiteStore struct {
	db    *sql.DB
	quota model.Quota
}

// NewSQLiteStore creates a data store backed by SQLite
func NewSQLiteStore(datasource *url.URL, quota model.Quota) (*SQLiteStore, error) {
	dbPath := datasource.Host + datasource.Path
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open SQLite DB: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("unable to connect to SQLite DB: %w", err)
	}

	// Create tables
	if err := createTables(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("unable to create tables: %w", err)
	}

	return &SQLiteStore{
		db:    db,
		quota: quota,
	}, nil
}

func createTables(db *sql.DB) error {
	// Execute the entire schema in a single transaction
	// SQLite can handle multiple statements separated by semicolons
	// when using multiple Exec calls

	// Split on semicolon but preserve trigger bodies
	var statements []string
	var currentStmt strings.Builder
	inTrigger := false

	lines := strings.Split(schemaSQL, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Skip comments and empty lines
		if trimmed == "" || strings.HasPrefix(trimmed, "--") {
			continue
		}

		currentStmt.WriteString(line)
		currentStmt.WriteString("\n")

		// Detect trigger start
		if strings.Contains(strings.ToUpper(trimmed), "CREATE TRIGGER") {
			inTrigger = true
		}

		// End of statement is semicolon, but not inside a trigger
		if strings.HasSuffix(trimmed, ";") {
			if inTrigger && strings.HasSuffix(trimmed, "END;") {
				// End of trigger
				inTrigger = false
				statements = append(statements, currentStmt.String())
				currentStmt.Reset()
			} else if !inTrigger {
				// Regular statement
				statements = append(statements, currentStmt.String())
				currentStmt.Reset()
			}
		}
	}

	// Execute each statement
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("failed to execute statement: %w\nStatement: %s", err, stmt)
		}
	}

	return nil
}

// Close the DB.
func (store *SQLiteStore) Close() error {
	return store.db.Close()
}
