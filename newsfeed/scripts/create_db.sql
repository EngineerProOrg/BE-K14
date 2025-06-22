create database newsfeed;

create table users (
    id bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_name varchar(20),
    hash_password varchar(32),
    email varchar(50),
    display_name nvarchar(20),
    dob varchar(8),
    removed boolean
);

create table user_user (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id bigint,
    follower_id bigint,
    follow_timestamp int,
    removed boolean
);
