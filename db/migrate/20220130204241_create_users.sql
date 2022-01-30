DROP TABLE users;

CREATE TABLE users (
        ID              SERIAL PRIMARY KEY NOT NULL,
        UUID            VARCHAR(64) NOT NULL UNIQUE,
        name            VARCHAR(255) NOT NULL,
        email           VARCHAR(255) NOT NULL UNIQUE,
        password        VARCHAR(255) NOT NULL,
        created_at      timestamp NOT NULL
);
