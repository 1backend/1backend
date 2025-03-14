/* tslint:disable */
/* eslint-disable */
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
 * @interface DeploySvcListDeploymentsResponse
 */
export interface DeploySvcListDeploymentsResponse {
    /**
     * 
     * @type {Array<DeploySvcDeployment>}
     * @memberof DeploySvcListDeploymentsResponse
     */
    deployments?: Array<DeploySvcDeployment>;
}

/**
 * Check if a given object implements the DeploySvcListDeploymentsResponse interface.
 */
export function instanceOfDeploySvcListDeploymentsResponse(value: object): value is DeploySvcListDeploymentsResponse {
    return true;
}

export function DeploySvcListDeploymentsResponseFromJSON(json: any): DeploySvcListDeploymentsResponse {
    return DeploySvcListDeploymentsResponseFromJSONTyped(json, false);
}

export function DeploySvcListDeploymentsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcListDeploymentsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'deployments': json['deployments'] == null ? undefined : ((json['deployments'] as Array<any>).map(DeploySvcDeploymentFromJSON)),
    };
}

export function DeploySvcListDeploymentsResponseToJSON(json: any): DeploySvcListDeploymentsResponse {
    return DeploySvcListDeploymentsResponseToJSONTyped(json, false);
}

export function DeploySvcListDeploymentsResponseToJSONTyped(value?: DeploySvcListDeploymentsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'deployments': value['deployments'] == null ? undefined : ((value['deployments'] as Array<any>).map(DeploySvcDeploymentToJSON)),
    };
}

