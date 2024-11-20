CREATE DATABASE task_management;

CREATE TABLE tasks
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    status TEXT NOT NULL,
    description TEXT,
    due_date DATE,
    priority VARCHAR(20)
);

CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

INSERT INTO users
    (username, password)
VALUES
    ('admin', 'password');