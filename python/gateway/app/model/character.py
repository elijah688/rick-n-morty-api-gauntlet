from pydantic import BaseModel,HttpUrl
from typing import Optional, List
from datetime import datetime
from model.location import Location
from model.episode import Episode

class CharacterClientResponse(BaseModel):
    id: int
    name: str
    status: Optional[str] = None
    species: Optional[str] = None
    type: Optional[str] = None
    gender: Optional[str] = None
    origin_id: Optional[int] = None
    location_id: Optional[int] = None
    image: Optional[str] = None
    url: Optional[str] = None
    created: Optional[datetime] = None
    
    class Config:
        from_attributes = True
     
class CharacterResponse(BaseModel):
    id: int
    name: str
    status: Optional[str] = None
    species: Optional[str] = None
    type: Optional[str] = None
    gender: Optional[str] = None
    origin: Optional[Location] = None
    location: Optional[Location] = None
    image: Optional[str] = None
    url: Optional[str] = None
    created: Optional[datetime] = None
    episodes: Optional[List[int]] = None
    debut: Optional[Episode] = None

    class Config:
        from_attributes = True
   


class CharacterLocationInput(BaseModel):
    url: Optional[HttpUrl]

class CharacterInput(BaseModel):
    id: Optional[int]
    name: str
    status: Optional[str]
    species: Optional[str]
    type: Optional[str]
    gender: Optional[str]
    origin: CharacterLocationInput
    location: CharacterLocationInput
    image: Optional[HttpUrl]
    episodes: Optional[List[int]]
    url: Optional[HttpUrl]
    created: Optional[str]


class CharactersResponse(BaseModel):
    characters: List[CharacterResponse]
    total: int
