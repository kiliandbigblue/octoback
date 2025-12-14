CREATE TABLE IF NOT EXISTS grocery_item (
    id INTEGER PRIMARY KEY,
    grocery_list_id INTEGER NOT NULL,
    name TEXT NOT NULL CHECK (length(name) > 0 AND length(name) <= 100),
    quantity INTEGER NOT NULL DEFAULT 1 CHECK (quantity > 0),
    checked INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    version INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY (grocery_list_id) REFERENCES grocery_list (id) ON DELETE CASCADE
);
