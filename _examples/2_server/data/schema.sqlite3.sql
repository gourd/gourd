--
-- example schema
-- --------------
-- schema file for sqlite3 database
-- to be used in this example
--

CREATE TABLE posts (
  `id`    INTEGER PRIMARY KEY,
  `uid`   INTEGER,
  `title` TEXT,
  `body`  TEXT,
  `size`  INTEGER,
  `date`  INTEGER
);

CREATE TABLE comments (
  `id`      INTEGER PRIMARY KEY,
  `post_id` INTEGER,
  `title`   TEXT,
  `body`    TEXT,
  `size`    INTEGER,
  `date`    INTEGER
);
