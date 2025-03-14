/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { PolicySvcInstance } from './PolicySvcInstance';
/**
 *
 * @export
 * @interface PolicySvcUpsertInstanceRequest
 */
export interface PolicySvcUpsertInstanceRequest {
    /**
     *
     * @type {PolicySvcInstance}
     * @memberof PolicySvcUpsertInstanceRequest
     */
    instance?: PolicySvcInstance;
}
/**
 * Check if a given object implements the PolicySvcUpsertInstanceRequest interface.
 */
export declare function instanceOfPolicySvcUpsertInstanceRequest(value: object): value is PolicySvcUpsertInstanceRequest;
export declare function PolicySvcUpsertInstanceRequestFromJSON(json: any): PolicySvcUpsertInstanceRequest;
export declare function PolicySvcUpsertInstanceRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcUpsertInstanceRequest;
export declare function PolicySvcUpsertInstanceRequestToJSON(json: any): PolicySvcUpsertInstanceRequest;
export declare function PolicySvcUpsertInstanceRequestToJSONTyped(value?: PolicySvcUpsertInstanceRequest | null, ignoreDiscriminator?: boolean): any;
