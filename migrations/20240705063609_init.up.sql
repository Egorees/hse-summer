CREATE TABLE IF NOT EXISTS users
(
    id            serial not null unique,
    username      varchar(255) unique,
    amount        int
);