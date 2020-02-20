create database video_server;
use video_server;

CREATE TABLE users
(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    login_name VARCHAR(64) UNIQUE KEY,
    pwd TEXT
);

CREATE TABLE video_info
(
    id VARCHAR(64) PRIMARY KEY NOT NULL,
    author_id INT UNSIGNED,
    name TEXT,
    display_ctime TEXT,
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments
(
    id VARCHAR(64) PRIMARY KEY NOT NULL,
    video_id VARCHAR(64),
    author_id INT UNSIGNED,
    content TEXT,
    time DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions
(
    -- session_id TINYTEXT PRIMARY KEY NOT NULL,
    session_id VARCHAR(64) PRIMARY KEY NOT NULL,
    TTL TINYTEXT,
    login_name VARCHAR(64)
);