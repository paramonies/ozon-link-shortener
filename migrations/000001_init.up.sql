CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    long_url VARCHAR(300),
    short_id VARCHAR(30),
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);