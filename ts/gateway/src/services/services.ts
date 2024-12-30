import { Config } from "../config/config.ts";
import {
  GetPageOfCharactersOptions,
  CharacterCrud,
  Character,
} from "../types/characters.ts";
import { Episode } from "../types/episode.ts";
import { Location, validateLocation } from "../types/location.ts";
import { compileCharacters } from "../utils/utils.ts";
export class Services {
  private config: Config;
  private apiUrl: string;

  constructor(config: Config) {
    this.config = config;
    this.apiUrl = config.crudServicsHost;
  }

  private async getDebutes({
    ids,
  }: {
    ids: number[];
  }): Promise<Record<number, Episode>> {
    const locationsRes = await fetch(`${this.apiUrl}/character/list/debut`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ids }),
    });

    if (!locationsRes.ok) {
      throw new Error(`Failed to fetch episodes: ${locationsRes.statusText}`);
    }

    return await locationsRes.json();
  }
  private async getLocations({
    ids,
  }: {
    ids: number[];
  }): Promise<Record<number, Location>> {
    if (ids.length == 0) {
      return {};
    }
    const locationsRes = await fetch(`${this.apiUrl}/location/list`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ids }),
    });

    if (!locationsRes.ok) {
      throw new Error(`Failed to fetch episodes: ${locationsRes.statusText}`);
    }

    const locs: Location[] = await locationsRes.json();
    const res: { [id: string]: Location } = {};
    locs.forEach((loc) => (res[loc.id] = loc));

    return res;
  }
  private async getEpisodes({
    ids,
  }: {
    ids: number[];
  }): Promise<Record<number, number[]>> {
    const episodesRes = await fetch(`${this.apiUrl}/character/list/episodes`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ids }),
    });

    if (!episodesRes.ok) {
      throw new Error(`Failed to fetch episodes: ${episodesRes.statusText}`);
    }

    return await episodesRes.json();
  }
  async getCharacters(ops: GetPageOfCharactersOptions): Promise<Character[]> {
    const charsRes = await fetch(
      `${this.apiUrl}/character?limit=${ops.limit}&offset=${ops.offset}`
    );

    if (!charsRes.ok) {
      throw new Error("Failed to fetch characters");
    }

    const chars: CharacterCrud[] = await charsRes.json();

    return await compileCharacters(
      chars,
      this.getEpisodes.bind(this),
      this.getLocations.bind(this),
      this.getDebutes.bind(this)
    );
  }

  async getCharacterByID(id: number): Promise<Character | null> {
    const charRes = await fetch(`${this.apiUrl}/character/${id}`);
    if (!charRes.ok) {
      throw new Error("Failed to fetch characters");
    }

    if (charRes.body) {
      const char: CharacterCrud = await charRes.json();

      const rs = await compileCharacters(
        [char],
        this.getEpisodes.bind(this),
        this.getLocations.bind(this),
        this.getDebutes.bind(this)
      );

      if (rs.length > 0) {
        return rs[0];
      }
    }
    return null;
  }
  async upsertCharacter(char: Character): Promise<Character | null> {
    const { origin, location } = char;
    if (!validateLocation(origin)) {
      char.origin = undefined;
    }

    if (!validateLocation(location)) {
      char.location = undefined;
    }
    const upsertRes = await fetch(`${this.apiUrl}/character`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(char),
    });

    if (!upsertRes.ok) {
      throw new Error(
        `Failed to upsert character episodes: ${upsertRes.statusText}`
      );
    }

    if (upsertRes.body) {
      const char = await upsertRes.json();

      const rs = await compileCharacters(
        [char],
        this.getEpisodes.bind(this),
        this.getLocations.bind(this),
        this.getDebutes.bind(this)
      );

      if (rs.length > 0) {
        return rs[0];
      }
    }
    return null;
  }

  async deleteCharacter(id: number): Promise<{ id: number }> {
    const res = await fetch(`${this.apiUrl}/character/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      throw new Error(`Failed to fetch episodes: ${res.statusText}`);
    }

    return { id };
  }

  async total(): Promise<number> {
    const res = await fetch(`${this.apiUrl}/total`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      throw new Error(`Failed to get total: ${res.statusText}`);
    }

    const t: { total: number } = await res.json();
    return t.total;
  }
}
