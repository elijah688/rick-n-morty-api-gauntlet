import { QueryArguments } from "https://deno.land/x/postgres@v0.17.0/query/query.ts";
export type QuerySignature = <T>(queryString: string, queryArgs?:QueryArguments) => Promise<T[]>;
