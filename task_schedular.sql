CREATE DATABASE task_schedular;

use task_schedular;

CREATE TABLE tasks (
                       id VARCHAR(250)  ,
                       name VARCHAR(255) NOT NULL,
                       created_at DATETIME  DEFAULT CURRENT_TIMESTAMP ,
                       status VARCHAR(20),
                       description TEXT,
                       deadline DATETIME,
                       priority INT,
                       assignee VARCHAR(100)
);
