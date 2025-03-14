/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  SecretSvcDecryptValueRequest,
  SecretSvcDecryptValueResponse,
  SecretSvcEncryptValueRequest,
  SecretSvcEncryptValueResponse,
  SecretSvcIsSecureResponse,
  SecretSvcListSecretsRequest,
  SecretSvcListSecretsResponse,
  SecretSvcRemoveSecretsRequest,
  SecretSvcSaveSecretsRequest,
} from '../models/index';
import {
    SecretSvcDecryptValueRequestFromJSON,
    SecretSvcDecryptValueRequestToJSON,
    SecretSvcDecryptValueResponseFromJSON,
    SecretSvcDecryptValueResponseToJSON,
    SecretSvcEncryptValueRequestFromJSON,
    SecretSvcEncryptValueRequestToJSON,
    SecretSvcEncryptValueResponseFromJSON,
    SecretSvcEncryptValueResponseToJSON,
    SecretSvcIsSecureResponseFromJSON,
    SecretSvcIsSecureResponseToJSON,
    SecretSvcListSecretsRequestFromJSON,
    SecretSvcListSecretsRequestToJSON,
    SecretSvcListSecretsResponseFromJSON,
    SecretSvcListSecretsResponseToJSON,
    SecretSvcRemoveSecretsRequestFromJSON,
    SecretSvcRemoveSecretsRequestToJSON,
    SecretSvcSaveSecretsRequestFromJSON,
    SecretSvcSaveSecretsRequestToJSON,
} from '../models/index';

export interface DecryptValueRequest {
    body: SecretSvcDecryptValueRequest;
}

export interface EncryptValueRequest {
    body: SecretSvcEncryptValueRequest;
}

export interface ListSecretsRequest {
    body?: SecretSvcListSecretsRequest;
}

export interface RemoveSecretsRequest {
    body: SecretSvcRemoveSecretsRequest;
}

export interface SaveSecretsRequest {
    body: SecretSvcSaveSecretsRequest;
}

/**
 * 
 */
export class SecretSvcApi extends runtime.BaseAPI {

    /**
     * Decrypt a value and return the encrypted result
     * Decrypt a Value
     */
    async decryptValueRaw(requestParameters: DecryptValueRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SecretSvcDecryptValueResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling decryptValue().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/secret-svc/decrypt`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SecretSvcDecryptValueRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SecretSvcDecryptValueResponseFromJSON(jsonValue));
    }

    /**
     * Decrypt a value and return the encrypted result
     * Decrypt a Value
     */
    async decryptValue(requestParameters: DecryptValueRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SecretSvcDecryptValueResponse> {
        const response = await this.decryptValueRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Encrypt a value and return the encrypted result
     * Encrypt a Value
     */
    async encryptValueRaw(requestParameters: EncryptValueRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SecretSvcEncryptValueResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling encryptValue().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/secret-svc/encrypt`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SecretSvcEncryptValueRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SecretSvcEncryptValueResponseFromJSON(jsonValue));
    }

    /**
     * Encrypt a value and return the encrypted result
     * Encrypt a Value
     */
    async encryptValue(requestParameters: EncryptValueRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SecretSvcEncryptValueResponse> {
        const response = await this.encryptValueRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Returns true if the encryption key is sufficiently secure.
     * Check Security Status
     */
    async isSecureRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SecretSvcIsSecureResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/secret-svc/is-secure`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SecretSvcIsSecureResponseFromJSON(jsonValue));
    }

    /**
     * Returns true if the encryption key is sufficiently secure.
     * Check Security Status
     */
    async isSecure(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SecretSvcIsSecureResponse> {
        const response = await this.isSecureRaw(initOverrides);
        return await response.value();
    }

    /**
     * List secrets by key(s) if authorized.
     * List Secrets
     */
    async listSecretsRaw(requestParameters: ListSecretsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SecretSvcListSecretsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/secret-svc/secrets`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SecretSvcListSecretsRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SecretSvcListSecretsResponseFromJSON(jsonValue));
    }

    /**
     * List secrets by key(s) if authorized.
     * List Secrets
     */
    async listSecrets(requestParameters: ListSecretsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SecretSvcListSecretsResponse> {
        const response = await this.listSecretsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Remove secrets if authorized to do so
     * Remove Secrets
     */
    async removeSecretsRaw(requestParameters: RemoveSecretsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling removeSecrets().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/secret-svc/secrets`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
            body: SecretSvcRemoveSecretsRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Remove secrets if authorized to do so
     * Remove Secrets
     */
    async removeSecrets(requestParameters: RemoveSecretsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.removeSecretsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Save secrets if authorized to do so
     * Save Secrets
     */
    async saveSecretsRaw(requestParameters: SaveSecretsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling saveSecrets().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/secret-svc/secrets`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: SecretSvcSaveSecretsRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Save secrets if authorized to do so
     * Save Secrets
     */
    async saveSecrets(requestParameters: SaveSecretsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.saveSecretsRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
