-- +goose Up
-- +goose StatementBegin
CREATE TYPE dog_action AS ENUM ('POPIS', 'WALK');
-- +goose StatementEnd
