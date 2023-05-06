-- +goose Up
alter table note add column updated_at timestamp;

-- +goose Down
alter table note drop column updated_at;
