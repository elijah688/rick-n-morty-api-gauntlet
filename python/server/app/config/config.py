from pydantic import  Field
from pydantic_settings import  BaseSettings
BaseSettings
class Config(BaseSettings):
    RIKI_API_BASE_URL: str = Field(..., env="RIKI_API_BASE_URL")
    MAIN_DB_HOST: str = Field(..., env="MAIN_DB_HOST")
    MAIN_DB_PORT: int = Field(..., env="MAIN_DB_PORT")
    MAIN_DB_USER: str = Field(..., env="MAIN_DB_USER")
    MAIN_DB_PASS: str = Field(..., env="MAIN_DB_PASS")
    MAIN_DB_NAME: str = Field(..., env="MAIN_DB_NAME")
    APP_DB_NAME: str = Field(..., env="APP_DB_NAME")
    PORT: str = Field(..., env="PORT")
    

    class Config:
        pass
        # env_file = ".env"
        # env_file_encoding = 'utf-8'

config = Config()
