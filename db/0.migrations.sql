CREATE TABLE file_links (
    id UUID PRIMARY KEY,
    link TEXT NOT NULL,
    caption TEXT,
    category TEXT,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);