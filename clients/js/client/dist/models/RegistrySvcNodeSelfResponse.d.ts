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
import type { RegistrySvcNode } from './RegistrySvcNode';
/**
 *
 * @export
 * @interface RegistrySvcNodeSelfResponse
 */
export interface RegistrySvcNodeSelfResponse {
    /**
     *
     * @type {RegistrySvcNode}
     * @memberof RegistrySvcNodeSelfResponse
     */
    node: RegistrySvcNode;
}
/**
 * Check if a given object implements the RegistrySvcNodeSelfResponse interface.
 */
export declare function instanceOfRegistrySvcNodeSelfResponse(value: object): value is RegistrySvcNodeSelfResponse;
export declare function RegistrySvcNodeSelfResponseFromJSON(json: any): RegistrySvcNodeSelfResponse;
export declare function RegistrySvcNodeSelfResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcNodeSelfResponse;
export declare function RegistrySvcNodeSelfResponseToJSON(json: any): RegistrySvcNodeSelfResponse;
export declare function RegistrySvcNodeSelfResponseToJSONTyped(value?: RegistrySvcNodeSelfResponse | null, ignoreDiscriminator?: boolean): any;
