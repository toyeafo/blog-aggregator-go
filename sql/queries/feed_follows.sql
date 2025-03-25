-- name: CreateFeedFollow :one
WITH inserted_feed_follow as (insert into feed_follows (id, user_id, feed_id, created_at, updated_at)
values (
    $1,
    $2,
    $3,
    $4,
    $5
)
returning *)

select 
inserted_feed_follow.*, 
users.name as user_name, 
feeds.name as feed_name 
from inserted_feed_follow
inner join feeds on feeds.id = inserted_feed_follow.feed_id
inner join users on users.id = inserted_feed_follow.user_id;

-- name: DeleteFeedFollow :exec
delete from feed_follows;

-- name: GetFeedFollow :many
select * from feed_follows;

-- name: GetFeedsFollowForUser :many
select 
users.name as user_name, 
feeds.name as feed_name,
feed_follows.* from feed_follows
left join users on users.id = feed_follows.user_id
left join feeds on feeds.id = feed_follows.feed_id
where feed_follows.user_id = $1;