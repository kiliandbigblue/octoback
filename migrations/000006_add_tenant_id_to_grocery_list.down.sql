DROP INDEX IF EXISTS idx_grocery_list_tenant_id;

-- SQLite doesn't support DROP COLUMN directly, so we need to recreate the table
CREATE TABLE grocery_list_backup AS SELECT id, name, created_at, version FROM grocery_list;
DROP TABLE grocery_list;
CREATE TABLE grocery_list (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL CHECK (length(name) > 0 AND length(name) <= 100),
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    version INTEGER NOT NULL DEFAULT 1
);
INSERT INTO grocery_list (id, name, created_at, version) SELECT id, name, created_at, version FROM grocery_list_backup;
DROP TABLE grocery_list_backup;

