from pydantic import BaseModel
from datetime import datetime
from typing import Optional

class Location(BaseModel):
    id: int
    name: str
    type: Optional[str] = None
    dimension: Optional[str] = None
    url: Optional[str] = None
    created: Optional[datetime] = None
  
    class Config:
        from_attributes = True