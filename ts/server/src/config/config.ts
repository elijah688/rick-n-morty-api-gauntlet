import { config } from "https://deno.land/x/dotenv@v3.2.2/mod.ts";
import { ENV_VARS } from "../constants/constants.ts"; 

export interface Config {
  rikiApiBaseUrl: string;
  mainDbHost: string;
  mainDbPort: number;
  mainDbUser: string;
  mainDbPass: string;
  mainDbName: string;
  appDbName: string;
  port: number
}

export class AppConfig {
  private config: Config;

  constructor() {
    const env = config();

    const requiredConfigKeys = Object.values(ENV_VARS);

    for (const key of requiredConfigKeys) {
      if (!env[key]) {
        throw new Error(`Missing required environment variable: ${key}`);
      }
    }

    this.config = {
      rikiApiBaseUrl: env[ENV_VARS.RIKI_API_BASE_URL],
      mainDbHost: env[ENV_VARS.MAIN_DB_HOST],
      mainDbPort: parseInt(env[ENV_VARS.MAIN_DB_PORT], 10),
      mainDbUser: env[ENV_VARS.MAIN_DB_USER],
      mainDbPass: env[ENV_VARS.MAIN_DB_PASS],
      mainDbName: env[ENV_VARS.MAIN_DB_NAME],
      appDbName: env[ENV_VARS.APP_DB_NAME]!,
      port:  parseInt(env[ENV_VARS.PORT], 10),
    };
  }

  getConfig(): Config {
    return this.config;
  }
}

export default AppConfig;
