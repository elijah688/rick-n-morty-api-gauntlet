
CREATE SEQUENCE if not exists location_id_seq START 0 MINVALUE 0;
CREATE SEQUENCE if not exists character_id_seq START 0 MINVALUE 0;
CREATE SEQUENCE if not exists episode_id_seq START 0 MINVALUE 0;

-- Create location table with custom sequence
CREATE TABLE IF NOT EXISTS location (
    id INT PRIMARY KEY DEFAULT nextval('location_id_seq'),
    name TEXT NOT NULL,
    type TEXT,
    dimension TEXT,
    url TEXT,
    created TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_location_name ON location (name);

INSERT INTO location (id, name, type, dimension, url, created)
VALUES (0, 'unknown', NULL, NULL, NULL, NULL)
ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS character (
    id INT PRIMARY KEY DEFAULT nextval('character_id_seq'),
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

CREATE INDEX IF NOT EXISTS idx_character_name ON character (name);
CREATE INDEX IF NOT EXISTS idx_character_status ON character (status);
CREATE INDEX IF NOT EXISTS idx_character_location_id ON character (location_id);

-- Create episode table with custom sequence
CREATE TABLE IF NOT EXISTS episode (
    id INT PRIMARY KEY DEFAULT nextval('episode_id_seq'),
    name TEXT NOT NULL,
    air_date TEXT,
    episode_code TEXT,
    url TEXT,
    created TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_episode_name ON episode (name);
CREATE INDEX IF NOT EXISTS idx_episode_air_date ON episode (air_date);

-- Create character_episode table
CREATE TABLE IF NOT EXISTS character_episode (
    character_id INT NOT NULL,
    episode_id INT NOT NULL,
    PRIMARY KEY (character_id, episode_id),
    FOREIGN KEY (character_id) REFERENCES character (id),
    FOREIGN KEY (episode_id) REFERENCES episode (id)
);

CREATE INDEX IF NOT EXISTS idx_character_episode_character_id ON character_episode (character_id);
CREATE INDEX IF NOT EXISTS idx_character_episode_episode_id ON character_episode (episode_id);