-- +migrate Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
  id uuid DEFAULT uuid_generate_v4 (),
  external_id VARCHAR(100) NOT NULL,
  name TEXT NOT NULL,
  -- TODO: change according to IETF language tag standard
  language VARCHAR(10) NOT NULL,
  joined_at timestamp default NOW(),
  updated_at timestamp,
  PRIMARY KEY (id)
);

-- +migrate Down

DROP TABLE users;