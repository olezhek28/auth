-- +goose Up
alter table users add column role text not null default 'user';

-- +goose Down
alter table users drop column role;
