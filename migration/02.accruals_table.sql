CREATE TABLE IF NOT EXISTS accruals (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date DATETIME NOT NULL,
    past TEXT NOT NULL,
    last TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    status
);


CREATE UNIQUE INDEX email_uindex ON users (email);
CREATE UNIQUE INDEX user_uindex ON users (username);