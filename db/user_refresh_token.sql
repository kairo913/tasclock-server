CREATE TABLE IF NOT EXISTS user_refresh_token {
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    token varchar(255) NOT NULL,
    expire_at TIMESTAMP NOT NULL,
    is_used BOOLEAN DEFAULT FALSE,
}

CREATE EVENT
    delete_expired_token
ON SCHEDULE EVERY 1 DAY
DO
BEGIN
DELETE FROM
    refresh_token
WHERE expire_at < NOW()
END