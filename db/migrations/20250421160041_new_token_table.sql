-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS token
(
    id SERIAL PRIMARY KEY,
    Refresh TEXT,
    UserID INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS token;
-- +goose StatementEnd
