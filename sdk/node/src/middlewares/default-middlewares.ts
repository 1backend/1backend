import { IncomingMessage, ServerResponse } from "node:http";
import { gzipSync } from "node:zlib";
import { TokenRefresher } from "../endpoint/token-refresher";

export type Next = () => Promise<void>;
export type ServiceMiddleware = (
  req: IncomingMessage & { headers: Record<string, string | string[] | undefined> },
  res: ServerResponse,
  next: Next,
) => Promise<void> | void;

export async function applyMiddlewares(
  middlewares: ServiceMiddleware[],
  req: IncomingMessage & { headers: Record<string, string | string[] | undefined> },
  res: ServerResponse,
  handler: () => Promise<void> | void,
): Promise<void> {
  let index = -1;

  const run = async (i: number): Promise<void> => {
    if (i <= index) throw new Error("next() called multiple times");
    index = i;

    const middleware = middlewares[i];
    if (!middleware) {
      await handler();
      return;
    }

    await middleware(req, res, async () => run(i + 1));
  };

  await run(0);
}

export function defaultMiddlewares(tokenRefresher: TokenRefresher, tokenAutoRefreshOff: boolean): ServiceMiddleware[] {
  return [
    tokenRefreshMiddleware(tokenRefresher, tokenAutoRefreshOff),
    recoverMiddleware,
    corsMiddleware,
    gzipEncodeMiddleware,
  ];
}

export function tokenRefreshMiddleware(tokenRefresher: TokenRefresher, disabled: boolean): ServiceMiddleware {
  return async (req, _res, next) => {
    if (!disabled) {
      await tokenRefresher.ensureValidToken(req);
    }
    await next();
  };
}

export const recoverMiddleware: ServiceMiddleware = async (_req, res, next) => {
  try {
    await next();
  } catch (err) {
    if (!res.headersSent) {
      res.statusCode = 500;
      res.setHeader("content-type", "application/json");
    }
    res.end(JSON.stringify({ error: err instanceof Error ? err.message : "unknown error" }));
  }
};

export const corsMiddleware: ServiceMiddleware = async (req, res, next) => {
  res.setHeader("Access-Control-Allow-Origin", "*");
  res.setHeader("Access-Control-Allow-Headers", "Content-Type, Authorization");
  res.setHeader("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS");

  if (req.method === "OPTIONS") {
    res.statusCode = 204;
    res.end();
    return;
  }

  await next();
};

export const gzipEncodeMiddleware: ServiceMiddleware = async (req, res, next) => {
  const acceptEncoding = String(req.headers["accept-encoding"] ?? "");
  if (!acceptEncoding.includes("gzip")) {
    await next();
    return;
  }

  const originalEnd = res.end.bind(res);
  res.end = ((chunk?: any, ...args: any[]) => {
    if (chunk && typeof chunk === "string") {
      res.setHeader("content-encoding", "gzip");
      return originalEnd(gzipSync(Buffer.from(chunk)), ...args);
    }
    return originalEnd(chunk, ...args);
  }) as ServerResponse["end"];

  await next();
};
