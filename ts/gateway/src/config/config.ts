import { config } from "https://deno.land/x/dotenv@v3.2.2/mod.ts";
import { ENV_VARS } from "../constants/constants.ts"

export interface Config {
  crudServicsHost: string;
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
      crudServicsHost: env[ENV_VARS.CRUD_SVC_HOST],
      port:  parseInt(env[ENV_VARS.GATEWAY_SERVER_PORT], 10),
    };
  }

  getConfig(): Config {
    return this.config;
  }
}

export default AppConfig;
