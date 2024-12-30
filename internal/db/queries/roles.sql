-- name: GetRoleWithID :one
select * from user_roles
where user_role = $1 limit 1;

-- name: GetRoleWithName :one
select * from user_roles
where user_role = $1 limit 1;

-- name: ListRoles :many
select * from user_roles order by id;

-- name: CreateRole :one
insert into user_roles (
    user_role, role_info
) values (
    $1, $2
) returning *;

-- name: UpdateRole :one
update user_roles
    set user_role = $2,
        role_info = $3
    where id = $1
returning *;

-- name: DeleteRole :exec
delete from user_roles where id = $1;
