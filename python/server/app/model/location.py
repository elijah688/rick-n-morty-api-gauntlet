from datetime import datetime
from sqlalchemy import Column, Integer, Text, DateTime
from sqlalchemy.ext.declarative import declarative_base
from pydantic import BaseModel, Field
from typing import Optional,Dict, List

Base = declarative_base()

class Location(Base):
    __tablename__ = "location"

    id = Column(Integer, primary_key=True)
    name = Column(Text, nullable=False)
    type = Column(Text, nullable=True)
    dimension = Column(Text, nullable=True)
    url = Column(Text, nullable=True)
    created = Column(DateTime, nullable=True)

class LocationResponse(BaseModel):
    id: int
    name: str
    type: Optional[str] = Field(default=None)
    dimension: Optional[str] = Field(default=None)
    url: Optional[str] = Field(default=None)
    created: Optional[datetime] = Field(default=None)

    class Config:
        from_attributes = True



class LocationRequest(BaseModel):
    ids:  List[int]

   