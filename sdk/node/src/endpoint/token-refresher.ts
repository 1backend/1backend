import { LRUCache } from "lru-cache";
import { parseJwt, ParsedClaims } from "../auth/jwt";
import { ApiClientFactory } from "../client/api-client-factory";
import { RequestLike } from "../types";

function getAuthorization(headers: RequestLike["headers"]): string {
  const raw = headers.authorization ?? headers.Authorization;
  return Array.isArray(raw) ? raw[0] ?? "" : raw ?? "";
}

export interface TokenRefresher {
  ensureValidToken(request: RequestLike): Promise<{ token: string; claims: ParsedClaims | null }>;
}

export class TokenRefresherImpl implements TokenRefresher {
  private readonly tokenReplacementCache = new LRUCache<string, string>({ max: 1000, ttl: 60_000 });
  private publicKey?: string;

  constructor(private readonly clientFactory: ApiClientFactory) {}

  async ensureValidToken(request: RequestLike): Promise<{ token: string; claims: ParsedClaims | null }> {
    const authHeader = getAuthorization(request.headers);
    if (!authHeader.startsWith("Bearer ")) {
      return { token: "", claims: null };
    }

    let token = authHeader.slice(7).trim();
    const publicKey = await this.getPublicKey();

    const cached = this.tokenReplacementCache.get(token);
    if (cached) {
      token = cached;
      request.headers.authorization = `Bearer ${token}`;
    }

    const claims = this.tryParseClaims(publicKey, token);
    if (claims && claims.exp && claims.exp * 1000 > Date.now()) {
      return { token, claims };
    }

    const refreshClient = this.clientFactory.create(token);
    const refresh = await refreshClient.userSvc.refreshToken();
    const newToken = refresh.body.token?.token;
    if (!newToken) {
      throw new Error("token refresh response did not contain a token");
    }

    request.headers.authorization = `Bearer ${newToken}`;
    this.tokenReplacementCache.set(token, newToken);

    return { token: newToken, claims: this.tryParseClaims(publicKey, newToken) };
  }

  private tryParseClaims(publicKey: string, token: string): ParsedClaims | null {
    try {
      return parseJwt(publicKey, token);
    } catch {
      return null;
    }
  }

  private async getPublicKey(): Promise<string> {
    if (this.publicKey) return this.publicKey;
    const api = this.clientFactory.create();
    const response = await api.userSvc.getPublicKey();
    this.publicKey = response.body.publicKey;
    return this.publicKey;
  }
}
