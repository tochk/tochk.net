CREATE TABLE IF NOT EXISTS tags
(
    id       SERIAL PRIMARY KEY NOT NULL,
    name     TEXT               NOT NULL,
    language TEXT               NOT NULL
);

CREATE TABLE IF NOT EXISTS team_members
(
    id       SERIAL PRIMARY KEY NOT NULL,
    name     TEXT               NOT NULL,
    language TEXT               NOT NULL
);

CREATE TABLE IF NOT EXISTS projects_tags
(
    tag_id     INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    PRIMARY KEY (tag_id, project_id)
);

CREATE TABLE IF NOT EXISTS projects_team_members
(
    team_member_id INTEGER NOT NULL,
    project_id     INTEGER NOT NULL,
    PRIMARY KEY (team_member_id, project_id)
);
