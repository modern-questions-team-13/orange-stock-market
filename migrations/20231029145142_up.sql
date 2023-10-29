-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE buys
(
    id         serial primary key NOT NULL,
    price      INT NULL,
    created_at DATETIME NULL,
    user_id    INT                NOT NULL,
    company_id INT                NOT NULL
) COMMENT 'биды на покупку';

CREATE TABLE companies
(
    id   serial primary key NOT NULL,
    name VARCHAR NULL,
);

CREATE TABLE operations
(
    id         serial primary key NOT NULL,
    price      INT NULL,
    time       DATETIME NULL,
    seller_id  INT                NOT NULL,
    company_id INT                NOT NULL,
    buyer_id   INT                NOT NULL,
) COMMENT 'история успешных операций';

CREATE TABLE sales
(
    id         serial primary key NOT NULL,
    price      INT NULL,
    created_at DATETIME NULL,
    user_id    INT                NOT NULL,
    company_id INT                NOT NULL,
) COMMENT 'биды на продажу';

CREATE TABLE stock_portfolios
(
    id         serial primary key NOT NULL,
    count      INT                NOT NULL,
    user_id    INT primary        NOT NULL,
    company_id INT                NOT NULL
);

CREATE TABLE users
(
    id     serial primary key NOT NULL,
    login  varchar(255)       NOT NULL,
    wealth INT NULL,
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

drop table users, buys, sales, operations, companies, stock_portfolios;
-- +goose StatementEnd
