-- migrate:up
create table memory_references (
    id bigint primary key AUTO_INCREMENT not null,
    memory_id bigint,
        FOREIGN KEY (memory_id) REFERENCES memories(id) ON DELETE CASCADE,
    title varchar(128) not null,
    link text,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

-- migrate:down
drop table memory_references;

