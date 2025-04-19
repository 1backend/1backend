/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { DeploySvcDeployment } from './DeploySvcDeployment';
import {
    DeploySvcDeploymentFromJSON,
    DeploySvcDeploymentFromJSONTyped,
    DeploySvcDeploymentToJSON,
    DeploySvcDeploymentToJSONTyped,
} from './DeploySvcDeployment';

/**
 * 
 * @export
 * @interface DeploySvcSaveDeploymentRequest
 */
export interface DeploySvcSaveDeploymentRequest {
    /**
     * 
     * @type {DeploySvcDeployment}
     * @memberof DeploySvcSaveDeploymentRequest
     */
    deployment?: DeploySvcDeployment;
}

/**
 * Check if a given object implements the DeploySvcSaveDeploymentRequest interface.
 */
export function instanceOfDeploySvcSaveDeploymentRequest(value: object): value is DeploySvcSaveDeploymentRequest {
    return true;
}

export function DeploySvcSaveDeploymentRequestFromJSON(json: any): DeploySvcSaveDeploymentRequest {
    return DeploySvcSaveDeploymentRequestFromJSONTyped(json, false);
}

export function DeploySvcSaveDeploymentRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcSaveDeploymentRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'deployment': json['deployment'] == null ? undefined : DeploySvcDeploymentFromJSON(json['deployment']),
    };
}

export function DeploySvcSaveDeploymentRequestToJSON(json: any): DeploySvcSaveDeploymentRequest {
    return DeploySvcSaveDeploymentRequestToJSONTyped(json, false);
}

export function DeploySvcSaveDeploymentRequestToJSONTyped(value?: DeploySvcSaveDeploymentRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'deployment': DeploySvcDeploymentToJSON(value['deployment']),
    };
}

