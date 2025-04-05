/* tslint:disable */
/* eslint-disable */
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

import { mapValues } from '../runtime';
import type { PolicySvcTemplateId } from './PolicySvcTemplateId';
import {
    PolicySvcTemplateIdFromJSON,
    PolicySvcTemplateIdFromJSONTyped,
    PolicySvcTemplateIdToJSON,
    PolicySvcTemplateIdToJSONTyped,
} from './PolicySvcTemplateId';
import type { PolicySvcParameters } from './PolicySvcParameters';
import {
    PolicySvcParametersFromJSON,
    PolicySvcParametersFromJSONTyped,
    PolicySvcParametersToJSON,
    PolicySvcParametersToJSONTyped,
} from './PolicySvcParameters';

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
export function instanceOfPolicySvcInstance(value: object): value is PolicySvcInstance {
    if (!('parameters' in value) || value['parameters'] === undefined) return false;
    if (!('templateId' in value) || value['templateId'] === undefined) return false;
    return true;
}

export function PolicySvcInstanceFromJSON(json: any): PolicySvcInstance {
    return PolicySvcInstanceFromJSONTyped(json, false);
}

export function PolicySvcInstanceFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcInstance {
    if (json == null) {
        return json;
    }
    return {
        
        'endpoint': json['endpoint'] == null ? undefined : json['endpoint'],
        'id': json['id'] == null ? undefined : json['id'],
        'parameters': PolicySvcParametersFromJSON(json['parameters']),
        'templateId': PolicySvcTemplateIdFromJSON(json['templateId']),
    };
}

export function PolicySvcInstanceToJSON(json: any): PolicySvcInstance {
    return PolicySvcInstanceToJSONTyped(json, false);
}

export function PolicySvcInstanceToJSONTyped(value?: PolicySvcInstance | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'endpoint': value['endpoint'],
        'id': value['id'],
        'parameters': PolicySvcParametersToJSON(value['parameters']),
        'templateId': PolicySvcTemplateIdToJSON(value['templateId']),
    };
}

