CREATE TABLE IF NOT EXISTS user_info {
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id varchar(36) NOT NULL,
    lastname varchar(255) NOT NULL,
    firstname varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
}

CREATE TABLE IF NOT EXISTS user_auth {
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uid INTEGER NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    foreign key (uid) references user(id),
}
