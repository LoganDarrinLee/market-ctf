-- name: ListUserAccessLogs :many
select * from user_access_logs
where user_id = $1;

-- name: LastLoggedIn :one
select * from user_access_logs 
where user_id = $1 
order by logged_in desc
limit 1;

-- create does not return an acces log.
-- name: CreateAccessLog :exec
insert into user_access_logs (
    user_id, 
    logged_in,
    logged_out
) values (
    $1, $2, $3
);

-- name: DeleteAccessLog :exec
delete from user_access_logs where id = $1;
