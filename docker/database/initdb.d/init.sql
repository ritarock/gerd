CREATE DATABASE IF NOT EXISTS app;

CREATE TABLE IF NOT EXISTS books (
    id             INT         NOT NULL,
    title          VARCHAR(60) NOT NULL,
    published_date DATE,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS categories (
    id       INT         NOT NULL,
    category VARCHAR(60) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS books_categories (
    book_id     INT NOT NULL,
    category_id INT NOT NULL,

    FOREIGN KEY (book_id)     REFERENCES books (id)      ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS authors (
    id INT NOT NULL,
    name VARCHAR(45) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS books_authors (
    book_id   INT NOT NULL,
    author_id INT NOT NULL,
    FOREIGN KEY (book_id)   REFERENCES books (id)   ON DELETE CASCADE,
    FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE CASCADE
);
