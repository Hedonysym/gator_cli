-- +goose up
create table posts (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    title text not null,
    url text not null unique,
    description text, 
    published_at timestamp,
    feed_id uuid not null references feeds(id) on delete cascade
);

-- +goose down
drop table posts;