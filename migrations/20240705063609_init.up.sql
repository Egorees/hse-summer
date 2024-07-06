CREATE TABLE IF NOT EXISTS users
(
    username      varchar(255) unique,
    amount        int
);