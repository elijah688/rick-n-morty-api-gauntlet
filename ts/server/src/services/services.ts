import { DB } from "../db/db.ts";
import { Character, CharacterByIDParams } from "../types/character.ts";
import { CharactersPageParams } from "../types/character.ts";
import { DebutsResponse } from "../types/episodes.ts";
import { GetCharactersEpisodesByIDs } from "../types/episodes.ts";
import { CharacterEpisodesResponse } from "../types/episodes.ts";
import { DebutIDs } from "../types/episodes.ts";
import { GetLocationsByIDs, Location } from "../types/location.ts";
import { PingResult } from "../types/ping.ts";

export class Services {
  private db: DB;

  constructor(db: DB) {
    this.db = db;
  }

  async Ping(): Promise<PingResult> {
    return await this.db.ping()
  }
  async GetCharacters(params: CharactersPageParams): Promise<Character[]> {
    return await this.db.GetCharacters(params)
  }

  async GetCharacterByID(params: CharacterByIDParams): Promise<Character|null> {
    return await this.db.GetCharacterByID(params)
  }

  async UpsertCharacter(params: Character): Promise<Character> {
    return await this.db.UpsertCharacter(params)
  }

  async DeleteCharacterByID(params: CharacterByIDParams): Promise<void> {
    return await this.db.DeleteCharacterByID(params)
  }

  async GetLocationsByIDs(params: GetLocationsByIDs): Promise<Location[]> {
    return await this.db.GetLocationsByIDs(params)
  }

  async GetEpisodesByCharIDs(params: GetCharactersEpisodesByIDs): Promise<CharacterEpisodesResponse> {
    return await this.db.GetEpisodesByCharIDs(params)
  }


  async GetDebuts(params: DebutIDs): Promise<DebutsResponse> {
    return await this.db.GetDebuts(params)
  }

  async GetTotal(): Promise<number> {
    return await this.db.GetTotal()
  }


}

