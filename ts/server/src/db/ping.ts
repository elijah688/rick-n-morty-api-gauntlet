import { PingResult } from "../types/ping.ts";
import { QuerySignature } from "../types/queryFunc.ts";

export async function ping(queryFunc: QuerySignature): Promise<PingResult> {
  const rows = await queryFunc<PingResult>("select version();");
  if (rows.length > 0) {
    return { version: rows[0].version };
  }

  return { version: null };
}

