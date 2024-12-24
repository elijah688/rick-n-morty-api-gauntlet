from pydantic import BaseModel
from datetime import datetime
from typing import Optional

class Episode(BaseModel):
    id: int
    name: str
    air_date: Optional[str] = None
    episode_code: Optional[str] = None
    url: Optional[str] = None
    created: Optional[datetime] = None
    
    class Config:
        from_attributes = True