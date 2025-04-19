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
 * @interface PromptSvcErrorResponse
 */
export interface PromptSvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof PromptSvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the PromptSvcErrorResponse interface.
 */
export declare function instanceOfPromptSvcErrorResponse(value: object): value is PromptSvcErrorResponse;
export declare function PromptSvcErrorResponseFromJSON(json: any): PromptSvcErrorResponse;
export declare function PromptSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcErrorResponse;
export declare function PromptSvcErrorResponseToJSON(json: any): PromptSvcErrorResponse;
export declare function PromptSvcErrorResponseToJSONTyped(value?: PromptSvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
