CREATE DATABASE IF NOT EXISTS cleanarch_go_db;

USE cleanarch_go_db;

DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
    identifier VARCHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    age int NOT NULL
    )  ENGINE=INNODB;

INSERT into users (identifier, name, lastname, age) values ("78535066-801c-4d4a-906f-73949b530e52", "Guilherme", "Rodrigues", 26);