/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import http from 'http';
import { SecretSvcDecryptValueRequest } from '../model/secretSvcDecryptValueRequest';
import { SecretSvcDecryptValueResponse } from '../model/secretSvcDecryptValueResponse';
import { SecretSvcEncryptValueRequest } from '../model/secretSvcEncryptValueRequest';
import { SecretSvcEncryptValueResponse } from '../model/secretSvcEncryptValueResponse';
import { SecretSvcIsSecureResponse } from '../model/secretSvcIsSecureResponse';
import { SecretSvcListSecretsRequest } from '../model/secretSvcListSecretsRequest';
import { SecretSvcListSecretsResponse } from '../model/secretSvcListSecretsResponse';
import { SecretSvcRemoveSecretsRequest } from '../model/secretSvcRemoveSecretsRequest';
import { SecretSvcSaveSecretsRequest } from '../model/secretSvcSaveSecretsRequest';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum SecretSvcApiApiKeys {
    BearerAuth = 0
}
export declare class SecretSvcApi {
    protected _basePath: string;
    protected _defaultHeaders: any;
    protected _useQuerystring: boolean;
    protected authentications: {
        default: Authentication;
        BearerAuth: ApiKeyAuth;
    };
    protected interceptors: Interceptor[];
    constructor(basePath?: string);
    set useQuerystring(value: boolean);
    set basePath(basePath: string);
    set defaultHeaders(defaultHeaders: any);
    get defaultHeaders(): any;
    get basePath(): string;
    setDefaultAuthentication(auth: Authentication): void;
    setApiKey(key: SecretSvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Decrypt a value and return the encrypted result
     * @summary Decrypt a Value
     * @param body Decrypt Value Request
     */
    decryptValue(body: SecretSvcDecryptValueRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: SecretSvcDecryptValueResponse;
    }>;
    /**
     * Encrypt a value and return the encrypted result
     * @summary Encrypt a Value
     * @param body Encrypt Value Request
     */
    encryptValue(body: SecretSvcEncryptValueRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: SecretSvcEncryptValueResponse;
    }>;
    /**
     * Returns true if the encryption key is sufficiently secure.
     * @summary Check Security Status
     */
    isSecure(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: SecretSvcIsSecureResponse;
    }>;
    /**
     * List secrets by key(s) if authorized.
     * @summary List Secrets
     * @param body List Secret Request
     */
    listSecrets(body?: SecretSvcListSecretsRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: SecretSvcListSecretsResponse;
    }>;
    /**
     * Remove secrets if authorized to do so
     * @summary Remove Secrets
     * @param body Remove Secret Request
     */
    removeSecrets(body: SecretSvcRemoveSecretsRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
    /**
     * Save secrets if authorized to do so
     * @summary Save Secrets
     * @param body Save Secret Request
     */
    saveSecrets(body: SecretSvcSaveSecretsRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
}
