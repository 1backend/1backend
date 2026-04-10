import { createHash } from "node:crypto";
import { LRUCache } from "lru-cache";
import type { UserSvcHasPermissionResponse } from "@1backend/client";
import { ApiClientFactory } from "../client/api-client-factory";
import { RequestLike } from "../types";

function extractJwt(request: RequestLike): string {
  const header = request.headers.authorization ?? request.headers.Authorization;
  const value = Array.isArray(header) ? header[0] : header;
  if (!value) return "";
  return value.replace(/^Bearer\s+/i, "").trim();
}

function cacheKey(token: string, permission: string): string {
  const hash = createHash("sha256").update(token).digest("hex");
  return `${hash}:${permission}`;
}

export interface PermissionChecker {
  hasPermission(request: RequestLike, permission: string): Promise<{ response: UserSvcHasPermissionResponse; statusCode: number }>;
}

export class PermissionCheckerImpl implements PermissionChecker {
  private readonly cache = new LRUCache<string, { response: UserSvcHasPermissionResponse; statusCode: number }>({
    max: 2000,
    ttl: 5 * 60 * 1000,
  });

  constructor(private readonly clientFactory: ApiClientFactory) {}

  async hasPermission(request: RequestLike, permission: string): Promise<{ response: UserSvcHasPermissionResponse; statusCode: number }> {
    const token = extractJwt(request);
    const key = cacheKey(token, permission);

    if (token) {
      const hit = this.cache.get(key);
      if (hit) return hit;
    }

    const client = this.clientFactory.create(token);
    const res = await client.userSvc.hasPermission(permission);
    const payload = { response: res.body, statusCode: res.response.statusCode ?? 0 };

    if (token && res.body.until) {
      const expiresAt = Date.parse(res.body.until);
      if (!Number.isNaN(expiresAt)) {
        this.cache.set(key, payload, { ttl: Math.max(1000, expiresAt - Date.now()) });
      }
    }

    return payload;
  }
}
