-- migrate:up
create table memories (
    id varchar(255) primary key not null,
    category_id varchar(255),
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

