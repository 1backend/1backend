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
 * @interface PromptSvcRemovePromptRequest
 */
export interface PromptSvcRemovePromptRequest {
    /**
     *
     * @type {string}
     * @memberof PromptSvcRemovePromptRequest
     */
    promptId?: string;
}
/**
 * Check if a given object implements the PromptSvcRemovePromptRequest interface.
 */
export declare function instanceOfPromptSvcRemovePromptRequest(value: object): value is PromptSvcRemovePromptRequest;
export declare function PromptSvcRemovePromptRequestFromJSON(json: any): PromptSvcRemovePromptRequest;
export declare function PromptSvcRemovePromptRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcRemovePromptRequest;
export declare function PromptSvcRemovePromptRequestToJSON(json: any): PromptSvcRemovePromptRequest;
export declare function PromptSvcRemovePromptRequestToJSONTyped(value?: PromptSvcRemovePromptRequest | null, ignoreDiscriminator?: boolean): any;
