from fastapi import APIRouter, Query,Body
from services.services import Services
from model.character import CharacterResponse
from typing import List
from model.character import CharactersDebuteRequest
from model.character import CharactersEpisodesRequest
class CharacterRouter:
    def __init__(self, services: Services):
        self.__services = services
        self.router = APIRouter()  
        
        @self.router.get("/ping")
        async def ping():
            return await self.__services.ping()


        @self.router.get("/total")
        async def get_total():
            return await self.__services.get_total()
      
        @self.router.get("/{id}")
        async def get_character_by_id(id: int):
            return await self.__services.get_character_by_id(id)

        @self.router.get("")
        async def get_characters(limit: int = Query(10), offset: int = Query(0)):
            return await self.__services.get_characters(limit=limit, offset=offset)

        @self.router.post("")
        async def upsert_character(character:CharacterResponse):
            return await self.__services.upsert_character(character)


        @self.router.post("/list/debut")
        async def get_chars_debut(req:CharactersDebuteRequest):
            return await self.__services.get_chars_debut(req.ids)
        
        @self.router.post("/list/episodes")
        async def get_episodes_by_ids(req:CharactersEpisodesRequest):
            return await self.__services.get_episodes_by_ids(req.ids)
        
        @self.router.delete("/{id}")
        async def delete_character(id:int):
            await self.__services.delete_character(id)
            return {"id":id}
      
       