-- +migrate Up

CREATE TABLE IF NOT EXISTS maintenance
(
    id          BIGSERIAL PRIMARY KEY,
    images      VARCHAR(100)[] NOT NULL,
    description VARCHAR(300)   NOT NULL,
    user_id     BIGINT REFERENCES users (id),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    unique_id   VARCHAR(100)   NOT NULL UNIQUE
);

-- +migrate Down

DROP TABLE IF EXISTS maintenance;