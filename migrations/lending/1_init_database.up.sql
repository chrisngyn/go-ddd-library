CREATE TYPE patron_type AS ENUM ('Regular', 'Researcher');
CREATE TYPE book_type AS ENUM ('Restricted', 'Circulating');
CREATE TYPE book_status AS ENUM ('Available', 'OnHold', 'CheckedOut');

CREATE TABLE patrons
(
    id          VARCHAR(255) PRIMARY KEY,
    patron_type patron_type NOT NULL,
    created_at  timestamp   NOT NULL DEFAULT current_timestamp,
    updated_at  timestamp   NOT NULL DEFAULT current_timestamp,
    deleted_at  timestamp            DEFAULT NULL
);

CREATE TABLE books
(
    id                uuid PRIMARY KEY,
    library_branch_id uuid        NOT NULL,
    book_type         book_type   NOT NULL,
    book_status       book_status NOT NULL,
    patron_id         VARCHAR(255),
    hold_till         timestamp,
    checked_out_at    timestamp,
    created_at        timestamp   NOT NULL DEFAULT current_timestamp,
    updated_at        timestamp   NOT NULL DEFAULT current_timestamp,
    deleted_at        timestamp            DEFAULT NULL
);

CREATE TABLE holds
(
    id                BIGSERIAL PRIMARY KEY,
    patron_id         VARCHAR(255) NOT NULL,
    book_id           uuid         NOT NULL,
    library_branch_id uuid         NOT NULL,
    hold_from         timestamp    NOT NULL,
    hold_till         timestamp,
    created_at        timestamp    NOT NULL DEFAULT current_timestamp,
    updated_at        timestamp    NOT NULL DEFAULT current_timestamp,
    deleted_at        timestamp             DEFAULT NULL,
    CONSTRAINT fk_holds_patrons FOREIGN KEY (patron_id) REFERENCES patrons (id),
    CONSTRAINT uq_holds_patron_id_book_id UNIQUE (patron_id, book_id)
);

CREATE TABLE overdue_checkouts
(
    id                BIGSERIAL PRIMARY KEY,
    patron_id         VARCHAR(255) NOT NULL,
    book_id           uuid         NOT NULL,
    library_branch_id uuid         NOT NULL,
    created_at        timestamp    NOT NULL DEFAULT current_timestamp,
    updated_at        timestamp    NOT NULL DEFAULT current_timestamp,
    deleted_at        timestamp             DEFAULT NULL,
    CONSTRAINT fk_overdue_checkouts_patrons FOREIGN KEY (patron_id) REFERENCES patrons(id),
    CONSTRAINT uq_overdue_checkouts_patron_id_book_id UNIQUE (patron_id, book_id)
);