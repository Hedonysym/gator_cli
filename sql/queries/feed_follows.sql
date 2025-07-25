-- name: CreateFeedFollow :one
with inserted_feed_follow as (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT 
inserted_feed_follow.*,
feeds.name AS feed_name,
users.name AS user_name
FROM inserted_feed_follow
inner join feeds on feeds.id = inserted_feed_follow.feed_id
inner join users on users.id = inserted_feed_follow.user_id;

-- name: GetFeedFollowsForUser :many
select feeds.name as feed_name, feed_follows.*
from feed_follows
inner join feeds on feeds.id = feed_follows.feed_id
where feed_follows.user_id = $1
order by feed_follows.created_at desc;

-- name: FeedUnfollow :exec
delete from feed_follows
where user_id = $1 and feed_id = $2;