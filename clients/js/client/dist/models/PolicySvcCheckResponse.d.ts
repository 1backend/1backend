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
 * @interface PolicySvcCheckResponse
 */
export interface PolicySvcCheckResponse {
    /**
     *
     * @type {boolean}
     * @memberof PolicySvcCheckResponse
     */
    allowed: boolean;
}
/**
 * Check if a given object implements the PolicySvcCheckResponse interface.
 */
export declare function instanceOfPolicySvcCheckResponse(value: object): value is PolicySvcCheckResponse;
export declare function PolicySvcCheckResponseFromJSON(json: any): PolicySvcCheckResponse;
export declare function PolicySvcCheckResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcCheckResponse;
export declare function PolicySvcCheckResponseToJSON(json: any): PolicySvcCheckResponse;
export declare function PolicySvcCheckResponseToJSONTyped(value?: PolicySvcCheckResponse | null, ignoreDiscriminator?: boolean): any;
