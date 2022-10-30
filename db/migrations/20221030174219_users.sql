-- migrate:up
create table users (
    id bigint primary key AUTO_INCREMENT not null,
    name varchar(128) not null,
    photo varchar(128) not null,
    email varchar(255),
    token_auth text,
    password varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

-- migrate:down
drop table users;

