DROP TABLE messages;

CREATE TABLE messages (
        ID              SERIAL PRIMARY KEY NOT NULL,
        name            VARCHAR(255) NOT NULL,
        message         TEXT NOT NULL,
        created_at      timestamp NOT NULL
);

-- INSERT INTO messages
-- (name, message, created_at)
-- VALUES ('Bifur', 'Hey folks!', '2022-02-01'),
-- ('Bofur', 'WHATUP!', '2022-02-01'),
-- ('BLOINfur', 'YO YOYO!', '2022-02-01');
