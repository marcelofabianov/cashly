-- +goose Up
-- +goose StatementBegin
CREATE TABLE identity_users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    public_id UUID NOT NULL,
    identity_document VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    version BIGINT NOT NULL DEFAULT 1
);

-- Índices
CREATE INDEX idx_identity_users_public_id ON identity_users (public_id);
CREATE INDEX idx_identity_users_identity_document ON identity_users (identity_document);
CREATE INDEX idx_identity_users_email ON identity_users (email);
CREATE INDEX idx_identity_users_enabled ON identity_users (enabled);
CREATE INDEX idx_identity_users_deleted_at ON identity_users (deleted_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_identity_users_public_id;
DROP INDEX IF EXISTS idx_identity_users_identity_document;
DROP INDEX IF EXISTS idx_identity_users_email;
DROP INDEX IF EXISTS idx_identity_users_enabled;
DROP INDEX IF EXISTS idx_identity_users_deleted_at;
DROP TABLE IF EXISTS identity_users;
-- +goose StatementEnd
