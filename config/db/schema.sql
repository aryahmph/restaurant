CREATE TYPE role_type AS ENUM ('WAITER','CASHIER', 'ADMIN');

DROP TYPE IF EXISTS role_type;

CREATE TABLE IF NOT EXISTS users
(
    id            uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    username      varchar(100)     NOT NULL UNIQUE,
    password_hash varchar(255)     NOT NULL,
    role          role_type        NOT NULL,
    created_at    timestamp        NOT NULL DEFAULT (NOW()),
    updated_at    timestamp        NOT NULL DEFAULT (NOW())
);

DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS categories
(
    id         uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name       varchar(255)     NOT NULL UNIQUE,
    created_at timestamp        NOT NULL DEFAULT (NOW()),
    updated_at timestamp        NOT NULL DEFAULT (NOW())
);

DROP TABLE IF EXISTS categories;

CREATE TABLE IF NOT EXISTS uploads
(
    id         uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    filename   uuid             NOT NULL UNIQUE,
    created_at timestamp        NOT NULL DEFAULT (NOW()),
    updated_at timestamp        NOT NULL DEFAULT (NOW())
);

DROP TABLE IF EXISTS uploads;

CREATE TABLE IF NOT EXISTS menus
(
    id             uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name           varchar(255)     NOT NULL UNIQUE,
    price          integer          NOT NULL,
    category_name  varchar(255)     NOT NULL REFERENCES categories (name) ON DELETE SET NULL ON UPDATE CASCADE,
    image_filename uuid             NOT NULL REFERENCES uploads (filename) ON DELETE SET NULL ON UPDATE CASCADE,
    created_at     timestamp        NOT NULL DEFAULT (NOW()),
    updated_at     timestamp        NOT NULL DEFAULT (NOW()),
    deleted_at     timestamp
);

DROP TABLE IF EXISTS menus;

CREATE TYPE order_status_type AS ENUM ('TODO','ON PROGRESS', 'SERVED','DONE');

DROP TYPE IF EXISTS order_status_type;

CREATE TABLE IF NOT EXISTS orders
(
    id         uuid PRIMARY KEY  NOT NULL DEFAULT gen_random_uuid(),
    id_waiter  uuid              NOT NULL REFERENCES users ON DELETE SET NULL ON UPDATE RESTRICT,
    id_cashier uuid              NOT NULL REFERENCES users ON DELETE SET NULL ON UPDATE RESTRICT,
    name       varchar(100)      NOT NULL,
    status     order_status_type NOT NULL DEFAULT 'TODO',
    created_at timestamp         NOT NULL DEFAULT (NOW()),
    updated_at timestamp         NOT NULL DEFAULT (NOW())
);

DROP TABLE IF EXISTS orders;

CREATE TABLE IF NOT EXISTS order_menus
(
    id         uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    id_order   uuid             NOT NULL REFERENCES orders ON DELETE CASCADE ON UPDATE RESTRICT,
    id_menu    serial           NOT NULL REFERENCES menus ON DELETE RESTRICT ON UPDATE RESTRICT,
    has_served bool             NOT NULL DEFAULT FALSE,
    notes      text,
    created_at timestamp        NOT NULL DEFAULT (NOW()),
    updated_at timestamp        NOT NULL DEFAULT (NOW())
);

DROP TABLE IF EXISTS order_menus;