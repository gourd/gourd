--
-- example schema
-- --------------
-- schema file for sqlite3 database
-- to be used in this example
--

--
-- example user table 1
--

DROP TABLE IF EXISTS user_a;
CREATE TABLE user_a (
  `id`       INTEGER PRIMARY KEY,
  `username` TEXT,
  `email`    TEXT,
  `password` TEXT,
  `name`     INTEGER,
  `created`  INTEGER,
  `updated`  INTEGER
);

--
-- example user table 2
--

DROP TABLE IF EXISTS user_b;
CREATE TABLE user_b (
  `id`       TEXT PRIMARY KEY,
  `username` TEXT,
  `email`    TEXT,
  `password` TEXT,
  `name`     INTEGER,
  `created`  INTEGER,
  `updated`  INTEGER
);

