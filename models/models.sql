DROP TABLE IF EXISTS wallets;
CREATE TABLE wallets
(
    id         TEXT PRIMARY KEY,
    owner_type INTEGER NOT NULL,
    owner_id   TEXT    NOT NULL UNIQUE,
    coin       INTEGER DEFAULT 0,
    updated_at TIMESTAMPTZ
);


DROP TABLE IF EXISTS booths;
CREATE TABLE booths
(
    id          TEXT PRIMARY KEY,
    wallet_id   TEXT        NOT NULL,
    name        VARCHAR(15) NOT NULL UNIQUE,
    description VARCHAR(200),

    status      INTEGER DEFAULT 1,
    updated_at  TIMESTAMPTZ,

    FOREIGN KEY (wallet_id) REFERENCES wallets (id)
);

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id         TEXT PRIMARY KEY,

    wallet_id  TEXT,
    booth_id   TEXT,
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

    FOREIGN KEY (booth_id) REFERENCES booths (id),
    FOREIGN KEY (wallet_id) REFERENCES wallets (id)
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
    session_id TEXT NOT NULL,
    ip         TEXT NOT NULL,
    method     TEXT NOT NULL,
    path       TEXT NOT NULL,

    created_at TIMESTAMPTZ DEFAULT current_timestamp,

    FOREIGN KEY (session_id) REFERENCES sessions (id) ON DELETE RESTRICT
);

DROP TABLE IF EXISTS orders CASCADE;
CREATE TABLE orders
(
    id              TEXT PRIMARY KEY,

    staff_id        TEXT    NOT NULL,
    from_id         TEXT    NOT NULL,
    to_id           TEXT,
    amount          INTEGER NOT NULL,
    refund_order_id TEXT,

    access_log_id   TEXT,
    created_at      TIMESTAMPTZ DEFAULT current_timestamp,
    closed_at       TIMESTAMPTZ,

    FOREIGN KEY (staff_id) REFERENCES users (id) ON DELETE RESTRICT,
    FOREIGN KEY (from_id) REFERENCES wallets (id) ON DELETE RESTRICT,
    FOREIGN KEY (to_id) REFERENCES wallets (id) ON DELETE RESTRICT,
    FOREIGN KEY (refund_order_id) REFERENCES orders (id) ON DELETE CASCADE,
    -- FOREIGN KEY (access_log_id) REFERENCES access_logs (id) ON DELETE RESTRICT
);
