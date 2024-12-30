import { Services } from "../services/services.ts";
import {
  Context,
  Router,
  RouterContext,
} from "https://deno.land/x/oak@v17.1.3/mod.ts";
import { Character } from "../types/characters.ts";

export class CharacterRouter {
  private _router: Router;
  private svc: Services;

  constructor(s: Services) {
    this.svc = s;

    this._router = new Router();
    this._router.get(
      "/character/:id",
      async (ctx: RouterContext<"/character/:id", { id: string }>) => {
        const parsedId = parseInt(ctx.params.id, 10);
        if (isNaN(parsedId)) {
          ctx.response.status = 429;
          ctx.response.body = {
            error: "Invalid 'id' parameter. Must be a valid integer.",
          };
          return;
        }
        const character = await this.svc.getCharacterByID(parsedId);
        ctx.response.body = character;
      }
    );

    this._router.get("/character", async (ctx: Context) => {
      const limit = ctx.request.url.searchParams.get("limit");
      const offset = ctx.request.url.searchParams.get("offset");
      const parsedLimit = limit ? parseInt(limit, 10) : 10;
      const parsedOffset = offset ? parseInt(offset, 10) : 0;

      if (isNaN(parsedLimit) || isNaN(parsedOffset)) {
        ctx.response.status = 400;
        ctx.response.body = {
          error: "'limit' and 'offset' must be valid numbers",
        };
        return;
      }

      const [characters, total] = await Promise.all([
        this.svc.getCharacters({
          limit: parsedLimit,
          offset: parsedOffset,
        }),
        this.svc.total(),
      ]);

      ctx.response.body = { characters, total };
    });

    this._router.delete(
      "/character/:id",
      async (ctx: RouterContext<"/character/:id", { id: string }>) => {
        const parsedId = parseInt(ctx.params.id, 10);
        if (isNaN(parsedId)) {
          ctx.response.status = 429;
          ctx.response.body = {
            error: "Invalid 'id' parameter. Must be a valid integer.",
          };
          return;
        }
        const character = await this.svc.deleteCharacter(parsedId);
        ctx.response.body = character;
      }
    );

    this._router.post("/character", async (ctx: Context) => {
      const char: Character = await ctx.request.body.json();
      ctx.response.body =  await this.svc.upsertCharacter(char);
    });
  }

  get router() {
    return this._router;
  }
}
