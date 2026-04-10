import jwt from "jsonwebtoken";

export interface ParsedClaims {
  exp?: number;
  [key: string]: unknown;
}

export function parseJwt(userSvcPublicKey: string, token: string): ParsedClaims {
  return jwt.verify(token, userSvcPublicKey, { algorithms: ["RS256"] }) as ParsedClaims;
}
