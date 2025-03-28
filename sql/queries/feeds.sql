-- name: CreateFeed :one
insert into feeds (id, name, url, user_id, created_at, updated_at, last_fetched_at)
values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
returning *;

-- name: GetFeed :one
select * from feeds where name = $1;

-- name: GetFeedByURL :one
select * from feeds where url = $1;

-- name: DeleteFeeds :exec
delete from feeds;

-- name: GetFeeds :many
select * from feeds;

-- name: GetFeedsWithName :many
select feeds.*, users.name from feeds left join users on feeds.user_id = users.id;

-- name: MarkFeedFetched :exec
update feeds set last_fetched_at = now(), updated_at = now() where id = $1 
returning *;

-- name: GetNextFeedToFetch :one
select * from feeds order by last_fetched_at asc nulls first limit 1;