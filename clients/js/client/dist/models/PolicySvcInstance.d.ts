/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { PolicySvcTemplateId } from './PolicySvcTemplateId';
import type { PolicySvcParameters } from './PolicySvcParameters';
/**
 *
 * @export
 * @interface PolicySvcInstance
 */
export interface PolicySvcInstance {
    /**
     *
     * @type {string}
     * @memberof PolicySvcInstance
     */
    endpoint?: string;
    /**
     *
     * @type {string}
     * @memberof PolicySvcInstance
     */
    id?: string;
    /**
     *
     * @type {PolicySvcParameters}
     * @memberof PolicySvcInstance
     */
    parameters: PolicySvcParameters;
    /**
     *
     * @type {PolicySvcTemplateId}
     * @memberof PolicySvcInstance
     */
    templateId: PolicySvcTemplateId;
}
/**
 * Check if a given object implements the PolicySvcInstance interface.
 */
export declare function instanceOfPolicySvcInstance(value: object): value is PolicySvcInstance;
export declare function PolicySvcInstanceFromJSON(json: any): PolicySvcInstance;
export declare function PolicySvcInstanceFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcInstance;
export declare function PolicySvcInstanceToJSON(json: any): PolicySvcInstance;
export declare function PolicySvcInstanceToJSONTyped(value?: PolicySvcInstance | null, ignoreDiscriminator?: boolean): any;
