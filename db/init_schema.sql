CREATE TABLE IF NOT EXISTS users
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255)        NOT NULL,
    role     VARCHAR(8)          NOT NULL
);

CREATE TABLE IF NOT EXISTS customers
(
    user_id       INTEGER PRIMARY KEY REFERENCES users (id) NOT NULL,
    start_capital INTEGER                                   NOT NULL
);

CREATE TABLE IF NOT EXISTS porters
(
    user_id    INTEGER PRIMARY KEY REFERENCES users (id) NOT NULL,
    max_weight INTEGER                                   NOT NULL,
    drunk      BOOLEAN                                   NOT NULL,
    fatigue    DOUBLE PRECISION                          NOT NULL,
    salary     INTEGER                                   NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255)                  NOT NULL,
    weight      INTEGER                       NOT NULL,
    customer_id INTEGER REFERENCES users (id) NOT NULL,
    porter_id   INTEGER REFERENCES users (id)
);