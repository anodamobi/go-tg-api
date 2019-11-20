-- +migrate Up

CREATE TABLE IF NOT EXISTS notifications
(
    id         BIGSERIAL PRIMARY KEY,
    type       VARCHAR(50)   NOT NULL,
    title      VARCHAR(50)   NOT NULL,
    body       VARCHAR(1000) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    sender     BIGINT        NOT NULL REFERENCES users (id),
    continent  VARCHAR(20)   NOT NULL,
    country    VARCHAR(50)   NOT NULL,
    city       VARCHAR(20)   NOT NULL,
    location   VARCHAR(20)   NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS notifications;