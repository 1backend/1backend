import { DEFAULT_SERVER_URL } from "../constants";
import { OneBackendApiClientFactory, ApiClientFactory } from "../client/api-client-factory";
import { PermissionChecker, PermissionCheckerImpl } from "../endpoint/permission-checker";
import { TokenRefresher, TokenRefresherImpl } from "../endpoint/token-refresher";
import { defaultMiddlewares, ServiceMiddleware } from "../middlewares/default-middlewares";

export interface BootOptions {
  test?: boolean;
  serverUrl?: string;
  selfUrl?: string;
  tokenAutoRefreshOff?: boolean;
  clientFactory?: ApiClientFactory;
  tokenRefresher?: TokenRefresher;
  permissionChecker?: PermissionChecker;
  middlewares?: ServiceMiddleware[];
}

export interface LoadedBootOptions {
  test: boolean;
  serverUrl: string;
  selfUrl: string;
  tokenAutoRefreshOff: boolean;
  clientFactory: ApiClientFactory;
  tokenRefresher: TokenRefresher;
  permissionChecker: PermissionChecker;
  middlewares: ServiceMiddleware[];
  api: ReturnType<ApiClientFactory["create"]>;
}

export async function createDefaultOptions(options: BootOptions = {}): Promise<LoadedBootOptions> {
  const serverUrl = options.serverUrl || process.env.OB_SERVER_URL || DEFAULT_SERVER_URL;
  const selfUrl = options.selfUrl || process.env.OB_SELF_URL || "";
  const test = options.test ?? process.env.OB_TEST === "true";
  const tokenAutoRefreshOff = options.tokenAutoRefreshOff ?? process.env.OB_TOKEN_AUTO_REFRESH_OFF === "true";

  const clientFactory = options.clientFactory ?? new OneBackendApiClientFactory(serverUrl);
  const tokenRefresher = options.tokenRefresher ?? new TokenRefresherImpl(clientFactory);
  const permissionChecker = options.permissionChecker ?? new PermissionCheckerImpl(clientFactory);

  const middlewares = options.middlewares ?? defaultMiddlewares(tokenRefresher, tokenAutoRefreshOff);

  return {
    test,
    serverUrl,
    selfUrl,
    tokenAutoRefreshOff,
    clientFactory,
    tokenRefresher,
    permissionChecker,
    middlewares,
    api: clientFactory.create(),
  };
}
