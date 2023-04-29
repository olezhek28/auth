-- +goose Up
create table note
(
    id         bigserial primary key,
    title      text      not null,
    content    text      not null,
    created_at timestamp not null
);

-- +goose Down
drop table note;
