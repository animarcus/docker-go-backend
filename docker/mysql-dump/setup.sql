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





create table posts (
  id serial primary key,
  content text,
  author varchar(255)
);

INSERT INTO
  posts (content, author)
VALUES
  ('boop', 'marcus'),
  ('poop', 'david'),
  ('i like dicks', 'philip');


-- DELETE FROM posts
-- WHERE id = 6
