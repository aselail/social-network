-- Migration to create the user table
CREATE TABLE IF NOT EXISTS user
(
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    archive         BOOLEAN DEFAULT FALSE NOT NULL,
    public          BOOLEAN DEFAULT TRUE  NOT NULL,
    email           TEXT                  NOT NULL UNIQUE,
    password        TEXT                  NOT NULL,
    first_name      TEXT                  NOT NULL,
    last_name       TEXT                  NOT NULL,
    age             INT                   NOT NULL,
    nickname        TEXT,
    about           TEXT,
    gender          INT,
    profile_picture INT
);

-- Create an index on user.email for faster lookups
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON user (email);