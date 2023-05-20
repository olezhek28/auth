-- +goose Up
create table users
(
    username text not null,
    password text not null,
    unique (username)
);

-- +goose Down
drop table users;
