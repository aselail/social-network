-- Migration to create the user table
CREATE TABLE IF NOT EXISTS user
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    archive    BOOLEAN DEFAULT FALSE NOT NULL,
    email      TEXT                  NOT NULL UNIQUE,
    password   TEXT                  NOT NULL,
    first_name TEXT                  NOT NULL,
    last_name  TEXT                  NOT NULL,
    metadata   JSON
);