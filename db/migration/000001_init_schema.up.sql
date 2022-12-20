CREATE TABLE users (
    "id" bigserial PRIMARY KEY,
    "username" varchar NOT NULL UNIQUE,
    "email" varchar NOT NULL UNIQUE,
    "password" varchar NOT NULL
);

CREATE TABLE images (
    "id" bigserial PRIMARY KEY,
    "url" varchar NOT NULL,
    "author" varchar NOT NULL REFERENCES users (username) ON DELETE CASCADE
);
