-- CREATE DATABASE testDB;

-- create table posts (
--     id serial primary key,
--     content text,
--     author varchar(255)
-- );

-- CREATE TABLE `Posts`(
--   `id` SERIAL PRIMARY KEY AUTO_INCREMENT,
--   `title` CHAR(255) NOT NULL,
--   `description` TEXT,
--   `year` YEAR
-- );
SELECT
  *
FROM
  posts;


drop TABLE posts;

create table posts (
  id serial primary key NOT NULL,
  created DATETIME DEFAULT CURRENT_TIMESTAMP,
  title VARCHAR(255) not null,
  content text
);
INSERT INTO
  posts(title, content)
VALUES ( "c", "d" );












-- DELETE FROM posts
-- WHERE id = 6
