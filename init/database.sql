CREATE TABLE library_user
(
  id SERIAL PRIMARY KEY,
  username VARCHAR(20) NOT NULL,
  email VARCHAR(20) NOT NULL,
  password VARCHAR(256) NOT NULL
);

CREATE INDEX ON library_user(email);

CREATE TABLE author
(
  id SERIAL PRIMARY KEY,
  firstname VARCHAR(64) NOT NULL,
  lastname VARCHAR(64) NOT NULL,
  birthdate DATE NOT NULL
);

CREATE INDEX ON author(firstname, lastname);

CREATE TABLE book
(
  id SERIAL PRIMARY KEY,
  title VARCHAR(64) NOT NULL,
  release_date DATE NOT NULL
);

CREATE INDEX ON book(title);

CREATE TABLE book_copy
(
  id SERIAL PRIMARY KEY,
  book_id INTEGER REFERENCES book(id)
);

CREATE TABLE book_author
(
  book_id INTEGER REFERENCES book(id) ON UPDATE CASCADE ON DELETE CASCADE,
  author_id INTEGER REFERENCES author(id) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT book_author_id PRIMARY KEY (author_id, book_id)
);

CREATE TABLE history
(
  id SERIAL PRIMARY KEY,
  library_user_id INTEGER REFERENCES library_user(id) ON UPDATE CASCADE ON DELETE CASCADE,
  book_copy_id INTEGER REFERENCES book_copy(id) ON UPDATE CASCADE ON DELETE CASCADE,
  date_from DATE NOT NULL,
  date_to DATE
);

CREATE INDEX ON history(book_copy_id) WHERE date_to isnull;
CREATE INDEX ON history(date_from) WHERE date_to isnull;
CREATE INDEX ON history(library_user_id) WHERE date_to isnull;
