from db.db import Database
from model.location import LocationResponse
from model.character import CharacterResponse 
from model.episode import EpisodeResponse
from typing import List
import datetime
class Services:
    def __init__(self, db: Database):
        self.__db = db  

    async def ping(self):
        return await self.__db.ping()
    

    async def get_characters(self, limit:int, offset:int):
        res = await self.__db.get_characters(limit,offset)
        return [CharacterResponse.model_validate(character_sql) for character_sql in res]
    

    async def get_character_by_id(self,id):
        return CharacterResponse.model_validate(await self.__db.get_character_by_id(id))
    
    async def upsert_character(self, character: CharacterResponse):
        if character.created==None:
            character.created= datetime.datetime.now(datetime.timezone.utc).replace(tzinfo=None)
        return CharacterResponse.model_validate(await self.__db.upsert_character(character.to_sql_model()))

    async def get_chars_debut(self, ids: List[int])->List[dict[int,EpisodeResponse]]:
        return { row[0]: EpisodeResponse(
                        id=row[1],
                        name=row[2],
                        air_date=row[3],
                        episode_code=row[4],
                        url=row[5],
                        created=row[6]
                    ) for row in await self.__db.get_chars_debut(ids) 
                }
            

    async def get_episodes_by_ids(self, ids: List[int]):
        return  await self.__db.get_episodes_by_ids(ids)
    
    async def get_locations_by_ids(self, ids: List[int]):
        return  {k:LocationResponse.model_validate(v) for k,v in (await self.__db.get_locations_by_ids(ids)).items()}
             
             
    async def delete_character(self,id):
        return await self.__db.delete_character(id)
   
    async def get_total(self):
        return await self.__db.get_total()
        
    