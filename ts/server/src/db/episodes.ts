import {
  CharacterEpisode,
  GetCharactersEpisodesByIDs,
  Episode,
  DebutsResponse,
  DebutIDs
} from "../types/episodes.ts";
import { QuerySignature } from "../types/queryFunc.ts";

export async function getEpisodesByIDs(
  queryFunc: QuerySignature,
  { ids }: GetCharactersEpisodesByIDs
): Promise<CharacterEpisode[]> {
  const paramSymbols = ids.map((_, i) => `$${i + 1}`).join(", ");
  const query = `
    SELECT * FROM character_episode
    WHERE character_id in (${paramSymbols})
  `;

  return await queryFunc<CharacterEpisode>(query, ids);
}

export async function getFirstAppearancesByCharIDs(
  queryFunc: QuerySignature,
  { ids: characterIDs }: DebutIDs
): Promise<DebutsResponse> {
  const paramSymbols = characterIDs.map((_, i) => `$${i + 1}`).join(", ");

  const query = `
    WITH ranked_episodes AS (
      SELECT 
        ce.character_id,
        ce.episode_id,
        e.name AS episode_name,
        e.air_date,
        e.episode_code,
        e.url AS episode_url,
        e.created AS episode_created,
        ROW_NUMBER() OVER (PARTITION BY ce.character_id ORDER BY ce.episode_id ASC) AS rn
      FROM character_episode ce
      INNER JOIN episode e
      ON ce.episode_id = e.id
      WHERE ce.character_id IN (${paramSymbols})
    )
    SELECT 
      character_id,
      episode_id AS id,
      episode_name AS name,
      air_date,
      episode_code,
      episode_url AS url,
      episode_created AS created
    FROM ranked_episodes
    WHERE rn = 1
    ORDER BY character_id ASC;
  `;

  const rows = await queryFunc<{ character_id: number } & Episode>(
    query,
    characterIDs
  );

  const characterEpisodes = new Map<number, Episode>();

  for (const { character_id, ...episode } of rows) {
    characterEpisodes.set(character_id, episode);
  }

  

  return <DebutsResponse>Object.fromEntries(characterEpisodes);
}
