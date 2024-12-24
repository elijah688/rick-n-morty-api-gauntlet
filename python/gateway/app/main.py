from fastapi import FastAPI, Depends
from config.config import Config
from services.services import Services
from router.character_router import CharacterRouter
import uvicorn
from fastapi.middleware.cors import CORSMiddleware


app = FastAPI()

config = Config()
services = Services(config)

async def lifespan(app: FastAPI):
    yield  

app = FastAPI(lifespan=lifespan)

app.include_router(
    router=CharacterRouter(services).get_router(),
    prefix="/character",
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  
    allow_credentials=True,
    allow_methods=["*"],  
    allow_headers=["*"], 
)

if __name__ == "__main__":
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=int(config.GATEWAY_SERVER_PORT),
        reload=True
    )
