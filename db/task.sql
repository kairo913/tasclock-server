CREATE TABLE IF NOT EXISTS task {
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id varchar(36) NOT NULL,
    user_id varchar(36) NOT NULL,
    title varchar(255) NOT NULL,
    description TEXT NOT NULL,
    is_done BOOLEAN NOT NULL,
    reward INTEGER NOT NULL,
    elapsed INTEGER NOT NULL,
    deadline TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user_info(user_id)
}
