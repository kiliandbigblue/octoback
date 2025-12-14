CREATE TABLE IF NOT EXISTS tenant_invitation (
    id INTEGER PRIMARY KEY,
    tenant_id INTEGER NOT NULL,
    invite_code TEXT NOT NULL UNIQUE CHECK (length(invite_code) > 0 AND length(invite_code) <= 32),
    created_by TEXT NOT NULL CHECK (length(created_by) > 0 AND length(created_by) <= 255),
    expires_at TEXT,
    used_at TEXT,
    used_by TEXT,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    FOREIGN KEY (tenant_id) REFERENCES tenant (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_tenant_invitation_invite_code ON tenant_invitation (invite_code);
CREATE INDEX IF NOT EXISTS idx_tenant_invitation_tenant_id ON tenant_invitation (tenant_id);

