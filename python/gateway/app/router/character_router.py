from fastapi import APIRouter, Query,Body
from services.services import Services
from model.character import CharacterResponse,CharactersResponse,CharacterInput
class CharacterRouter:
    def __init__(self, services: Services):
        self.__router = APIRouter()
        self.__services = services


        @self.__router.post("")
        async def upsert_character(character =Body()) -> CharacterResponse:
          return  await self.__services.upsert_character(character)

     
        @self.__router.get("")
        async def get_characters(limit: int = Query(10), offset: int = Query(0)) -> CharactersResponse:
          return await self.__services.get_characters(limit, offset)
     
        @self.__router.get("/{id}")
        async def get_character_by_id(id) -> CharacterResponse:
          return await self.__services.get_character_by_id(id)

        @self.__router.delete("/{id}")
        async def delete_character(id) -> dict[str,int]:
          return await self.__services.delete_character(id)

    def get_router(self):
        return self.__router

