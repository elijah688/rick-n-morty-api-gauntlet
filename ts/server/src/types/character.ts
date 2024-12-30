
export interface CharactersPageParams {
  limit: number;
  offset: number;
}

export interface CharacterByIDParams {
  id: number;
}


export interface CharacterByIDResponse {
  id: number;
}


export interface Character {
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

export interface DeleteCharacterResponse {
  id:number
}