/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ModelSvcModelStatus } from './ModelSvcModelStatus';
/**
 *
 * @export
 * @interface ModelSvcStatusResponse
 */
export interface ModelSvcStatusResponse {
    /**
     *
     * @type {ModelSvcModelStatus}
     * @memberof ModelSvcStatusResponse
     */
    status?: ModelSvcModelStatus;
}
/**
 * Check if a given object implements the ModelSvcStatusResponse interface.
 */
export declare function instanceOfModelSvcStatusResponse(value: object): value is ModelSvcStatusResponse;
export declare function ModelSvcStatusResponseFromJSON(json: any): ModelSvcStatusResponse;
export declare function ModelSvcStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcStatusResponse;
export declare function ModelSvcStatusResponseToJSON(json: any): ModelSvcStatusResponse;
export declare function ModelSvcStatusResponseToJSONTyped(value?: ModelSvcStatusResponse | null, ignoreDiscriminator?: boolean): any;
