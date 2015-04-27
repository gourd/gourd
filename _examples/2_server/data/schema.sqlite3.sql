--
-- example schema
-- --------------
-- schema file for sqlite3 database
-- to be used in this example
--

DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
  `id`    INTEGER PRIMARY KEY,
  `uid`   INTEGER,
  `title` TEXT,
  `body`  TEXT,
  `size`  INTEGER,
  `date`  INTEGER
);

DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
  `id`      INTEGER PRIMARY KEY,
  `post_id` INTEGER,
  `title`   TEXT,
  `body`    TEXT,
  `size`    INTEGER,
  `date`    INTEGER
);
