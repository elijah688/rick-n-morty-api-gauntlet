CREATE SCHEMA IF NOT EXISTS ricknmorty;

CREATE TABLE IF NOT EXISTS ricknmorty.character (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    status TEXT,
    species TEXT,
    type TEXT,
    gender TEXT,
    origin_id INT,
    location_id INT,
    image TEXT,
    url TEXT,
    created TIMESTAMP,
    FOREIGN KEY (origin_id) REFERENCES location (id),
    FOREIGN KEY (location_id) REFERENCES location (id)
);

CREATE TABLE IF NOT EXISTS ricknmorty.location (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT,
    dimension TEXT,
    url TEXT,
    created TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ricknmorty.episode (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    air_date TEXT,
    episode_code TEXT,
    url TEXT,
    created TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ricknmorty.character_episode (
    character_id INT NOT NULL,
    episode_id INT NOT NULL,
    PRIMARY KEY (character_id, episode_id),
    FOREIGN KEY (character_id) REFERENCES character (id),
    FOREIGN KEY (episode_id) REFERENCES episode (id)
);
