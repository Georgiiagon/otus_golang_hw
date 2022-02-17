-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    id serial primary key,
    title varchar,
    duration varchar,
    description varchar,
    notify_at timestamp,
    created_at timestamp,
    updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
