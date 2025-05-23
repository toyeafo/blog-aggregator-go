-- +goose Up
create table posts (
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    title text not null,
    url text unique not null,
    description text,
    published_at timestamp,
    feed_id uuid not null REFERENCES feeds on delete CASCADE,
    FOREIGN KEY(feed_id) REFERENCES feeds (id)
);

-- +goose Down
drop table posts;