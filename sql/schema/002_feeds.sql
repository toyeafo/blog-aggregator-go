-- +goose Up
create table feeds (
    id uuid primary key,
    name text,
    url text unique not null,
    user_id uuid NOT NULL REFERENCES users ON DELETE CASCADE,
    created_at timestamp not null,
    updated_at timestamp not null,
    last_fetched_at timestamp,
    FOREIGN KEY(user_id) REFERENCES users (id)
);

-- +goose Down
drop table feeds;