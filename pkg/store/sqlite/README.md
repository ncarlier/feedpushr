# SQLite Store Implementation

This directory contains the SQLite database implementation for feedpushr.

## Features

- Full support for all store interfaces (FeedRepository, OutputRepository, CacheRepository, SearchRepository)
- Uses modernc.org/sqlite (pure Go SQLite driver)
- JSON serialization for complex fields (tags, props, filters)
- **SQLite FTS5 (Full-Text Search) for feed indexing**
- Automatic index maintenance via triggers
- Quota enforcement
- ACID transactions via SQLite

## Usage

To use SQLite as the database backend, specify the `sqlite://` URL scheme:

```bash
# Use SQLite with file path
feedpushr --db "sqlite:///var/opt/feedpushr/data.db"

# Or set environment variable
export FP_DB="sqlite:///var/opt/feedpushr/data.db"
feedpushr
```

## Schema

The SQLite store creates three main tables and one FTS5 virtual table:

### `feeds`
- `id` (TEXT, PRIMARY KEY) - Feed ID (MD5 of XML URL)
- `xml_url` (TEXT) - RSS/Atom feed URL
- `html_url` (TEXT) - Website URL
- `hub_url` (TEXT) - PubSubHubbub hub URL
- `title` (TEXT) - Feed title
- `status` (TEXT) - Aggregation status (running/stopped)
- `tags` (TEXT) - JSON array of tags
- `cdate` (DATETIME) - Creation date
- `mdate` (DATETIME) - Modification date

### `feeds_fts` (FTS5 Virtual Table)
Full-text search index for feeds containing:
- `id` (UNINDEXED) - Feed ID reference
- `title` - Searchable feed title
- `xml_url` - Searchable feed URL
- `tags` - Searchable tags

**Automatic Triggers**: The FTS index is automatically maintained through SQLite triggers on INSERT, UPDATE, and DELETE operations on the feeds table.

### `outputs`
- `id` (TEXT, PRIMARY KEY) - Output ID
- `alias` (TEXT) - Output alias
- `name` (TEXT) - Output name
- `description` (TEXT) - Output description
- `condition` (TEXT) - Conditional expression
- `props` (TEXT) - JSON object of properties
- `filters` (TEXT) - JSON array of filters
- `enabled` (INTEGER) - Whether output is enabled
- `nb_success` (INTEGER) - Success counter
- `nb_error` (INTEGER) - Error counter

### `cache`
- `key` (TEXT, PRIMARY KEY) - Cache key
- `value` (TEXT) - Cache value
- `date` (DATETIME) - Cache entry date

## Implementation Details

### Schema Management
The database schema is defined in `schema.sql` and embedded into the binary using Go's `//go:embed` directive. This approach provides:
- **Version control**: Schema is tracked in a standard SQL file
- **Readability**: Easy to review and understand the complete schema
- **Maintainability**: Schema changes are centralized in one file
- **No runtime dependencies**: Schema is compiled into the binary

### Full-Text Search
The SQLite store uses **SQLite FTS5** for full-text search instead of external indexing libraries:
- **Native integration**: FTS5 is built into SQLite
- **Automatic updates**: Triggers keep the FTS index synchronized with the feeds table
- **Efficient queries**: Uses SQLite's optimized FTS5 ranking algorithm
- **No external dependencies**: Unlike BoltDB which uses Bleve, SQLite handles search internally

### Transaction Safety
All write operations use SQLite transactions to ensure ACID properties.

### JSON Fields
Complex data structures (tags, props, filters) are stored as JSON strings and marshaled/unmarshaled transparently.

### Search Index Maintenance
The FTS5 index is maintained automatically through three triggers:
- `feeds_ai`: After INSERT - adds new feed to FTS index
- `feeds_ad`: After DELETE - removes feed from FTS index  
- `feeds_au`: After UPDATE - updates feed in FTS index

### Quota Management
Feed and output quotas are enforced at save time, preventing quota violations.

## Testing

Tests are located in `pkg/store/test/sqlite_test.go`. Run tests with:

```bash
go test ./pkg/store/test/... -v
```

## Performance Considerations

- SQLite is optimized for single-writer scenarios
- For high-concurrency writes, consider using the BoltDB backend instead
- **FTS5 search is very fast** for typical use cases (thousands of feeds)
- FTS index is stored within the same database file
- Pagination is implemented using LIMIT/OFFSET
- Search results are ranked by SQLite's BM25 algorithm

## Compatibility

- Go 1.24.0+
- modernc.org/sqlite v1.44.3+
- SQLite FTS5 (included in modernc.org/sqlite)

## Comparison with BoltDB Store

| Feature | SQLite Store | BoltDB Store |
|---------|--------------|--------------|
| Search | SQLite FTS5 (built-in) | Bleve (external library) |
| Index Storage | Same database file | Separate index directory |
| Index Maintenance | Automatic triggers | Manual indexing |
| Dependencies | Pure Go SQLite only | BoltDB + Bleve + dependencies |
| Portability | Single database file | Database + index directory |
| Write Concurrency | Single writer | Single writer |
