-- migrate:up
create table memories (
    id bigint primary key AUTO_INCREMENT not null,
    category_id bigint,
        FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
    title varchar(128) not null,
    body text,
    photo varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

-- migrate:down
drop table memories;

