import {
  Context,
  Router,
  RouterContext,
} from "https://deno.land/x/oak@v17.1.3/mod.ts";

import { Services } from "../services/services.ts";
import { Character, } from "../types/character.ts";
import {GetLocationsByIDs} from "../types/location.ts"
import {GetCharactersEpisodesByIDs} from "../types/episodes.ts"
export class CharacterRouter {
  private svc: Services;
  public router: Router;

  constructor(svc: Services) {
    this.svc = svc;
    this.router = new Router();

    this.router.get("/total", async (ctx: Context) => {
      const big = await this.svc.GetTotal();
      ctx.response.body = { total: Number(big) };
    });

    this.router.get("/character", async (ctx: Context) => {
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

      ctx.response.body = await this.svc.GetCharacters({
        limit: parsedLimit,
        offset: parsedOffset,
      });
    });

    this.router.get(
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
        const character = await this.svc.GetCharacterByID({ id: parsedId });
        ctx.response.body = character;
      }
    );

    this.router.post("/character", async (ctx: Context) => {
      const char: Character = await ctx.request.body.json();
      const character = await this.svc.UpsertCharacter(char);

      ctx.response.body = character;
    });

    this.router.delete(
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
        await this.svc.DeleteCharacterByID({ id: parsedId });

        ctx.response.body = { id: ctx.params.id };
      }
    );

    this.router.post("/location/list", async (ctx: Context) => {
      const lids: GetLocationsByIDs = await ctx.request.body.json();
      const locs = await this.svc.GetLocationsByIDs(lids);

      ctx.response.body = locs;
    });

    this.router.post("/character/list/episodes", async (ctx: Context) => {
      const cids: GetCharactersEpisodesByIDs = await ctx.request.body.json();
      const episodess = await this.svc.GetEpisodesByCharIDs(cids);

      ctx.response.body = episodess;
    });

    this.router.post("/character/list/debut", async (ctx: Context) => {
      const cids: GetCharactersEpisodesByIDs = await ctx.request.body.json();
      const episodess = await this.svc.GetDebuts(cids);

      ctx.response.body = episodess;
    });


  }
}

