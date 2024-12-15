
SELECT setval('location_id_seq', COALESCE((SELECT MAX(id) + 1 FROM location), 1));
SELECT setval('character_id_seq', COALESCE((SELECT MAX(id) + 1 FROM character), 1));
SELECT setval('episode_id_seq', COALESCE((SELECT MAX(id) + 1 FROM episode), 1));
