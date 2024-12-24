
from config.config import Config
from typing import List
from model.episode import Episode
import httpx
from pydantic import TypeAdapter

class EpisodeService:
    def __init__(self, config: Config):
        self.__config = config

    async def get_episodes_by_ids(self, ids: List[int]) -> dict[int,List[int]]:
        url = f"{self.__config.CRUD_SVC_HOST}/character/list/episodes"
        payload = {"ids": ids}

        async with httpx.AsyncClient() as client:
            response = await client.post(url, json=payload)

            if response.status_code == 200:
                try:
                    res = response.json()
                    return TypeAdapter(dict[int,List[int]]).validate_python(response.json())
                except Exception as e:
                    return f"Error parsing response: {str(e)}"
            else:
                return f"Error: {response.status_code} - {response.text}"

    async def get_characters_debut_by_ids(self, ids: List[int]) -> dict[int,Episode]:
        url = f"{self.__config.CRUD_SVC_HOST}/character/list/debut"
        payload = {"ids": ids}

        async with httpx.AsyncClient() as client:
            response = await client.post(url, json=payload)

            if response.status_code == 200:
                try:
                    return TypeAdapter(dict[int,Episode]).validate_python(response.json())
                except Exception as e:
                    return f"Error parsing response: {str(e)}"
            else:
                return f"Error: {response.status_code} - {response.text}"





