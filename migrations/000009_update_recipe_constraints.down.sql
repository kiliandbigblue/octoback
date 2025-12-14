-- Revert to original constraints (servings > 0)

-- Create table with original constraints
CREATE TABLE IF NOT EXISTS recipe_new (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL CHECK (length(name) > 0 AND length(name) <= 100),
    description TEXT NOT NULL DEFAULT '' CHECK (length(description) <= 1000),
    servings INTEGER NOT NULL DEFAULT 1 CHECK (servings > 0),
    prep_time_minutes INTEGER NOT NULL DEFAULT 0 CHECK (prep_time_minutes >= 0),
    cook_time_minutes INTEGER NOT NULL DEFAULT 0 CHECK (cook_time_minutes >= 0),
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    version INTEGER NOT NULL DEFAULT 1,
    tenant_id INTEGER REFERENCES tenant (id) ON DELETE CASCADE
);

-- Copy data from current table (update servings = 0 to 1 to satisfy constraint)
INSERT INTO recipe_new (id, name, description, servings, prep_time_minutes, cook_time_minutes, created_at, version, tenant_id)
SELECT id, name, description, CASE WHEN servings = 0 THEN 1 ELSE servings END, prep_time_minutes, cook_time_minutes, created_at, version, tenant_id
FROM recipe;

-- Drop current table
DROP TABLE recipe;

-- Rename new table to original name
ALTER TABLE recipe_new RENAME TO recipe;

-- Recreate index for tenant lookups
CREATE INDEX IF NOT EXISTS idx_recipe_tenant_id ON recipe (tenant_id);

