CREATE TABLE IF NOT EXISTS grocery_list (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
);

ALTER TABLE grocery_list ADD CONSTRAINT check_name_length CHECK (char_length(name) > 0 AND char_length(name) <= 100);

