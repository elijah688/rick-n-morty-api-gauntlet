from sqlalchemy import func,select,delete
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker, aliased
from sqlalchemy.ext.asyncio import create_async_engine, AsyncSession
from config.config import Config
from model.character import Character
from typing import List, Dict, Tuple
from model.episode import Episode,CharacterEpisode, EpisodeResponse
from collections import defaultdict
from model.location import Location

Base = declarative_base()

class Database:
    def __init__(self, config: Config):
        self.config = config
        self.database_url = f"postgresql+asyncpg://{config.MAIN_DB_USER}:{config.MAIN_DB_PASS}@{config.MAIN_DB_HOST}:{config.MAIN_DB_PORT}/{config.APP_DB_NAME}"
        self.engine = None
        self.SessionLocal = None

    async def connect(self):
        self.engine = create_async_engine(self.database_url, echo=True, future=True)
        self.SessionLocal = sessionmaker(
            bind=self.engine,
            class_=AsyncSession,
            expire_on_commit=False
        )
        print("Connected to the database.")

    async def close(self):
        if self.engine:
            await self.engine.dispose()
            print("Closed the database connection.")

    def get_session(self) -> AsyncSession:
        return self.SessionLocal()

    async def ping(self) -> str:
        async with self.get_session() as session:
            result = await session.execute(select(func.version()))
            version = result.scalar_one()  
            return version
        
    async def get_characters(self, limit:int, offset:int) -> List[Character]:
        async with self.get_session() as session:
            result = await session.execute(
                select(Character) 
                .limit(limit)
                .offset(offset)
                .order_by(Character.id.asc()) 
            )
            return result.scalars().all() 
        
    async def get_character_by_id(self,id:int) -> Character:
         async with self.get_session() as session:
            result = await session.execute(
                select(Character) 
                .where(Character.id == id)
            )
            return result.scalars().one()
        
        


    async def upsert_character(self, character: Character):
        async with self.get_session() as session:
            c= await session.merge(character)
            await session.commit()  
            return c
    
    
    async def get_chars_debut(self, ids: List[int])->Tuple:
        async with self.get_session() as session:
            ranked_subquery = (
                select(
                    CharacterEpisode.character_id,
                    CharacterEpisode.episode_id.label('id'),
                    Episode.name,
                    Episode.air_date,
                    Episode.episode_code,
                    Episode.url.label('url'),
                    Episode.created.label('created'),
                    func.row_number().over(
                        partition_by=CharacterEpisode.character_id,
                        order_by=CharacterEpisode.episode_id.asc()
                    ).label('rn')
                )
                .join(Episode, Episode.id == CharacterEpisode.episode_id)
                .where(CharacterEpisode.character_id.in_(ids))
                .alias('ranked_episodes')
            )

            main_query = (
                select(
                    ranked_subquery.c.character_id,
                    ranked_subquery.c.id,
                    ranked_subquery.c.name,
                    ranked_subquery.c.air_date,
                    ranked_subquery.c.episode_code,
                    ranked_subquery.c.url,
                    ranked_subquery.c.created
                )
                .where(ranked_subquery.c.rn == 1) 
                .order_by(ranked_subquery.c.character_id.asc())
            )

            result = await session.execute(main_query)
            
            return result.fetchall()

    async def get_episodes_by_ids(self, ids: List[int]) -> Dict[int,List[int]]:
        async with self.get_session() as session:
            result = await session.execute(
             select(
                 CharacterEpisode,
             ).
             where(CharacterEpisode.character_id.in_(ids))
            )

        res = defaultdict(list)
        for x in  result.scalars().all():
            res[x.character_id].append(x.episode_id)
        
        return res
        

    async def get_locations_by_ids(self, ids: List[int]) -> Dict[int,Location]:
         async with self.get_session() as session:
            result = await session.execute(
             select(
                Location
             ).
             where(Location.id.in_(ids))
            )
            

         return { location.id: location for location in result.scalars().all()}
        
   

    async def delete_character(self, id: int):
        async with self.get_session() as session:
            async with session.begin():  
                await session.execute(
                    delete(CharacterEpisode).where(CharacterEpisode.character_id == id)
                )

                await session.execute(
                    delete(Character).where(Character.id == id)
                )

            await session.commit()
    
    
    async def get_total(self) -> int:
        async with self.get_session() as session:
            result = await session.execute(
                select(func.count(Character.id))
            )
        return  result.scalar_one() 
        