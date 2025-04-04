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
import type { ModelSvcModel } from './ModelSvcModel';
import type { ModelSvcPlatform } from './ModelSvcPlatform';
/**
 *
 * @export
 * @interface ModelSvcGetModelResponse
 */
export interface ModelSvcGetModelResponse {
    /**
     *
     * @type {boolean}
     * @memberof ModelSvcGetModelResponse
     */
    _exists: boolean;
    /**
     *
     * @type {ModelSvcModel}
     * @memberof ModelSvcGetModelResponse
     */
    model: ModelSvcModel;
    /**
     *
     * @type {ModelSvcPlatform}
     * @memberof ModelSvcGetModelResponse
     */
    platform: ModelSvcPlatform;
}
/**
 * Check if a given object implements the ModelSvcGetModelResponse interface.
 */
export declare function instanceOfModelSvcGetModelResponse(value: object): value is ModelSvcGetModelResponse;
export declare function ModelSvcGetModelResponseFromJSON(json: any): ModelSvcGetModelResponse;
export declare function ModelSvcGetModelResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcGetModelResponse;
export declare function ModelSvcGetModelResponseToJSON(json: any): ModelSvcGetModelResponse;
export declare function ModelSvcGetModelResponseToJSONTyped(value?: ModelSvcGetModelResponse | null, ignoreDiscriminator?: boolean): any;
