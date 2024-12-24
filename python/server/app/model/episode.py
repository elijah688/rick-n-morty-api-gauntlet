
from sqlalchemy import Column, Integer, String, DateTime
from sqlalchemy.ext.declarative import declarative_base
from pydantic import BaseModel,field_validator
from typing import Optional
from datetime import datetime

Base = declarative_base()

class CharacterEpisode(Base):
    __tablename__ = 'character_episode'

    character_id = Column(Integer, primary_key=True)
    episode_id = Column(Integer, primary_key=True)

class Episode(Base):
    __tablename__ = 'episode'

    id = Column(Integer, primary_key=True)
    name = Column(String)
    air_date = Column(DateTime)
    episode_code = Column(String)
    url = Column(String)
    created = Column(DateTime)
  

        
class CharacterEpisodeResponse(BaseModel):
    character_id: int
    episode_id: int

    class Config:
        from_attributes = True


class EpisodeResponse(BaseModel):
    id: int
    name: str
    air_date: Optional[str] = None
    episode_code: Optional[str] = None
    url: Optional[str] = None
    created: Optional[datetime] = None
    
    class Config:
        from_attributes = True