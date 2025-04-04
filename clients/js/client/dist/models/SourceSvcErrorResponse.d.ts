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
 * @interface SourceSvcErrorResponse
 */
export interface SourceSvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof SourceSvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the SourceSvcErrorResponse interface.
 */
export declare function instanceOfSourceSvcErrorResponse(value: object): value is SourceSvcErrorResponse;
export declare function SourceSvcErrorResponseFromJSON(json: any): SourceSvcErrorResponse;
export declare function SourceSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SourceSvcErrorResponse;
export declare function SourceSvcErrorResponseToJSON(json: any): SourceSvcErrorResponse;
export declare function SourceSvcErrorResponseToJSONTyped(value?: SourceSvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
