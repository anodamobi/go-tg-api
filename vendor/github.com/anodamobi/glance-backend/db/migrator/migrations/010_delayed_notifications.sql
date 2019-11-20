-- +migrate Up

CREATE TABLE IF NOT EXISTS delayed_notifications
(
    id             BIGSERIAL PRIMARY KEY,
    -- if event_id == 0, this record doesn't relate to event notifications
    event_id       BIGINT        NOT NULL,
    user_device_id VARCHAR(200)  NOT NULL,
    time_to_notify TIMESTAMP     NOT NULL,
    comment        VARCHAR(1000) NOT NULL,
    title          VARCHAR(200)  NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS delayed_notifications;