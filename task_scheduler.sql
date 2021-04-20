CREATE DATABASE task_scheduler;

use task_scheduler;

CREATE TABLE task (
    id         VARCHAR(255),
    name       VARCHAR(50),
    created_at DATETIME,
    status     VARCHAR(50),
    data       VARCHAR(50)
);
