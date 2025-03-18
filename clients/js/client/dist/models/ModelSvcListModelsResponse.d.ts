/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ModelSvcModel } from './ModelSvcModel';
/**
 *
 * @export
 * @interface ModelSvcListModelsResponse
 */
export interface ModelSvcListModelsResponse {
    /**
     *
     * @type {Array<ModelSvcModel>}
     * @memberof ModelSvcListModelsResponse
     */
    models?: Array<ModelSvcModel>;
}
/**
 * Check if a given object implements the ModelSvcListModelsResponse interface.
 */
export declare function instanceOfModelSvcListModelsResponse(value: object): value is ModelSvcListModelsResponse;
export declare function ModelSvcListModelsResponseFromJSON(json: any): ModelSvcListModelsResponse;
export declare function ModelSvcListModelsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcListModelsResponse;
export declare function ModelSvcListModelsResponseToJSON(json: any): ModelSvcListModelsResponse;
export declare function ModelSvcListModelsResponseToJSONTyped(value?: ModelSvcListModelsResponse | null, ignoreDiscriminator?: boolean): any;
