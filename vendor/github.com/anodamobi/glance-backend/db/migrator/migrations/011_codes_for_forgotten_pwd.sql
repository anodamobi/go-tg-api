-- +migrate Up

CREATE TABLE IF NOT EXISTS codes_for_forgotten_pwd
(
    id    BIGSERIAL PRIMARY KEY,
    code  VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL

);

-- +migrate Down

DROP TABLE IF EXISTS codes_for_forgotten_pwd;