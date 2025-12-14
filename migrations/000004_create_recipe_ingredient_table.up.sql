CREATE TABLE IF NOT EXISTS recipe_ingredient (
    id INTEGER PRIMARY KEY,
    recipe_id INTEGER NOT NULL,
    name TEXT NOT NULL CHECK (length(name) > 0 AND length(name) <= 100),
    quantity TEXT NOT NULL DEFAULT '' CHECK (length(quantity) <= 50),
    unit TEXT NOT NULL DEFAULT '' CHECK (length(unit) <= 50),
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    version INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY (recipe_id) REFERENCES recipe (id) ON DELETE CASCADE
);

