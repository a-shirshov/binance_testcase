CREATE TABLE IF NOT EXISTS symbol (
    id serial PRIMARY KEY not null UNIQUE,
    symbol varchar(20)
);

CREATE TABLE IF NOT EXISTS bprice (
    id serial PRIMARY KEY not null UNIQUE,
    symbol_id int references symbol(id) on delete cascade not null,
    price text,
    timestamp bigint
);