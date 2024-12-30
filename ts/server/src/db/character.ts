import {
  Character,
  CharactersPageParams,
  CharacterByIDParams,
} from "../types/character.ts";
import { QuerySignature } from "../types/queryFunc.ts";
import {
  Pool,
  PoolClient,
  Transaction,
} from "https://deno.land/x/postgres@v0.17.0/mod.ts";

export async function getCharacters(
  queryFunc: QuerySignature,
  { limit, offset }: CharactersPageParams
): Promise<Character[]> {
  const query = "select * from character order by id asc limit $1  offset $2 ";
  return await queryFunc<Character>(query, [limit, offset]);
}

export async function getCharacterByID(
  queryFunc: QuerySignature,
  { id }: CharacterByIDParams
): Promise<Character[]> {
  const query = "select * from character where id = $1";
  return await queryFunc<Character>(query, [id]);
}


export async function deleteCharacter(
  pool: Pool,
  params: CharacterByIDParams
): Promise<void> {
  let client: PoolClient | undefined = undefined;
  let transaction: Transaction | undefined = undefined;
  try {
    client = await pool.connect();

    transaction = client.createTransaction("transaction");
    await transaction.begin();

    await transaction.queryObject(
      "DELETE FROM character_episode WHERE character_id = $1",
      [params.id]
    );

    await transaction.queryObject("DELETE FROM character WHERE id = $1", [
      params.id,
    ]);

    await transaction.commit();
    // https://deno-postgres.com/#/?id=transaction-errors
    // Transaction errors
    // When you are inside a Transaction block in PostgreSQL, reaching an error is terminal for the transaction. Executing the following in PostgreSQL will cause all changes to be undone and the transaction to become unusable until it has ended.
    //
    // "Therefore, we do not roll back here due to the fact that the transaction has already been rolled back."
    // } catch (error) {
    //   await transaction?.rollback({ chain: true });
    //   throw error;
  } finally {
    client?.release();
  }
}

export async function getTotal(queryFunc: QuerySignature): Promise<number> {
  const query = "select count(*) from character;";
  const rows = await queryFunc<{ count: number }>(query);
  if (rows.length > 0) {
    return rows[0].count;
  }
  throw Error("failed getting count");
}

export async function upsertCharacter(
  queryFunc: QuerySignature,
  character: Character
): Promise<Character[]> {
  const fields: (keyof Character)[] = [
    "id",
    "name",
    "status",
    "species",
    "type",
    "gender",
    "origin_id",
    "location_id",
    "image",
    "url",
    "created",
  ];

  const hasId = character.id !== null && character.id !== undefined;
  const filteredFields = hasId ? fields : fields.filter((field) => field !== "id");

  const columnNames = filteredFields.join(", ");
  const placeholders = filteredFields.map((_, i) => `$${i + 1}`).join(", ");

  const values = filteredFields.map((field) => character[field]);

  const updateColumns = fields
    .filter((field) => field !== "id") 
    .map((field) => `${field} = EXCLUDED.${field}`)
    .join(", ");

  const query = `
    INSERT INTO character (${columnNames})
    VALUES (${placeholders})
    ON CONFLICT (id)
    DO UPDATE SET ${updateColumns}
    RETURNING *;
  `;

  return await queryFunc<Character>(query, values);
}
