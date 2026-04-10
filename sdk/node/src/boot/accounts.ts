import { promises as fs } from "node:fs";
import { dirname } from "node:path";
import type { UserSvcToken } from "@1backend/client";
import { RegistrySvcApi, UserSvcApi } from "../client/api-client-factory";
import { randomUUID } from "node:crypto";
import { DEFAULT_APP_HOST } from "../constants";
import { CredentialStore, Credential } from "../types";

export async function registerServiceAccount(
  userService: UserSvcApi,
  serviceSlug: string,
  serviceName: string,
  store: CredentialStore,
): Promise<UserSvcToken> {
  let credential = await store.get();

  if (!credential) {
    credential = { slug: serviceSlug, password: randomUUID() };
    await store.set(credential);

    const registerRsp = await userService.register({
      appHost: DEFAULT_APP_HOST,
      slug: serviceSlug,
      name: serviceName,
      password: credential.password,
    });

    return registerRsp.body.token;
  }

  try {
    const loginRsp = await userService.login({
      appHost: DEFAULT_APP_HOST,
      slug: serviceSlug,
      password: credential.password,
    });

    return loginRsp.body.token;
  } catch (loginErr) {
    const registerRsp = await userService.register({
      appHost: DEFAULT_APP_HOST,
      slug: serviceSlug,
      name: serviceName,
      password: credential.password,
    });
    if (!registerRsp?.body?.token) {
      throw loginErr;
    }

    return registerRsp.body.token;
  }
}

export async function registerUserAccount(
  userService: UserSvcApi,
  appHost: string,
  slug: string,
  password: string,
  name: string,
): Promise<UserSvcToken> {
  await userService.register({ appHost, slug, password, name });
  const loginRsp = await userService.login({ appHost, slug, password });
  return loginRsp.body.token;
}

export async function registerServiceInstance(registryService: RegistrySvcApi, selfUrl: string): Promise<void> {
  await registryService.registerInstance({ url: selfUrl });
}

export function listenAddress(url: string): string {
  return url.replace("https://", "").replace("http://", "");
}

export function createCredentialFileStore(filePath: string): CredentialStore {
  return {
    async get(): Promise<Credential | null> {
      try {
        const raw = await fs.readFile(filePath, "utf8");
        return JSON.parse(raw) as Credential;
      } catch {
        return null;
      }
    },
    async set(credential: Credential): Promise<void> {
      await fs.mkdir(dirname(filePath), { recursive: true });
      await fs.writeFile(filePath, JSON.stringify(credential, null, 2), "utf8");
    },
  };
}
