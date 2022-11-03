-- migrate:up
create table memory_tags (
    id bigint primary key AUTO_INCREMENT not null,
    memory_id bigint,
        FOREIGN KEY (memory_id) REFERENCES memories(id) ON DELETE CASCADE,
    tag_id bigint,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

-- migrate:down
drop table memory_tags;

