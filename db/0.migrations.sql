CREATE TABLE IF NOT EXISTS file_links (
    id UUID PRIMARY KEY,
    link TEXT NOT NULL,
    caption TEXT,
    category TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


DROP TABLE IF EXISTS file_links;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

-- Insert default admin user (password needs to be hashed)
INSERT INTO users (id, email, password, role, created_at, updated_at)
VALUES (
    gen_random_uuid(), 
    'admin@example.com',
    '$2a$10$your_hashed_password', -- You'll need to generate this
    'admin',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

DROP INDEX IF EXISTS idx_users_email;
DROP TABLE IF EXISTS users;

ALTER TABLE file_links
ADD COLUMN user_id UUID NOT NULL,
ADD CONSTRAINT fk_file_links_user 
    FOREIGN KEY (user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE;

CREATE INDEX idx_file_links_user_id ON file_links(user_id);

DROP INDEX IF EXISTS idx_file_links_user_id;
ALTER TABLE file_links
DROP CONSTRAINT IF EXISTS fk_file_links_user,
DROP COLUMN IF EXISTS user_id;