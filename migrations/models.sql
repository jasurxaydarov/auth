CREATE DATABASE auth;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users(

    user_id     UUID            PRIMARY KEY,
    first_name  VARCHAR(124)    NOT NULL,
    last_name   VARCHAR(124)    NOT NULL ,
    user_name   VARCHAR(124)    NOT NULL,
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
INSERT INTO "users" (user_id, first_name, last_name, user_name, gmail, password) VALUES
    (gen_random_uuid(), 'Jasur', 'Xayderov', 'JasurXaydarov', 'xaydarovjasur2005@gmail.com', 'jasur2005'),
    (gen_random_uuid(), 'Jane', 'Smith', 'janesmith', 'janesmith@gmail.com', 'password456'),
    (gen_random_uuid(), 'Alice', 'Johnson', 'alicej', 'alicej@example.com', 'alice@789'),
    (gen_random_uuid(), 'Bob', 'Brown', 'bobbybrown', 'bobby.brown@example.com', 'bob_password'),
    (gen_random_uuid(), 'Charlie', 'Davis', 'charlied', 'charlie.d@example.com', 'charlie123!');