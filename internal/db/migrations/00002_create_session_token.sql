-- +goose Up
create table if not exists user_sessions (
    id int primary key generated always as identity,
    request_id varchar(80),
    created_at timestamp,
    expires_at timestamp,
    session_token varchar(80),
    user_id int,
    foreign key (user_id) references users(id)
);

-- +goose Down
drop table user_sessions;
