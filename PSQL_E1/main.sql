-- \d to reveal tables
DROP TABLE tblname;
CREATE TABLE employees (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    score INT DEFAULT 10,
    salary REAL
);
--
INSERT INTO employees (name, score, salary) VALUES('Daniel', 23, 55000);
INSERT INTO employees (name, score, salary) VALUES('Arin', 25, 65000);
INSERT INTO employees (name, score, salary) VALUES('Juan', 24, 72000);
INSERT INTO employees (name, score, salary) VALUES('Shen', 26, 64000);
INSERT INTO employees (name, score, salary) VALUES('Myke', 27, 58000);
INSERT INTO employees (name, score, salary) VALUES('McLeod', 26, 72000);
INSERT INTO employees (name, score, salary) VALUES('James', 32, 35000);
--
SELECT * FROM employees;
