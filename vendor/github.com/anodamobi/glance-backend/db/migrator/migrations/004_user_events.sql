-- +migrate Up

CREATE TABLE IF NOT EXISTS user_events
(
    id       BIGSERIAL PRIMARY KEY,
    event_id BIGINT REFERENCES events (id),
    user_id  BIGINT REFERENCES users (id),
    is_saved BOOLEAN DEFAULT FALSE,
    status   INT NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS user_events;