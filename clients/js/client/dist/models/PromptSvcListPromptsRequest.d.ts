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
import type { DatastoreQuery } from './DatastoreQuery';
/**
 *
 * @export
 * @interface PromptSvcListPromptsRequest
 */
export interface PromptSvcListPromptsRequest {
    /**
     *
     * @type {DatastoreQuery}
     * @memberof PromptSvcListPromptsRequest
     */
    query?: DatastoreQuery;
}
/**
 * Check if a given object implements the PromptSvcListPromptsRequest interface.
 */
export declare function instanceOfPromptSvcListPromptsRequest(value: object): value is PromptSvcListPromptsRequest;
export declare function PromptSvcListPromptsRequestFromJSON(json: any): PromptSvcListPromptsRequest;
export declare function PromptSvcListPromptsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcListPromptsRequest;
export declare function PromptSvcListPromptsRequestToJSON(json: any): PromptSvcListPromptsRequest;
export declare function PromptSvcListPromptsRequestToJSONTyped(value?: PromptSvcListPromptsRequest | null, ignoreDiscriminator?: boolean): any;
