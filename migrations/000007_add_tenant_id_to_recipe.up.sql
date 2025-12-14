-- Add tenant_id column to recipe table
ALTER TABLE recipe ADD COLUMN tenant_id INTEGER REFERENCES tenant (id) ON DELETE CASCADE;

-- Create index for faster tenant-based lookups
CREATE INDEX IF NOT EXISTS idx_recipe_tenant_id ON recipe (tenant_id);

