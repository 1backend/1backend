/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface DeploySvcDeleteDeploymentRequest
 */
export interface DeploySvcDeleteDeploymentRequest {
    /**
     * 
     * @type {string}
     * @memberof DeploySvcDeleteDeploymentRequest
     */
    deploymentId: string;
}

/**
 * Check if a given object implements the DeploySvcDeleteDeploymentRequest interface.
 */
export function instanceOfDeploySvcDeleteDeploymentRequest(value: object): value is DeploySvcDeleteDeploymentRequest {
    if (!('deploymentId' in value) || value['deploymentId'] === undefined) return false;
    return true;
}

export function DeploySvcDeleteDeploymentRequestFromJSON(json: any): DeploySvcDeleteDeploymentRequest {
    return DeploySvcDeleteDeploymentRequestFromJSONTyped(json, false);
}

export function DeploySvcDeleteDeploymentRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcDeleteDeploymentRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'deploymentId': json['deploymentId'],
    };
}

export function DeploySvcDeleteDeploymentRequestToJSON(json: any): DeploySvcDeleteDeploymentRequest {
    return DeploySvcDeleteDeploymentRequestToJSONTyped(json, false);
}

export function DeploySvcDeleteDeploymentRequestToJSONTyped(value?: DeploySvcDeleteDeploymentRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'deploymentId': value['deploymentId'],
    };
}

