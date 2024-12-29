// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
insert into users (
    role_id, 
    private_username, 
    public_username, 
    password_hash, 
    session_token
) values (
    $1, $2, $3, $4, $5
) returning id, role_id, private_username, public_username, password_hash, session_token
`

type CreateUserParams struct {
	RoleID          pgtype.Int4
	PrivateUsername string
	PublicUsername  string
	PasswordHash    string
	SessionToken    pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.RoleID,
		arg.PrivateUsername,
		arg.PublicUsername,
		arg.PasswordHash,
		arg.SessionToken,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.PrivateUsername,
		&i.PublicUsername,
		&i.PasswordHash,
		&i.SessionToken,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
delete from users where id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getByID = `-- name: GetByID :one
select id, role_id, private_username, public_username, password_hash, session_token from users
where id = $1 limit 1
`

func (q *Queries) GetByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.PrivateUsername,
		&i.PublicUsername,
		&i.PasswordHash,
		&i.SessionToken,
	)
	return i, err
}

const getBySessionToken = `-- name: GetBySessionToken :one
select id, role_id, private_username, public_username, password_hash, session_token from users 
where session_token = $1 limit 1
`

func (q *Queries) GetBySessionToken(ctx context.Context, sessionToken pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getBySessionToken, sessionToken)
	var i User
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.PrivateUsername,
		&i.PublicUsername,
		&i.PasswordHash,
		&i.SessionToken,
	)
	return i, err
}

const getWithPrivateUsername = `-- name: GetWithPrivateUsername :one
select id, role_id, private_username, public_username, password_hash, session_token from users
where private_username = $1 limit 1
`

func (q *Queries) GetWithPrivateUsername(ctx context.Context, privateUsername string) (User, error) {
	row := q.db.QueryRow(ctx, getWithPrivateUsername, privateUsername)
	var i User
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.PrivateUsername,
		&i.PublicUsername,
		&i.PasswordHash,
		&i.SessionToken,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
select id, role_id, private_username, public_username, password_hash, session_token from users order by id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.RoleID,
			&i.PrivateUsername,
			&i.PublicUsername,
			&i.PasswordHash,
			&i.SessionToken,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
update users
    set private_username = $2,
        public_username = $3,
        password_hash = $4, 
        session_token = $5
    where id = $1
returning id, role_id, private_username, public_username, password_hash, session_token
`

type UpdateUserParams struct {
	ID              int32
	PrivateUsername string
	PublicUsername  string
	PasswordHash    string
	SessionToken    pgtype.Text
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.PrivateUsername,
		arg.PublicUsername,
		arg.PasswordHash,
		arg.SessionToken,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.PrivateUsername,
		&i.PublicUsername,
		&i.PasswordHash,
		&i.SessionToken,
	)
	return i, err
}
