import type {
  UserSvcGetPublicKeyResponse as UserSvcPublicKeyResponse,
  UserSvcHasPermissionResponse,
  UserSvcToken,
} from "@1backend/client";

export interface ApiResponse<T> {
  response: { statusCode: number };
  body: T;
}

export interface UserSvcApi {
  getPublicKey(): Promise<ApiResponse<UserSvcPublicKeyResponse>>;
  register(body: { appHost: string; slug: string; password: string; name?: string }): Promise<ApiResponse<{ token: UserSvcToken }>>;
  login(body: { appHost: string; slug: string; password: string }): Promise<ApiResponse<{ token: UserSvcToken }>>;
  refreshToken(): Promise<ApiResponse<{ token: UserSvcToken }>>;
  hasPermission(permission: string): Promise<ApiResponse<UserSvcHasPermissionResponse>>;
}

export interface RegistrySvcApi {
  registerInstance(body: { url: string }): Promise<ApiResponse<object>>;
}

export interface ApiClient {
  userSvc: UserSvcApi;
  registrySvc: RegistrySvcApi;
}

export interface ApiClientFactory {
  create(token?: string): ApiClient;
}

async function request<T>(baseUrl: string, path: string, method: string, token?: string, body?: unknown): Promise<ApiResponse<T>> {
  const response = await fetch(new URL(path, baseUrl), {
    method,
    headers: {
      "content-type": "application/json",
      ...(token ? { authorization: `Bearer ${token}` } : {}),
    },
    body: body === undefined ? undefined : JSON.stringify(body),
  });

  const payload = (await response.json().catch(() => ({}))) as T;

  if (!response.ok) {
    const error = new Error(`1Backend API call failed: ${response.status}`) as Error & { body?: unknown; statusCode?: number };
    error.body = payload;
    error.statusCode = response.status;
    throw error;
  }

  return {
    response: { statusCode: response.status },
    body: payload,
  };
}

class UserSvcApiImpl implements UserSvcApi {
  constructor(private readonly baseUrl: string, private readonly token?: string) {}

  getPublicKey() {
    return request<UserSvcPublicKeyResponse>(this.baseUrl, "/user-svc/public-key", "GET", this.token);
  }

  register(body: { appHost: string; slug: string; password: string; name?: string }) {
    return request<{ token: UserSvcToken }>(this.baseUrl, "/user-svc/register", "POST", this.token, body);
  }

  login(body: { appHost: string; slug: string; password: string }) {
    return request<{ token: UserSvcToken }>(this.baseUrl, "/user-svc/login", "POST", this.token, body);
  }

  refreshToken() {
    return request<{ token: UserSvcToken }>(this.baseUrl, "/user-svc/refresh-token", "POST", this.token, {});
  }

  hasPermission(permission: string) {
    return request<UserSvcHasPermissionResponse>(
      this.baseUrl,
      `/user-svc/permission/${encodeURIComponent(permission)}`,
      "GET",
      this.token,
    );
  }
}

class RegistrySvcApiImpl implements RegistrySvcApi {
  constructor(private readonly baseUrl: string, private readonly token?: string) {}

  registerInstance(body: { url: string }) {
    return request<object>(this.baseUrl, "/registry-svc/instance", "POST", this.token, body);
  }
}

export class OneBackendApiClientFactory implements ApiClientFactory {
  constructor(private readonly serverUrl: string) {}

  create(token?: string): ApiClient {
    return {
      userSvc: new UserSvcApiImpl(this.serverUrl, token),
      registrySvc: new RegistrySvcApiImpl(this.serverUrl, token),
    };
  }
}
