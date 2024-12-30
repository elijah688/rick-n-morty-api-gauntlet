import { Character } from "../types/characters.ts";
import { CharacterCrud } from "../types/characters.ts";
import { Episode } from "../types/episode.ts";
import { Location } from "../types/location.ts";

function mapCharacterToFullDetails(
  chars: CharacterCrud[],
  episodes: Record<number, number[]>,
  locations: Record<number, Location>,
  debutes: Record<number, Episode>
): Character[] {
  return chars.map((x) => ({
    id: x.id,
    name: x.name,
    status: x.status,
    species: x.species,
    type: x.type,
    gender: x.gender,
    created: x.created,
    image: x.image,
    url: x.url,
    episodes: x.id ? episodes[x.id] || [] : [],
    debut: x.id ? debutes[x.id] : undefined,
    origin: x.origin_id ? locations[x.origin_id] : undefined,
    location: x.location_id ? locations[x.location_id] : undefined,
  }));
}

export async function compileCharacters(
  chars: CharacterCrud[],
  getEpisodes: (params: { ids: number[] }) => Promise<Record<number, number[]>>,
  getLocations: (params: {
    ids: number[];
  }) => Promise<Record<number, Location>>,
  getDebutes: (params: { ids: number[] }) => Promise<Record<number, Episode>>
): Promise<Character[]> {
  const ids: number[] = chars
    .map(({ id }) => id)
    .filter((id): id is number => !!id);

  const locIDs = Array.from(
    new Set(
      chars
        .map(({ origin_id, location_id }) => [origin_id, location_id])
        .flat()
        .filter((id): id is number => !!id)
    )
  );

  const [episodes, locations, debutes] = await Promise.all([
    getEpisodes({ ids }),
    getLocations({ ids: locIDs }),
    getDebutes({ ids }),
  ]);

  return mapCharacterToFullDetails(chars, episodes, locations, debutes);
}
