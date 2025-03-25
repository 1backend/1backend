/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface SecretSvcEncryptValueRequest
 */
export interface SecretSvcEncryptValueRequest {
    /**
     *
     * @type {string}
     * @memberof SecretSvcEncryptValueRequest
     */
    value?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof SecretSvcEncryptValueRequest
     */
    values?: Array<string>;
}
/**
 * Check if a given object implements the SecretSvcEncryptValueRequest interface.
 */
export declare function instanceOfSecretSvcEncryptValueRequest(value: object): value is SecretSvcEncryptValueRequest;
export declare function SecretSvcEncryptValueRequestFromJSON(json: any): SecretSvcEncryptValueRequest;
export declare function SecretSvcEncryptValueRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcEncryptValueRequest;
export declare function SecretSvcEncryptValueRequestToJSON(json: any): SecretSvcEncryptValueRequest;
export declare function SecretSvcEncryptValueRequestToJSONTyped(value?: SecretSvcEncryptValueRequest | null, ignoreDiscriminator?: boolean): any;
