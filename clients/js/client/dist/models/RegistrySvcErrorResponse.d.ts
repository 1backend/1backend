/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface RegistrySvcErrorResponse
 */
export interface RegistrySvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof RegistrySvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the RegistrySvcErrorResponse interface.
 */
export declare function instanceOfRegistrySvcErrorResponse(value: object): value is RegistrySvcErrorResponse;
export declare function RegistrySvcErrorResponseFromJSON(json: any): RegistrySvcErrorResponse;
export declare function RegistrySvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcErrorResponse;
export declare function RegistrySvcErrorResponseToJSON(json: any): RegistrySvcErrorResponse;
export declare function RegistrySvcErrorResponseToJSONTyped(value?: RegistrySvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
