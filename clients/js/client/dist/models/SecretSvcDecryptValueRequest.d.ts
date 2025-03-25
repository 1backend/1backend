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
 * @interface SecretSvcDecryptValueRequest
 */
export interface SecretSvcDecryptValueRequest {
    /**
     *
     * @type {string}
     * @memberof SecretSvcDecryptValueRequest
     */
    value?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof SecretSvcDecryptValueRequest
     */
    values?: Array<string>;
}
/**
 * Check if a given object implements the SecretSvcDecryptValueRequest interface.
 */
export declare function instanceOfSecretSvcDecryptValueRequest(value: object): value is SecretSvcDecryptValueRequest;
export declare function SecretSvcDecryptValueRequestFromJSON(json: any): SecretSvcDecryptValueRequest;
export declare function SecretSvcDecryptValueRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcDecryptValueRequest;
export declare function SecretSvcDecryptValueRequestToJSON(json: any): SecretSvcDecryptValueRequest;
export declare function SecretSvcDecryptValueRequestToJSONTyped(value?: SecretSvcDecryptValueRequest | null, ignoreDiscriminator?: boolean): any;
