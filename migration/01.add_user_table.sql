CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL,
    secondname  TEXT NOT NULL,
    patronymic  TEXT NOT NULL,
    benefits    TEXT NOT NULL,
    status      BOOLEAN NOT NULL,
    districts   TEXT NOT NULL,
    addr        TEXT NOT NULL
);

