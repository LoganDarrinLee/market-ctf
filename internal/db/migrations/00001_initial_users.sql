-- +goose Up

-- Ex: Vendor, customer, admin team.
create table if not exists user_roles (
    id int primary key generated always as identity,
    user_role varchar(50) not null unique,
    role_info varchar(500) null
);

create table if not exists users (
    id int primary key generated always as identity,
    role_id int,
    private_username varchar(35) not null unique,
    public_username varchar(35) not null unique,
    password_hash varchar(80) not null,
    session_token varchar(100) null,
    foreign key (role_id) references user_roles(id)
);

-- Log user access data to help with deanonymization
create table if not exists user_access_logs (
    id int primary key generated always as identity,
    user_id int,
    logged_in timestamp,
    -- Will not always be populated, they may just close the browser.
    logged_out timestamp null, 
    foreign key (user_id) references users(id)
);


-- +goose Down
drop table users;
drop table user_access_logs;
drop table user_account_changes;