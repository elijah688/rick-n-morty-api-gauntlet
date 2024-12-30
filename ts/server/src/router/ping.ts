import { Context, Router } from "https://deno.land/x/oak@v17.1.3/mod.ts";
import { Services } from "../services/services.ts";

export const router = new Router();

export class PingRouter {
  public router: Router;
  private svc: Services;

  constructor(svc: Services) {
    this.svc = svc;
    this.router = new Router();

    this.router.get("/ping", async (ctx: Context) => {
      ctx.response.body =  await this.svc.Ping() ;
    });
  }
}
