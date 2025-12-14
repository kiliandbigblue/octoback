-- Add tenant_id column to grocery_list table
ALTER TABLE grocery_list ADD COLUMN tenant_id INTEGER REFERENCES tenant (id) ON DELETE CASCADE;

-- Create index for faster tenant-based lookups
CREATE INDEX IF NOT EXISTS idx_grocery_list_tenant_id ON grocery_list (tenant_id);

