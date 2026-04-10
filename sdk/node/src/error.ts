export function oneBackendApiError(err: unknown): Error {
  if (!err) return new Error("unknown error");
  if (err instanceof Error) return err;

  if (typeof err === "object" && err !== null) {
    const maybe = err as { body?: unknown };
    if (maybe.body && typeof maybe.body === "object" && "error" in (maybe.body as Record<string, unknown>)) {
      return new Error(String((maybe.body as Record<string, unknown>).error));
    }
  }

  return new Error(String(err));
}
