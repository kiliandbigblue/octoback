DROP INDEX IF EXISTS idx_recipe_tenant_id;

-- SQLite doesn't support DROP COLUMN directly, so we need to recreate the table
CREATE TABLE recipe_backup AS SELECT id, name, description, servings, prep_time_minutes, cook_time_minutes, created_at, version FROM recipe;
DROP TABLE recipe;
CREATE TABLE recipe (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL CHECK (length(name) > 0 AND length(name) <= 100),
    description TEXT NOT NULL DEFAULT '' CHECK (length(description) <= 1000),
    servings INTEGER NOT NULL DEFAULT 1 CHECK (servings > 0),
    prep_time_minutes INTEGER NOT NULL DEFAULT 0 CHECK (prep_time_minutes >= 0),
    cook_time_minutes INTEGER NOT NULL DEFAULT 0 CHECK (cook_time_minutes >= 0),
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    version INTEGER NOT NULL DEFAULT 1
);
INSERT INTO recipe (id, name, description, servings, prep_time_minutes, cook_time_minutes, created_at, version) SELECT id, name, description, servings, prep_time_minutes, cook_time_minutes, created_at, version FROM recipe_backup;
DROP TABLE recipe_backup;

