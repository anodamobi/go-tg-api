-- +migrate Up

CREATE TABLE IF NOT EXISTS awaiting_parcels
(
    pid       BIGINT PRIMARY KEY REFERENCES parcels (id),
    take_time TIMESTAMP NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS awaiting_parcels;