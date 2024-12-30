import { AppConfig } from "./src/config/config.ts";
import { Services } from "./src/services/services.ts";
import { Application } from "https://deno.land/x/oak/mod.ts";
import { CharacterRouter } from "./src/router/router.ts";
import { oakCors } from "https://deno.land/x/cors/mod.ts";

const cfg = new AppConfig().getConfig();

const services = new Services(cfg);
const cr = new CharacterRouter(services);
const app = new Application();

app.use(oakCors());
app.use(cr.router.allowedMethods());
app.use(cr.router.routes());

console.log(`TS Gateway Listening on ${cfg.port}...`)
await app.listen({ port: cfg.port });