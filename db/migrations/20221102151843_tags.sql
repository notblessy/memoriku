-- migrate:up
create table tags (
    id bigint primary key AUTO_INCREMENT not null,
    name varchar(128) not null,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

-- migrate:down
drop table tags;

