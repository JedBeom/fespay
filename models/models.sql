DROP TABLE IF EXISTS students, booths, products, sellers,
    sessions, access_logs, orders;

CREATE TABLE students
(
    id         SERIAL PRIMARY KEY,
    grade      INTEGER           NOT NULL,
    class      INTEGER           NOT NULL,
    number     INTEGER           NOT NULL,
    unique (grade, class, number),
    name       VARCHAR(7)        NOT NULL,
    barcode_id CHAR(5) UNIQUE,
    coin       INTEGER DEFAULT 0 NOT NULL,

    updated_at TIMESTAMPTZ
);

CREATE TABLE booths
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(15)       NOT NULL UNIQUE,
    coin       INTEGER DEFAULT 0 NOT NULL,

    updated_at TIMESTAMPTZ
);

CREATE TABLE products
(
    id       TEXT PRIMARY KEY NOT NULL,
    name     VARCHAR(13)      NOT NULL,
    price    INTEGER          NOT NULL,
    booth_id INTEGER          NOT NULL,

    FOREIGN KEY (booth_id) REFERENCES booths (id)
);

CREATE TABLE sellers
(
    id         SERIAL PRIMARY KEY,
    student_id INTEGER               NOT NULL UNIQUE,
    booth_id   INTEGER               NOT NULL,

    login_id   TEXT                  NOT NULL UNIQUE,
    pin        CHAR(6)               NOT NULL,

    permission INTEGER     DEFAULT 0 NOT NULL,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (booth_id) REFERENCES booths (id) ON DELETE RESTRICT,
    FOREIGN KEY (student_id) REFERENCES students (id) ON DELETE RESTRICT
);

CREATE TABLE sessions
(
    id         TEXT PRIMARY KEY NOT NULL,
    seller_id  TEXT             NOT NULL,

    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE access_logs
(
    id         SERIAL PRIMARY KEY,
    date       TIMESTAMPTZ DEFAULT current_timestamp,
    path       TEXT NOT NULL,
    session_id TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    ip         TEXT NOT NULL,


    FOREIGN KEY (session_id) REFERENCES sessions (id)
);

CREATE TABLE orders
(
    id          TEXT PRIMARY KEY,
    date        TIMESTAMPTZ      DEFAULT current_timestamp,

    student_id  INTEGER NOT NULL,
    seller_id   INTEGER NOT NULL,

    sub_total   INTEGER NOT NULL,
    discount    INTEGER NOT NULL DEFAULT 0,
    grand_total INTEGER NOT NULL,

    is_canceled BOOL    NOT NULL DEFAULT false,

    FOREIGN KEY (student_id) REFERENCES students (id) ON DELETE RESTRICT,
    FOREIGN KEY (seller_id) REFERENCES sellers (id) ON DELETE RESTRICT

)