from fastapi import APIRouter, Query,Body
from services.services import Services
from model.location import LocationRequest

class LocationRouter:
    def __init__(self, services: Services):
        self.__services = services
        self.router = APIRouter()  
        

        @self.router.post("/list")
        async def get_locations_by_ids(req: LocationRequest):
            return await self.__services.get_locations_by_ids(req.ids)
