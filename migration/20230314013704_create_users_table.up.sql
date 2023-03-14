CREATE TABLE IF NOT EXISTS users (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL
);

INSERT INTO
    users (id, name)
VALUES
    ('id1', 'nathan'),
    ('id2', 'tobing'),
    ('id3', 'fernando');

