-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE buys
(
    id         serial primary key NOT NULL,
    price      INT                not NULL,
    created_at timestamp default now(),
    user_id    INT                NOT NULL,
    company_id INT                NOT NULL
);

CREATE TABLE companies
(
    id   serial primary key NOT NULL,
    name varchar(255) unique
);

CREATE TABLE operations
(
    id         serial primary key NOT NULL,
    price      INT                not NULL,
    time       timestamp default now(),
    seller_id  INT                NOT NULL,
    company_id INT                NOT NULL,
    buyer_id   INT                NOT NULL
);

CREATE TABLE sales
(
    id         serial primary key NOT NULL,
    price      INT                not NULL,
    created_at timestamp default now(),
    user_id    INT                NOT NULL,
    company_id INT                NOT NULL
);

CREATE TABLE stock_portfolios
(
    id         serial primary key NOT NULL,
    count      INT                NOT NULL,
    user_id    INT                NOT NULL,
    company_id INT                NOT NULL
);

CREATE TABLE users
(
    id     serial primary key  NOT NULL,
    login  varchar(255) unique NOT NULL,
    wealth INT                 not NULL
    check (wealth > 0)
);

CREATE TABLE secrets
(
    user_id serial primary key not null,
    token   varchar(255)       not null
);


CREATE OR REPLACE FUNCTION random_between(low INT ,high INT)
    RETURNS INT AS
$$
BEGIN
    RETURN floor(random()* (high-low + 1) + low);
END;
$$ language 'plpgsql' STRICT;

CREATE OR REPLACE FUNCTION random_string(len int)
    RETURNS text AS
$$
BEGIN
    RETURN substr(md5(random()::text), 0, len);
END;
$$ language 'plpgsql' STRICT;

insert into companies(id, name)
select
    random_between(1,100),
    random_string(20)

from generate_series(1, 100)
on conflict do nothing;


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

drop table users, buys, sales, operations, companies, stock_portfolios, secrets;
-- +goose StatementEnd
