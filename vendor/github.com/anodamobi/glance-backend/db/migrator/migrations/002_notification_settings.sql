-- +migrate Up

CREATE TABLE IF NOT EXISTS notification_settings
(
    user_id                    BIGINT NOT NULL,
    all_allowed                BOOLEAN DEFAULT TRUE,
    parcels_allowed            BOOLEAN DEFAULT TRUE,
    events_allowed             BOOLEAN DEFAULT TRUE,
    building_bulletins_allowed BOOLEAN DEFAULT TRUE,
    maintenance_allowed        BOOLEAN DEFAULT TRUE
);

-- +migrate Down

DROP TABLE IF EXISTS notification_settings;