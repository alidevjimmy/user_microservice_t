CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    phone      VARCHAR(12)  NOT NULL UNIQUE,
    username   VARCHAR(20)  NOT NULL UNIQUE,
    password   VARCHAR(300) NOT NULL,
    name       VARCHAR(255) NOT NULL,
    family     VARCHAR(255) NOT NULL,
    age        INT          NOT NULL,
    active     BOOL         NOT NULL DEFAULT (FALSE),
    blocked    BOOL         NOT NULL DEFAULT (FALSE),
    is_admin   BOOL         NOT NULL DEFAULT (FALSE),
    created_at TIMESTAMP    NOT NULL DEFAULT now(),
    updated_at TIMESTAMP    NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP    NOT NULL
);

CREATE TYPE code_purposes AS ENUM ('RESET_PASSWORD','ACTIVATION');

CREATE TABLE IF NOT EXISTS verification_codes
(
    id              SERIAL PRIMARY KEY,
    user_id         INT           NOT NULL,
    code            INT           NOT NULL,
    code_purpose    code_purposes NOT NULL,
    expiration_time TIMESTAMP     NOT NULL,
    created_at      TIMESTAMP     NOT NULL DEFAULT now(),
    CONSTRAINT users_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);