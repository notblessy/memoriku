-- migrate:up
create table categories (
   id varchar(255) primary key not null,
   group_id varchar(32),
   name varchar(128) not null,
   created_at timestamp,
   updated_at timestamp,
   deleted_at timestamp
);

-- migrate:down
drop table categories;

