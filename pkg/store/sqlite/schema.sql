-- feedpushr SQLite Schema
-- This file contains the complete database schema for the SQLite store implementation

-- Main feeds table
CREATE TABLE IF NOT EXISTS feeds (
	id TEXT PRIMARY KEY,
	xml_url TEXT NOT NULL,
	html_url TEXT,
	hub_url TEXT,
	title TEXT NOT NULL,
	status TEXT,
	tags TEXT,
	cdate DATETIME NOT NULL,
	mdate DATETIME NOT NULL
);

-- Index on XML URL for faster lookups
CREATE INDEX IF NOT EXISTS idx_feeds_xml_url ON feeds(xml_url);

-- FTS5 virtual table for full-text search on feeds
CREATE VIRTUAL TABLE IF NOT EXISTS feeds_fts USING fts5(
	id UNINDEXED,
	title,
	xml_url,
	tags,
	content='feeds',
	content_rowid='rowid'
);

-- Trigger: Maintain FTS index on INSERT
CREATE TRIGGER IF NOT EXISTS feeds_ai AFTER INSERT ON feeds BEGIN
	INSERT INTO feeds_fts(rowid, id, title, xml_url, tags)
	VALUES (new.rowid, new.id, new.title, new.xml_url, new.tags);
END;

-- Trigger: Maintain FTS index on DELETE
CREATE TRIGGER IF NOT EXISTS feeds_ad AFTER DELETE ON feeds BEGIN
	INSERT INTO feeds_fts(feeds_fts, rowid, id, title, xml_url, tags)
	VALUES ('delete', old.rowid, old.id, old.title, old.xml_url, old.tags);
END;

-- Trigger: Maintain FTS index on UPDATE
CREATE TRIGGER IF NOT EXISTS feeds_au AFTER UPDATE ON feeds BEGIN
	INSERT INTO feeds_fts(feeds_fts, rowid, id, title, xml_url, tags)
	VALUES ('delete', old.rowid, old.id, old.title, old.xml_url, old.tags);
	INSERT INTO feeds_fts(rowid, id, title, xml_url, tags)
	VALUES (new.rowid, new.id, new.title, new.xml_url, new.tags);
END;

-- Outputs table
CREATE TABLE IF NOT EXISTS outputs (
	id TEXT PRIMARY KEY,
	alias TEXT NOT NULL,
	name TEXT NOT NULL,
	description TEXT,
	condition TEXT,
	props TEXT,
	filters TEXT,
	enabled INTEGER NOT NULL,
	nb_success INTEGER NOT NULL DEFAULT 0,
	nb_error INTEGER NOT NULL DEFAULT 0
);

-- Cache table for storing temporary data
CREATE TABLE IF NOT EXISTS cache (
	key TEXT PRIMARY KEY,
	value TEXT NOT NULL,
	date DATETIME NOT NULL
);

-- Index on cache date for efficient eviction
CREATE INDEX IF NOT EXISTS idx_cache_date ON cache(date);
