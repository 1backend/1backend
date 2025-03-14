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
import type { ModelSvcPlatform } from './ModelSvcPlatform';
/**
 *
 * @export
 * @interface ModelSvcListPlatformsResponse
 */
export interface ModelSvcListPlatformsResponse {
    /**
     *
     * @type {Array<ModelSvcPlatform>}
     * @memberof ModelSvcListPlatformsResponse
     */
    platforms: Array<ModelSvcPlatform>;
}
/**
 * Check if a given object implements the ModelSvcListPlatformsResponse interface.
 */
export declare function instanceOfModelSvcListPlatformsResponse(value: object): value is ModelSvcListPlatformsResponse;
export declare function ModelSvcListPlatformsResponseFromJSON(json: any): ModelSvcListPlatformsResponse;
export declare function ModelSvcListPlatformsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcListPlatformsResponse;
export declare function ModelSvcListPlatformsResponseToJSON(json: any): ModelSvcListPlatformsResponse;
export declare function ModelSvcListPlatformsResponseToJSONTyped(value?: ModelSvcListPlatformsResponse | null, ignoreDiscriminator?: boolean): any;
