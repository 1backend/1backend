/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface SecretSvcListSecretsRequest
 */
export interface SecretSvcListSecretsRequest {
    /**
     *
     * @type {string}
     * @memberof SecretSvcListSecretsRequest
     */
    key?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof SecretSvcListSecretsRequest
     */
    keys?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof SecretSvcListSecretsRequest
     */
    namespace?: string;
}
/**
 * Check if a given object implements the SecretSvcListSecretsRequest interface.
 */
export declare function instanceOfSecretSvcListSecretsRequest(value: object): value is SecretSvcListSecretsRequest;
export declare function SecretSvcListSecretsRequestFromJSON(json: any): SecretSvcListSecretsRequest;
export declare function SecretSvcListSecretsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcListSecretsRequest;
export declare function SecretSvcListSecretsRequestToJSON(json: any): SecretSvcListSecretsRequest;
export declare function SecretSvcListSecretsRequestToJSONTyped(value?: SecretSvcListSecretsRequest | null, ignoreDiscriminator?: boolean): any;
