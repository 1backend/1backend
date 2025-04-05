/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { RegistrySvcInstance } from './RegistrySvcInstance';
/**
 *
 * @export
 * @interface RegistrySvcListInstancesResponse
 */
export interface RegistrySvcListInstancesResponse {
    /**
     *
     * @type {Array<RegistrySvcInstance>}
     * @memberof RegistrySvcListInstancesResponse
     */
    instances?: Array<RegistrySvcInstance>;
}
/**
 * Check if a given object implements the RegistrySvcListInstancesResponse interface.
 */
export declare function instanceOfRegistrySvcListInstancesResponse(value: object): value is RegistrySvcListInstancesResponse;
export declare function RegistrySvcListInstancesResponseFromJSON(json: any): RegistrySvcListInstancesResponse;
export declare function RegistrySvcListInstancesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcListInstancesResponse;
export declare function RegistrySvcListInstancesResponseToJSON(json: any): RegistrySvcListInstancesResponse;
export declare function RegistrySvcListInstancesResponseToJSONTyped(value?: RegistrySvcListInstancesResponse | null, ignoreDiscriminator?: boolean): any;
