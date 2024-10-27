CREATE TABLE IF NOT EXISTS user {
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id varchar(36) NOT NULL,
    lastname varchar(255) NOT NULL,
    firstname varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
}