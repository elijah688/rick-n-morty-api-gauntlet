
from config.config import Config
import httpx
from pydantic import TypeAdapter
from model.character import CharacterClientResponse, CharacterInput
from typing import List
from fastapi.encoders import jsonable_encoder


class CharacterService:
    def __init__(self, config: Config):
        self.__config = config

    async def get_characters(self, limit, offset) -> List[CharacterClientResponse]:
        url = f"{self.__config.CRUD_SVC_HOST}/character?limit={limit}&offset={offset}"
        
        async with httpx.AsyncClient() as client:
            response = await client.get(url)

            if response.status_code == 200:
                try:
                    return TypeAdapter(List[CharacterClientResponse]).validate_python(response.json())
                except Exception as e:
                    return f"Error parsing response: {str(e)}"
            else:
                return f"Error: {response.status_code} - {response.text}"
            
    async def get_character_by_id(self, id:int) -> CharacterClientResponse:
        url = f"{self.__config.CRUD_SVC_HOST}/character/{id}"
        
        async with httpx.AsyncClient() as client:
            response = await client.get(url)

            if response.status_code == 200:
                try:
                    return TypeAdapter(CharacterClientResponse).validate_python(response.json())
                except Exception as e:
                    return f"Error parsing response: {str(e)}"
            else:
                return f"Error: {response.status_code} - {response.text}"


    async def upsert_character(self, character:CharacterInput) -> CharacterClientResponse:
        url = f"{self.__config.CRUD_SVC_HOST}/character"
        
        async with httpx.AsyncClient() as client:
            response = await client.post(url,json=jsonable_encoder(character))

            if response.status_code == 200:
                try:
                    return TypeAdapter(CharacterClientResponse).validate_python(response.json())
                except Exception as e:
                    return f"Error parsing response: {str(e)}"
            else:
                return f"Error: {response.status_code} - {response.text}"


            

    async def delete_character(self, id:int) -> dict[str,int]:

        url = f"{self.__config.CRUD_SVC_HOST}/character/{id}"
        
        async with httpx.AsyncClient() as client:
            response = await client.delete(url)

            if response.status_code == 200:
                try:
                    return TypeAdapter(dict[str,int]).validate_python(response.json())
                except Exception as e:
                    return f"Error parsing response: {str(e)}"
            else:
                return f"Error: {response.status_code} - {response.text}"


    async def get_total(self) -> int:
        url = f"{self.__config.CRUD_SVC_HOST}/character/total"
        
        async with httpx.AsyncClient() as client:
            response = await client.get(url)

            if response.status_code == 200:
                try:
                    return TypeAdapter(int).validate_python(response.json())
                except Exception as e:
                    return f"Error parsing response: {str(e)}"
            else:
                return f"Error: {response.status_code} - {response.text}"