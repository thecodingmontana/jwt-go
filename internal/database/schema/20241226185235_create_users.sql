-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id TEXT PRIMARY KEY,
    name TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
