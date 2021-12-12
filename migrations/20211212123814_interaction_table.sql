-- +goose Up
-- +goose StatementBegin
CREATE TABLE interaction (
    id serial not null primary key,
    interaction_status dog_action,
    created_at timestamp default now(),
    is_deleted bool not null default false
);
-- +goose StatementEnd
