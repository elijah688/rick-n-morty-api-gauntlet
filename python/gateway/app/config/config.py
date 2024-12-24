from pydantic import  Field
from pydantic_settings import  BaseSettings

class Config(BaseSettings):
    CRUD_SVC_HOST: str = Field(..., env="CRUD_SVC_HOST")
    GATEWAY_SERVER_PORT: str = Field(..., env="GATEWAY_SERVER_PORT")

    class Config:
        pass
     
config = Config()
