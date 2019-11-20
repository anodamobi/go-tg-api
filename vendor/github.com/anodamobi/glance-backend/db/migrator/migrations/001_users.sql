-- +migrate Up

CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL PRIMARY KEY,
    image         VARCHAR(200),
    first_name    VARCHAR(50)  NOT NULL,
    last_name     VARCHAR(50)  NOT NULL,
    given_name    VARCHAR(50),
    room          INT          NOT NULL,
    email         VARCHAR(50)  NOT NULL UNIQUE,
    password      VARCHAR(100) NOT NULL,
    phone         VARCHAR(50),
    continent     VARCHAR(50)  NOT NULL,
    country       VARCHAR(50)  NOT NULL,
    city          VARCHAR(50)  NOT NULL,
    site          VARCHAR(50)  NOT NULL,
    block         VARCHAR(50)  NOT NULL,
    floor         INT          NOT NULL,
    temp_pass     VARCHAR(50)  NOT NULL,
    is_first_time BOOLEAN DEFAULT TRUE,
    device_id     VARCHAR(100) NOT NULL
);

INSERT INTO users (image, first_name, last_name, given_name, room, email, password, phone, continent, country, city,
                   site, block, floor, temp_pass, device_id)
VALUES ('https://images.unsplash.com/photo-1564889956728-9f3e4ad65498?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2134&q=80',
        'John', 'Doe', '', 312, 'test@email.com', '$2a$10$fQ.bzD8kSaTc8O6BPXfPe.ECuOV23dc5u272PKjol5/7GckYndocy', '',
        'Europe', 'UK', 'London', '6F', '12B', 5, 'Password1', '');

-- +migrate Down

DROP TABLE IF EXISTS users;