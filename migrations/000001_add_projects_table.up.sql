CREATE TABLE IF NOT EXISTS projects
(
    id                SERIAL PRIMARY KEY NOT NULL,
    title             TEXT               NOT NULL,
    short_description TEXT               NOT NULL,
    description       TEXT               NOT NULL,
    image_url         TEXT               NOT NULL,
    language          TEXT               NOT NULL DEFAULT 'ru',
    created_at        TIMESTAMP          NOT NULL DEFAULT current_timestamp,
    updated_at        TIMESTAMP          NOT NULL DEFAULT current_timestamp
);