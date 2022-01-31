DROP TABLE messages;

CREATE TABLE messages (
        ID              SERIAL PRIMARY KEY NOT NULL,
        name            VARCHAR(255) NOT NULL,
        message         TEXT NOT NULL,
        created_at      timestamp NOT NULL
);

INSERT INTO messages
(name, message, created_at) VALUES
('Dwalin', 'HEAD-BUTTTTT!!!', '2022-02-01'),
('Balin', 'May my beard be ultra Fluff.', '2022-02-01'),
('Kili', 'I like elfs.', '2022-02-01'),
('Fili', 'Yuck!', '2022-02-01'),
('Dori', 'Potato schnoz.', '2022-02-01'),
('Nori', 'igNORI...', '2022-02-01'),
('Ori', 'hmmmmm', '2022-02-01'),
('Oin', 'WHAT???', '2022-02-01'),
('Gloin', 'That is my son!!!', '2022-02-01'),
('Bifur', 'Hey folks!', '2022-02-01'),
('Bofur', 'WHATUP!', '2022-02-01'),
('Bombur', 'Ho Ho Ho, riding my keg!', '2022-02-01'),
('Thorin', 'MEIN SCHATZZZ!!!', '2022-02-01');
