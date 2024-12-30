import {
  Application,
} from "https://deno.land/x/oak@v17.1.3/mod.ts";

import { AppConfig } from "./src/config/config.ts";
import { DB } from "./src/db/db.ts";
import { Services } from "./src/services/services.ts";
import { CharacterRouter } from "./src/router/character.ts";
import { PingRouter } from "./src/router/ping.ts";
const app = new Application();

const config = new AppConfig().getConfig();
const db = new DB(config);
const services = new Services(db);

const pingRouter = new PingRouter(services)
app.use(pingRouter.router.routes());
app.use(pingRouter.router.allowedMethods());

const charRouter = new CharacterRouter(services)
app.use(charRouter.router.routes());
app.use(charRouter.router.allowedMethods());

console.log(`Server is running on http://localhost:${config.port}`);
await app.listen({ port: config.port });
