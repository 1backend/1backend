/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { RegistrySvcDefinition } from './RegistrySvcDefinition';
/**
 *
 * @export
 * @interface RegistrySvcListDefinitionsResponse
 */
export interface RegistrySvcListDefinitionsResponse {
    /**
     *
     * @type {Array<RegistrySvcDefinition>}
     * @memberof RegistrySvcListDefinitionsResponse
     */
    definitions?: Array<RegistrySvcDefinition>;
}
/**
 * Check if a given object implements the RegistrySvcListDefinitionsResponse interface.
 */
export declare function instanceOfRegistrySvcListDefinitionsResponse(value: object): value is RegistrySvcListDefinitionsResponse;
export declare function RegistrySvcListDefinitionsResponseFromJSON(json: any): RegistrySvcListDefinitionsResponse;
export declare function RegistrySvcListDefinitionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcListDefinitionsResponse;
export declare function RegistrySvcListDefinitionsResponseToJSON(json: any): RegistrySvcListDefinitionsResponse;
export declare function RegistrySvcListDefinitionsResponseToJSONTyped(value?: RegistrySvcListDefinitionsResponse | null, ignoreDiscriminator?: boolean): any;
