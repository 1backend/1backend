import { customAlphabet } from "nanoid";
import { randomUUID } from "node:crypto";
import { DEFAULT_APP_HOST, DEFAULT_TEST_APP_HOST } from "./constants";

const idSeparator = "-";
const nanoid = customAlphabet("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 10);

export function id(prefix = ""): string {
  const suffix = nanoid();
  if (!prefix) return suffix;
  return `${prefix.replace(/-+$/, "")}${idSeparator}${suffix}`;
}

export function opaqueId(prefix: string): string {
  return `${prefix.replace(/-+$/, "")}${idSeparator}${randomUUID()}`;
}

export function deterministicId(prefix: string, sourceId: string): string {
  const cleanPrefix = prefix.replace(/-+$/, "");
  const cleanId = sourceId.replaceAll("_", idSeparator);
  return `${cleanPrefix}${idSeparator}${cleanId}`;
}

export function marshal(value: unknown): string {
  return JSON.stringify(value);
}

export function internalId(appId: string, valueId: string): string {
  if (
    !appId.startsWith("app_") &&
    !appId.startsWith("app-") &&
    appId !== "*" &&
    appId !== DEFAULT_APP_HOST &&
    appId !== DEFAULT_TEST_APP_HOST
  ) {
    throw new Error(
      `appId must start with 'app_' or be '*', '${DEFAULT_APP_HOST}', '${DEFAULT_TEST_APP_HOST}', got: '${appId}'`,
    );
  }

  return `${appId}:${valueId}`;
}
