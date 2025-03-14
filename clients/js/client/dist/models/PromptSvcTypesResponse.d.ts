/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { PromptSvcStreamChunk } from './PromptSvcStreamChunk';
/**
 *
 * @export
 * @interface PromptSvcTypesResponse
 */
export interface PromptSvcTypesResponse {
    /**
     *
     * @type {PromptSvcStreamChunk}
     * @memberof PromptSvcTypesResponse
     */
    chunk?: PromptSvcStreamChunk;
}
/**
 * Check if a given object implements the PromptSvcTypesResponse interface.
 */
export declare function instanceOfPromptSvcTypesResponse(value: object): value is PromptSvcTypesResponse;
export declare function PromptSvcTypesResponseFromJSON(json: any): PromptSvcTypesResponse;
export declare function PromptSvcTypesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcTypesResponse;
export declare function PromptSvcTypesResponseToJSON(json: any): PromptSvcTypesResponse;
export declare function PromptSvcTypesResponseToJSONTyped(value?: PromptSvcTypesResponse | null, ignoreDiscriminator?: boolean): any;
