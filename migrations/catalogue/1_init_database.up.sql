CREATE TYPE book_type AS ENUM ('Restricted', 'Circulating');

CREATE TABLE books
(
    id         SERIAL PRIMARY KEY,
    isbn       VARCHAR(255) NOT NULL,
    title      VARCHAR(500) NOT NULL,
    author     VARCHAR(255) NOT NULL,
    created_at timestamptz  NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz  NOT NULL DEFAULT current_timestamp,
    deleted_at timestamptz           DEFAULT NULL,
    CONSTRAINT uq_books_isbn UNIQUE (isbn)
);

CREATE TABLE book_instances
(
    id                SERIAL PRIMARY KEY,
    book_id           uuid         NOT NULL,
    book_isbn         VARCHAR(255) NOT NULL,
    library_branch_id uuid         NOT NULL,
    book_type         book_type    NOT NULL,
    created_at        timestamptz  NOT NULL DEFAULT current_timestamp,
    updated_at        timestamptz  NOT NULL DEFAULT current_timestamp,
    deleted_at        timestamptz           DEFAULT NULL,
    CONSTRAINT uq_book_instances_book_id UNIQUE (book_id),
    CONSTRAINT fk_book_instances_books FOREIGN KEY (book_isbn) REFERENCES books(isbn)
);