from fastapi import FastAPI
from router.character_router import CharacterRouter
from router.location_router import LocationRouter
from services.services import Services
from db.db import Database
from config.config import Config
import uvicorn

config = Config()
db = Database(config)
services = Services(db)

async def lifespan(app: FastAPI):
    await db.connect()
    yield
    await db.close()

app = FastAPI(
    lifespan=lifespan
)

app.include_router(CharacterRouter(services).router, prefix="/character")
app.include_router(LocationRouter(services).router, prefix="/location")


if __name__ == "__main__":
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=int(config.PORT),
        reload=True
    )
