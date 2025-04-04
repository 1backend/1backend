/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface PolicySvcCheckRequest
 */
export interface PolicySvcCheckRequest {
    /**
     *
     * @type {string}
     * @memberof PolicySvcCheckRequest
     */
    endpoint?: string;
    /**
     *
     * @type {string}
     * @memberof PolicySvcCheckRequest
     */
    ip?: string;
    /**
     *
     * @type {string}
     * @memberof PolicySvcCheckRequest
     */
    method?: string;
    /**
     *
     * @type {string}
     * @memberof PolicySvcCheckRequest
     */
    userId?: string;
}
/**
 * Check if a given object implements the PolicySvcCheckRequest interface.
 */
export declare function instanceOfPolicySvcCheckRequest(value: object): value is PolicySvcCheckRequest;
export declare function PolicySvcCheckRequestFromJSON(json: any): PolicySvcCheckRequest;
export declare function PolicySvcCheckRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcCheckRequest;
export declare function PolicySvcCheckRequestToJSON(json: any): PolicySvcCheckRequest;
export declare function PolicySvcCheckRequestToJSONTyped(value?: PolicySvcCheckRequest | null, ignoreDiscriminator?: boolean): any;
