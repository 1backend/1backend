export type HeadersLike = Record<string, string | string[] | undefined>;

export interface RequestLike {
  headers: HeadersLike;
}

export interface Credential {
  slug: string;
  password: string;
}

export interface CredentialStore {
  get(): Promise<Credential | null>;
  set(credential: Credential): Promise<void>;
}
