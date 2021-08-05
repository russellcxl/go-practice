CREATE
    DATABASE IF NOT EXISTS assignment_3;

USE
    assignment_3;

DROP TABLE IF EXISTS leaves;
DROP TABLE IF EXISTS users;

CREATE TABLE users
(
    user_id       int          NOT NULL AUTO_INCREMENT,
    password      varchar(100) DEFAULT '/GKA',
    name          varchar(100) NOT NULL,
    team_id       int          DEFAULT 0,
    role          int          NOT NULL,
    leave_balance int          DEFAULT 10,
    PRIMARY KEY (user_id)
);

CREATE TABLE leaves
(
    leave_id    int NOT NULL AUTO_INCREMENT,
    user_id     int NOT NULL,
    team_id     int NOT NULL,
    start_time  int NOT NULL,
    end_time    int NOT NULL,
    days_taken  int NOT NULL,
    status      int DEFAULT 0,
    approver_id int NOT NULL,
    PRIMARY KEY (leave_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

INSERT INTO users (user_id, name, team_id, role)
VALUES (1000, 'Bill', 100, 2);
INSERT INTO users (name, team_id, role)
VALUES ('Beatrix', 100, 1),
       ('Budd', 100, 1),
       ('Elle', 100, 1),
       ('O-Ren', 100, 1);

INSERT INTO leaves (leave_id, user_id, team_id, start_time, end_time, days_taken, status, approver_id)
VALUES (2000, 1001, 100, 1636588800, 1636675200, 1, 0, 1000);
INSERT INTO leaves (user_id, team_id, start_time, end_time, days_taken, status, approver_id)
VALUES (1001, 100, 1633737600, 1633910400, 1, 0, 1000);



