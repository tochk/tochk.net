CREATE TABLE IF NOT EXISTS articles
(
    id         SERIAL PRIMARY KEY NOT NULL,
    title      TEXT               NOT NULL,
    short_text TEXT               NOT NULL,
    text       TEXT               NOT NULL,
    language   TEXT               NOT NULL DEFAULT 'ru',
    created_at TIMESTAMP          NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP          NOT NULL DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS articles_tags
(
    tag_id     INTEGER NOT NULL,
    article_id INTEGER NOT NULL,
    PRIMARY KEY (tag_id, article_id)
);

CREATE TABLE IF NOT EXISTS articles_team_members
(
    team_member_id INTEGER NOT NULL,
    article_id     INTEGER NOT NULL,
    PRIMARY KEY (team_member_id, article_id)
);
