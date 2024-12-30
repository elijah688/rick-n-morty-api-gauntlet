import { Episode } from "./episode.ts";
import { Location } from "./location.ts";

export interface GetPageOfCharactersOptions {
  limit: number;
  offset: number;
}

export interface CharacterCrud {
  id?: number;
  name: string;
  status?: string;
  species?: string;
  type?: string;
  gender?: string;
  origin_id?: number;
  location_id?: number;
  image?: string;
  url?: string;
  created?: string;
}

export interface Character {
  id?: number;
  name: string;
  status?: string;
  species?: string;
  type?: string;
  gender?: string;
  origin?: Location;
  location?: Location;
  image?: string;
  url?: string;
  created?: string;
  episodes?: number[];
  debut?: Episode;
}
