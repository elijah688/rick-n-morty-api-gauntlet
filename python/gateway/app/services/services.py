
from config.config import Config
from services.character import CharacterService
from services.location import LocationService
from services.episode import EpisodeService
from model.character import CharacterClientResponse, CharacterResponse,CharacterInput,CharactersResponse
from model.episode import Episode
from model.location import Location
from typing import List,Tuple
import asyncio

def map_related_data(character, debuts, episodes, locations):
    return {
        "debut": debuts.get(character.id),
        "episodes": episodes.get(character.id),
        "origin_name": locations.get(character.origin_id) if character.origin_id is not None else None,
        "location_name": locations.get(character.location_id) if character.location_id is not None else None,
    }

class Services:
    def __init__(self, config:Config,):
        self.__character_service = CharacterService(config)
        self.__location_service = LocationService(config)
        self.__episode_service = EpisodeService(config)

    async def __fetch_related_data(self,chars: List[CharacterClientResponse]) -> Tuple[dict[int,Episode],dict[int,List[int]],dict[int,Location]]:
        ids = [x.id for x in chars]
        loc_ids = list({id for c in chars for id in (c.origin_id, c.location_id)})

        return await asyncio.gather(
            self.__episode_service.get_characters_debut_by_ids(ids),
            self.__episode_service.get_episodes_by_ids(ids),
            self.__location_service.get_locations_by_ids(loc_ids),
        )
        

    async def get_characters(self, limit: int, offset: int) -> CharactersResponse:
       

        chars, total =  await asyncio.gather(
            self.__character_service.get_characters(limit, offset),
            self.__character_service.get_total()
        )
        debuts, episodes, locations = await self.__fetch_related_data(chars)


        res_chars= []
        for c in chars:
            res_chars.append(CharacterResponse(
                id=c.id,
                name=c.name,
                status=c.status,
                species=c.species,
                type=c.type,
                gender=c.gender,
                image=c.image,
                url=c.url,
                created=c.created,
                **map_related_data(c, debuts, episodes, locations),
            ))

            print(res_chars)
        return CharactersResponse(
            total=total,
            characters= res_chars
        )
    
    
       
    

    async def get_character_by_id(self, id: int) -> CharacterResponse:
        char = await self.__character_service.get_character_by_id(id)

        debuts, episodes, locations = await self.__fetch_related_data([char])
        
        return CharacterResponse(
                id=char.id,
                name=char.name,
                status=char.status,
                species=char.species,
                type=char.type,
                gender=char.gender,
                image=char.image,
                url=char.url,
                created=char.created,
                **map_related_data(char, debuts, episodes, locations),

            )
    

    async def upsert_character(self, character: CharacterInput) -> CharacterResponse:
        return await self.__character_service.upsert_character(character)
    
    async def delete_character(self, id: int) -> dict[str,int]:
        return await self.__character_service.delete_character(id)
    

    