/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface SecretSvcIsSecureResponse
 */
export interface SecretSvcIsSecureResponse {
    /**
     *
     * @type {boolean}
     * @memberof SecretSvcIsSecureResponse
     */
    isSecure: boolean;
}
/**
 * Check if a given object implements the SecretSvcIsSecureResponse interface.
 */
export declare function instanceOfSecretSvcIsSecureResponse(value: object): value is SecretSvcIsSecureResponse;
export declare function SecretSvcIsSecureResponseFromJSON(json: any): SecretSvcIsSecureResponse;
export declare function SecretSvcIsSecureResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcIsSecureResponse;
export declare function SecretSvcIsSecureResponseToJSON(json: any): SecretSvcIsSecureResponse;
export declare function SecretSvcIsSecureResponseToJSONTyped(value?: SecretSvcIsSecureResponse | null, ignoreDiscriminator?: boolean): any;
