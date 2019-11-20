-- +migrate Up

CREATE TABLE IF NOT EXISTS events
(
    id               BIGSERIAL PRIMARY KEY,
    title            VARCHAR(50) NOT NULL,
    date             TIMESTAMP   NOT NULL,
    location         VARCHAR(50) NOT NULL,
    category         INT         NOT NULL,
    type             INT         NOT NULL,
    open_to          VARCHAR(50) NOT NULL,
    max_attendees    INT         NOT NULL,
    is_waitlist_open BOOLEAN DEFAULT FALSE,
    description      VARCHAR(10000)
);

-- +migrate Down

DROP TABLE IF EXISTS events;