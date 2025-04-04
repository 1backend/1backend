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
 * @interface ModelSvcErrorResponse
 */
export interface ModelSvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof ModelSvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the ModelSvcErrorResponse interface.
 */
export declare function instanceOfModelSvcErrorResponse(value: object): value is ModelSvcErrorResponse;
export declare function ModelSvcErrorResponseFromJSON(json: any): ModelSvcErrorResponse;
export declare function ModelSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcErrorResponse;
export declare function ModelSvcErrorResponseToJSON(json: any): ModelSvcErrorResponse;
export declare function ModelSvcErrorResponseToJSONTyped(value?: ModelSvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
