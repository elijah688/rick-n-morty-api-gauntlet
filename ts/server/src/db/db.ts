import { Pool } from "https://deno.land/x/postgres@v0.17.0/mod.ts";
import { QueryArguments } from "https://deno.land/x/postgres@v0.17.0/query/query.ts";
import { Config } from "../config/config.ts";
import { PingResult } from "../types/ping.ts";
import {
  Character,
  CharactersPageParams,
  CharacterByIDParams,
} from "../types/character.ts";

import {
  GetCharactersEpisodesByIDs,
  CharacterEpisodesResponse,
  DebutIDs,
  DebutsResponse
} from "../types/episodes.ts";
import {
  getCharacters,
  getCharacterByID,
  upsertCharacter,
  deleteCharacter,
  getTotal
} from "./character.ts";
import { ping } from "./ping.ts";
import { GetLocationsByIDs, Location } from "../types/location.ts";
import { getLocationsByIDs } from "./location.ts";
import { getEpisodesByIDs } from "./episodes.ts";

import { getFirstAppearancesByCharIDs} from "./episodes.ts"
export class DB {
  private pool: Pool;

  constructor(config: Config) {
    this.pool = new Pool(
      {
        database: config.appDbName,
        hostname: config.mainDbHost,
        password: config.mainDbPass,
        port: config.mainDbPort,
        user: config.mainDbUser,
      },
      4,
      true
    );
  }

  async query<T>(
    queryString: string,
    queryArgs: QueryArguments = {}
  ): Promise<T[]> {
    const client = await this.pool.connect();
    try {
      const result = await client.queryObject<T>(queryString, queryArgs);
      return result.rows;
    } finally {
      client?.release();
    }
  }

  async ping(): Promise<PingResult> {
    return await ping(this.query.bind(this));
  }

  async GetCharacters(params: CharactersPageParams): Promise<Character[]> {
    return await getCharacters(this.query.bind(this), params);
  }

  async GetCharacterByID(params: CharacterByIDParams): Promise<Character|null> {
    const rows =await getCharacterByID(this.query.bind(this), params);
    if ((rows.length)>0){
      return rows[0]
    }
    return null

  }

  async UpsertCharacter(params: Character): Promise<Character> {
    const rows = await upsertCharacter(this.query.bind(this), params);
    if (rows.length > 0) {
      return rows[0];
    }

    throw Error("failed upserting character");
  }

  async DeleteCharacterByID(params: CharacterByIDParams): Promise<void> {
    return await deleteCharacter(this.pool, params);
  }

  async GetLocationsByIDs(params: GetLocationsByIDs): Promise<Location[]> {
    return await getLocationsByIDs(this.query.bind(this), params);
  }

  async GetEpisodesByCharIDs(
    params: GetCharactersEpisodesByIDs
  ): Promise<CharacterEpisodesResponse> {
    const eps = await getEpisodesByIDs(this.query.bind(this), params);

    const res: CharacterEpisodesResponse = {};
    for (const ce of eps) {
      if (res[ce.character_id]) {
        res[ce.character_id].push(ce.episode_id);
      } else{
        res[ce.character_id] = [ce.episode_id];
      }
    }
    return res;
  }

  async GetDebuts(
    params: DebutIDs
  ): Promise<DebutsResponse> {
    return await getFirstAppearancesByCharIDs(this.query.bind(this), params);
  }

  async GetTotal(
  ): Promise<number> {
    return await getTotal(this.query.bind(this));
  }


  
}

export default DB;
