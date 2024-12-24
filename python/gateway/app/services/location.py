
from config.config import Config
from model.location import Location
from typing import List
from pydantic import TypeAdapter
import httpx
class LocationService:
    def __init__(self, config: Config):
        self.__config = config

    async def get_locations_by_ids(self, ids:List[int]) -> dict[int, Location]:
        url = f"{self.__config.CRUD_SVC_HOST}/location/list"
        print(ids)
        async with httpx.AsyncClient() as client:
            response = await client.post(url, json={"ids":[x for x in ids if x is not None]})

            if response.status_code == 200:
                try:
                    return TypeAdapter(dict[int,Location]).validate_python(response.json())
                except Exception as e:
                    return f"Error parsing response: {str(e)}"
            else:
                return f"Error: {response.status_code} - {response.text}"