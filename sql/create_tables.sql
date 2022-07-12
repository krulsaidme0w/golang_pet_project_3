CREATE TABLE library_user
(
  id VARCHAR(64) NOT NULL PRIMARY KEY,
  username VARCHAR(20) NOT NULL,
  email VARCHAR(20) NOT NULL,
  password VARCHAR(256) NOT NULL
);

CREATE TABLE author
(
  id SERIAL PRIMARY KEY,
  firstname VARCHAR(64) NOT NULL,
  lastname VARCHAR(64) NOT NULL,
  birthdate DATE NOT NULL
);

CREATE TABLE book
(
  id SERIAL PRIMARY KEY,
  title VARCHAR(64) NOT NULL,
  release_date DATE NOT NULL
);

CREATE TABLE book_copy
(
  id SERIAL PRIMARY KEY,
  book_id INTEGER REFERENCES book(id)
);

CREATE TABLE book_author
(
  author_id INTEGER REFERENCES author(id) ON UPDATE CASCADE ON DELETE CASCADE,
  book_id INTEGER REFERENCES book(id) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT book_author_id PRIMARY KEY (author_id, book_id)
);

CREATE TABLE library_history
(
  id SERIAL PRIMARY KEY,
  library_user_id VARCHAR(64) REFERENCES library_user(id) ON UPDATE CASCADE ON DELETE CASCADE,
  book_copy_id INTEGER REFERENCES book_copy(id) ON UPDATE CASCADE ON DELETE CASCADE,
  date_from DATE NOT NULL,
  date_to DATE NOT NULL
);
