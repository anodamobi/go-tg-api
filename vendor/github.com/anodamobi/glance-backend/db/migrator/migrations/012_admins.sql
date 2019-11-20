-- +migrate Up

CREATE TABLE IF NOT EXISTS admins
(
    id                       BIGSERIAL PRIMARY KEY,
    first_name               VARCHAR(255) NOT NULL,
    last_name                VARCHAR(255) NOT NULL,
    email                    VARCHAR(255) NOT NULL,
    image                    VARCHAR(255) NOT NULL,
    password                 VARCHAR(255) NOT NULL,
    site_permissions         TEXT[]       NOT NULL,
    admin_permissions        TEXT[]       NOT NULL,
    resident_permissions     TEXT[]       NOT NULL,
    chat_permissions         TEXT[]       NOT NULL,
    delivery_permissions     TEXT[]       NOT NULL,
    maintenance_permissions  TEXT[]       NOT NULL,
    wellbeing_permissions    TEXT[]       NOT NULL,
    events_permissions       TEXT[]       NOT NULL,
    notification_permissions TEXT[]       NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS admins;