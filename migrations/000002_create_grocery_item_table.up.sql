CREATE TABLE IF NOT EXISTS grocery_item (
    id bigserial PRIMARY KEY,
    grocery_list_id bigint NOT NULL,
    name text NOT NULL,
    quantity integer NOT NULL DEFAULT 1,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1,
    FOREIGN KEY (grocery_list_id) REFERENCES grocery_list (id) ON DELETE CASCADE
);

ALTER TABLE grocery_item ADD CONSTRAINT check_name_length CHECK (char_length(name) > 0 AND char_length(name) <= 100);
ALTER TABLE grocery_item ADD CONSTRAINT check_quantity CHECK (quantity > 0);

