-- name: GetWithPrivateUsername :one
select * from users
where private_username = $1 limit 1;

-- name: GetBySessionToken :one
select 
    user_sessions.session_token,
    user_sessions.created_at,
    user_sessions.expires_at,
    users.id
from user_sessions
join users
on user_sessions.user_id = users.id;


-- name: GetByID :one
select * from users
where id = $1 limit 1;

-- name: ListUsers :many
select * from users order by id;

-- name: CreateUser :one
insert into users (
    role_id, 
    private_username, 
    public_username, 
    password_hash, 
    session_token
) values (
    $1, $2, $3, $4, $5
) returning *;

-- name: UpdateUser :one
update users
    set private_username = $2,
        public_username = $3,
        password_hash = $4, 
        session_token = $5
    where id = $1
returning *;

-- name: DeleteUser :exec
delete from users where id = $1;
