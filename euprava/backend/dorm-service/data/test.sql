CREATE TABLE IF NOT EXISTS student (
    id SERIAL PRIMARY KEY,
    name VARCHAR ( 50 ) NOT NULL
    surname VARCHAR ( 50 ) UNIQUE NOT NULL,
);

SELECT id,name,surname FROM student LIMIT 5 OFFSET 1;

UPDATE student 
SET name='nam56',surname='surn56'
WHERE id=8;

SELECT * FROM student;

SELECT * FROM student WHERE id=8;

DELETE FROM student WHERE id=1;

INSERT INTO student(name,surname) VALUES('name3','surname3') ON CONFLICT DO NOTHING;