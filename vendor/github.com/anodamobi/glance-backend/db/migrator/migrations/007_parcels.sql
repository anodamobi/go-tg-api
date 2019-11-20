-- +migrate Up

CREATE TABLE IF NOT EXISTS parcels
(
    id                 BIGSERIAL PRIMARY KEY,
    user_id            BIGINT REFERENCES users (id),
    room               INT           NOT NULL,
    delivery_company   VARCHAR(100)  NOT NULL,
    parcel_location    VARCHAR(100)  NOT NULL,
    notification_title VARCHAR(100)  NOT NULL,
    notification_text  VARCHAR(1000) NOT NULL,
    created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status             INT       DEFAULT 0
);

-- +migrate Down

DROP TABLE IF EXISTS parcels;