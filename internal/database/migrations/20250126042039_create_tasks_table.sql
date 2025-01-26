-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks
(
    id          INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    uuid        uuid                     DEFAULT gen_random_uuid(),
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    status      INTEGER                  DEFAULT 0,
    user_id     INTEGER      NOT NULL REFERENCES users (id),
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE 'UTC'),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE 'UTC')
);

CREATE
OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at
= (now() AT TIME ZONE 'UTC');
RETURN NEW;
END;
$$
language 'plpgsql';

CREATE TRIGGER update_tasks_updated_at
    BEFORE UPDATE
    ON tasks
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
DROP FUNCTION update_updated_at_column();
DROP
EVENT TRIGGER update_tasks_updated_at;
-- +goose StatementEnd