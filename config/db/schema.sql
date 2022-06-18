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
    id         serial PRIMARY KEY,
    name       varchar(255) NOT NULL UNIQUE,
    created_at timestamp    NOT NULL DEFAULT (NOW()),
    updated_at timestamp    NOT NULL DEFAULT (NOW())
    );

DROP TABLE IF EXISTS categories;

CREATE TABLE IF NOT EXISTS menus
(
    id          serial PRIMARY KEY,
    id_category serial       NOT NULL REFERENCES categories,
    name        varchar(255) NOT NULL UNIQUE,
    price       integer      NOT NULL,
    created_at  timestamp    NOT NULL DEFAULT (NOW()),
    updated_at  timestamp    NOT NULL DEFAULT (NOW())
    );

DROP TABLE IF EXISTS menus;

CREATE TABLE IF NOT EXISTS orders
(
    id         serial PRIMARY KEY,
    id_user    uuid         NOT NULL REFERENCES users,
    name       varchar(100) NOT NULL,
    created_at timestamp    NOT NULL DEFAULT (NOW()),
    updated_at timestamp    NOT NULL DEFAULT (NOW())
    );

DROP TABLE IF EXISTS orders;

CREATE TABLE IF NOT EXISTS order_menus
(
    id         serial PRIMARY KEY,
    id_order   serial    NOT NULL REFERENCES orders,
    id_menu    serial    NOT NULL REFERENCES menus,
    created_at timestamp NOT NULL DEFAULT (NOW()),
    updated_at timestamp NOT NULL DEFAULT (NOW())
    );

DROP TABLE IF EXISTS order_menus;

CREATE TABLE IF NOT EXISTS payments
(
    id          serial PRIMARY KEY,
    id_user     uuid      NOT NULL REFERENCES users,
    id_order    serial    NOT NULL REFERENCES orders,
    total_price numeric   NOT NULL,
    created_at  timestamp NOT NULL DEFAULT (NOW()),
    updated_at  timestamp NOT NULL DEFAULT (NOW())
    );

DROP TABLE IF EXISTS payments;