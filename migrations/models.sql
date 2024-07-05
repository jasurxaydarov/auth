CREATE DATABASE auth;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users(

    user_id     UUID            PRIMARY KEY,
    user_name   VARCHAR(124)    NOT NULL,
    user_role   INT             NOT NULL    DEFAULT 1,
    gmail       VARCHAR(256)    NOT NULL,
    password    VARCHAR(69)     NOT NULL

);


CREATE TABLE fruits(
    fruit id    UUID            PRIMARY KEY,
    fruit_name  VARCHAR(124)    NOT NULL
);

INSERT INTO fruits (fruit_id, fruit_name) VALUES
    (gen_random_uuid(), 'Apple'),
    (gen_random_uuid(), 'Banana'),
    (gen_random_uuid(), 'Cherry'),
    (gen_random_uuid(), 'Date'),
    (gen_random_uuid(), 'Elderberry');

-- Insert values into the "user" table
INSERT INTO "users" (user_id, user_role, user_name, gmail, password) VALUES
    (gen_random_uuid(), 2,'JasurXaydarov', 'xaydarovjasur2005@gmail.com', 'jasur2005'),
    (gen_random_uuid(),2, 'janesmith', 'janesmith@gmail.com', 'password456'),
    (gen_random_uuid(), 1,'alicej', 'alicej@example.com', 'alice@789'),
    (gen_random_uuid(), 1,'Bob', 'bobby.brown@example.com', 'bob_password'),
    (gen_random_uuid(), 1,'Charlie', 'charlie.d@example.com', 'charlie123!');