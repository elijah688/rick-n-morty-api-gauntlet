from pydantic import BaseModel
from typing import Optional,List
from datetime import datetime
from sqlalchemy import Column, Integer, String, DateTime
from sqlalchemy.ext.declarative import declarative_base
from datetime import datetime,timezone

class CharacterResponse(BaseModel):
    id: Optional[int] = None
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

    def to_sql_model(self) -> "Character":
        c = Character(
            id=self.id,
            name=self.name,
            status=self.status,
            species=self.species,
            type=self.type,
            gender=self.gender,
            origin_id=self.origin_id,
            location_id=self.location_id,
            image=self.image,
            url=self.url,
            created=datetime.now(timezone.utc)
        )

        if self.created is not None:
            c.created = self.created.replace(tzinfo=None)

        return c

    class Config:
        from_attributes = True

Base = declarative_base()

class Character(Base):
    __tablename__ = 'character'

    id = Column(Integer, primary_key=True)
    name = Column(String, nullable=False)
    status = Column(String)
    species = Column(String)
    type = Column(String)
    gender = Column(String)
    origin_id = Column(Integer)
    location_id = Column(Integer)
    image = Column(String)
    url = Column(String)
    created = Column(DateTime)

class CharactersEpisodesRequest(BaseModel):
    ids:  List[int]

class CharactersDebuteRequest(BaseModel):
    ids:  List[int]

