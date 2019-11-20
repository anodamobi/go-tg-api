-- +migrate Up

CREATE TABLE IF NOT EXISTS wellbeing_requests
(
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT REFERENCES users (id),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_resolved BOOLEAN   DEFAULT FALSE
);

-- +migrate Down

DROP TABLE IF EXISTS wellbeing_requests;