CREATE TABLE IF NOT EXISTS tenant (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL CHECK (length(name) > 0 AND length(name) <= 100),
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    version INTEGER NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS tenant_member (
    id INTEGER PRIMARY KEY,
    tenant_id INTEGER NOT NULL,
    user_id TEXT NOT NULL CHECK (length(user_id) > 0 AND length(user_id) <= 255),
    role TEXT NOT NULL DEFAULT 'member' CHECK (role IN ('owner', 'admin', 'member')),
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    FOREIGN KEY (tenant_id) REFERENCES tenant (id) ON DELETE CASCADE,
    UNIQUE (tenant_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_tenant_member_user_id ON tenant_member (user_id);
CREATE INDEX IF NOT EXISTS idx_tenant_member_tenant_id ON tenant_member (tenant_id);

