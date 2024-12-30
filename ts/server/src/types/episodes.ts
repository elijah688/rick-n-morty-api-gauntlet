export interface GetCharactersEpisodesByIDs {
  ids: number[];
}

export interface Episode {
  id: number;
  name: string;
  air_date?: string;
  episode_code?: string;
  url?: string;
  created?: string;
}

export interface CharacterEpisodesResponse {
  [id: number]: number[];
}

export interface CharacterEpisode {
  character_id: number;
  episode_id: number;
}


export interface DebutIDs {
  ids: number[];
}

export interface DebutsResponse {
  [id: number]: Episode;
}
