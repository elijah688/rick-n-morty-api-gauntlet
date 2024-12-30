import { GetLocationsByIDs, Location } from "../types/location.ts";
import { QuerySignature } from "../types/queryFunc.ts";

export async function getLocationsByIDs(
  queryFunc: QuerySignature,
  {ids}: GetLocationsByIDs
): Promise<Location[]> {
  const paramSymbols = ids.map((_, i) => `$${i+1}`).join(", ");
  const query = `
    SELECT * FROM location
    WHERE id in (${paramSymbols})
  `;


  return await queryFunc<Location>(query, ids);
}


