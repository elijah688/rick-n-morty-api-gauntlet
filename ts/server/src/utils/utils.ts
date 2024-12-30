/**
 * Ensures the given error is an instance of Error.
 * If it's not, converts it into an Error with a stringified message.
 *
 * @param error - The error to normalize
 * @returns An instance of Error
 */
export function normalizeError(error: unknown): Error {
  return error instanceof Error ? error : new Error(String(error));
}
