CREATE TABLE IF NOT EXISTS user_token (
    id TEXT PRIMARY KEY,
    access_token TEXT,
    refresh_token TEXT,
    expiry TIMESTAMP,
    token_type TEXT
);