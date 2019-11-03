DROP TABLE IF EXISTS booths;
CREATE TABLE booths
(
    id          TEXT PRIMARY KEY,
    coin        INTEGER DEFAULT 0,
    name        VARCHAR(15) NOT NULL UNIQUE,
    description VARCHAR(200),
    location    VARCHAR(15),

    status      INTEGER DEFAULT 1,
    updated_at  TIMESTAMPTZ

);

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id         TEXT PRIMARY KEY,

    coin       INTEGER DEFAULT 0,
    booth_id   TEXT,
    login_id   TEXT,
    password   TEXT,
    type       INTEGER    NOT NULL,

    card_code  CHAR(5) UNIQUE,
    grade      INTEGER,
    class      INTEGER,
    number     INTEGER,
    unique (grade, class, number),
    name       VARCHAR(7) NOT NULL,

    status     INTEGER DEFAULT 1,
    updated_at TIMESTAMPTZ,

    FOREIGN KEY (booth_id) REFERENCES booths (id)
);

DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions
(
    id         TEXT PRIMARY KEY NOT NULL,
    user_id    TEXT             NOT NULL,
    user_agent TEXT             NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id),

    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    deleted_at TIMESTAMPTZ
);

DROP TABLE IF EXISTS access_logs;
CREATE TABLE access_logs
(
    id         TEXT PRIMARY KEY,
    session_id TEXT,
    ip         TEXT NOT NULL,
    method     TEXT NOT NULL,
    path       TEXT NOT NULL,

    created_at TIMESTAMPTZ DEFAULT current_timestamp,

    FOREIGN KEY (session_id) REFERENCES sessions (id) ON DELETE RESTRICT
);

DROP TABLE IF EXISTS records CASCADE;
CREATE TABLE records
(
    id            TEXT PRIMARY KEY,

    staff_id      TEXT    NOT NULL,
    booth_id      TEXT    NOT NULL,
    user_id       TEXT,
    amount        INTEGER NOT NULL,

    type          INTEGER NOT NULL,

    access_log_id TEXT,
    created_at    TIMESTAMPTZ DEFAULT current_timestamp,
    paid_at       TIMESTAMPTZ,
    canceled_at   TIMESTAMPTZ,

    FOREIGN KEY (staff_id) REFERENCES users (id) ON DELETE RESTRICT,
    FOREIGN KEY (booth_id) REFERENCES booths (id) ON DELETE RESTRICT,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE RESTRICT
    -- FOREIGN KEY (access_log_id) REFERENCES access_logs (id) ON DELETE RESTRICT
);
