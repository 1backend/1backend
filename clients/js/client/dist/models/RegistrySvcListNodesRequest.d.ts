/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface RegistrySvcListNodesRequest
 */
export interface RegistrySvcListNodesRequest {
    /**
     * Node IDs to filter on
     * @type {Array<string>}
     * @memberof RegistrySvcListNodesRequest
     */
    ids?: Array<string>;
}
/**
 * Check if a given object implements the RegistrySvcListNodesRequest interface.
 */
export declare function instanceOfRegistrySvcListNodesRequest(value: object): value is RegistrySvcListNodesRequest;
export declare function RegistrySvcListNodesRequestFromJSON(json: any): RegistrySvcListNodesRequest;
export declare function RegistrySvcListNodesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcListNodesRequest;
export declare function RegistrySvcListNodesRequestToJSON(json: any): RegistrySvcListNodesRequest;
export declare function RegistrySvcListNodesRequestToJSONTyped(value?: RegistrySvcListNodesRequest | null, ignoreDiscriminator?: boolean): any;
