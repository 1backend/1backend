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
import type { RegistrySvcNode } from './RegistrySvcNode';
/**
 *
 * @export
 * @interface RegistrySvcListNodesResponse
 */
export interface RegistrySvcListNodesResponse {
    /**
     *
     * @type {Array<RegistrySvcNode>}
     * @memberof RegistrySvcListNodesResponse
     */
    nodes?: Array<RegistrySvcNode>;
}
/**
 * Check if a given object implements the RegistrySvcListNodesResponse interface.
 */
export declare function instanceOfRegistrySvcListNodesResponse(value: object): value is RegistrySvcListNodesResponse;
export declare function RegistrySvcListNodesResponseFromJSON(json: any): RegistrySvcListNodesResponse;
export declare function RegistrySvcListNodesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcListNodesResponse;
export declare function RegistrySvcListNodesResponseToJSON(json: any): RegistrySvcListNodesResponse;
export declare function RegistrySvcListNodesResponseToJSONTyped(value?: RegistrySvcListNodesResponse | null, ignoreDiscriminator?: boolean): any;
