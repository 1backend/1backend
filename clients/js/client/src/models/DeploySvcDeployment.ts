/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { DeploySvcDeploymentStrategy } from './DeploySvcDeploymentStrategy';
import {
    DeploySvcDeploymentStrategyFromJSON,
    DeploySvcDeploymentStrategyFromJSONTyped,
    DeploySvcDeploymentStrategyToJSON,
    DeploySvcDeploymentStrategyToJSONTyped,
} from './DeploySvcDeploymentStrategy';
import type { DeploySvcAutoScalingConfig } from './DeploySvcAutoScalingConfig';
import {
    DeploySvcAutoScalingConfigFromJSON,
    DeploySvcAutoScalingConfigFromJSONTyped,
    DeploySvcAutoScalingConfigToJSON,
    DeploySvcAutoScalingConfigToJSONTyped,
} from './DeploySvcAutoScalingConfig';
import type { DeploySvcDeploymentStatus } from './DeploySvcDeploymentStatus';
import {
    DeploySvcDeploymentStatusFromJSON,
    DeploySvcDeploymentStatusFromJSONTyped,
    DeploySvcDeploymentStatusToJSON,
    DeploySvcDeploymentStatusToJSONTyped,
} from './DeploySvcDeploymentStatus';
import type { DeploySvcTargetRegion } from './DeploySvcTargetRegion';
import {
    DeploySvcTargetRegionFromJSON,
    DeploySvcTargetRegionFromJSONTyped,
    DeploySvcTargetRegionToJSON,
    DeploySvcTargetRegionToJSONTyped,
} from './DeploySvcTargetRegion';
import type { DeploySvcResourceLimits } from './DeploySvcResourceLimits';
import {
    DeploySvcResourceLimitsFromJSON,
    DeploySvcResourceLimitsFromJSONTyped,
    DeploySvcResourceLimitsToJSON,
    DeploySvcResourceLimitsToJSONTyped,
} from './DeploySvcResourceLimits';

/**
 * 
 * @export
 * @interface DeploySvcDeployment
 */
export interface DeploySvcDeployment {
    /**
     * Optional: Auto-scaling rules
     * @type {DeploySvcAutoScalingConfig}
     * @memberof DeploySvcDeployment
     */
    autoScaling?: DeploySvcAutoScalingConfig;
    /**
     * DefinitionId is the id of the definition
     * @type {string}
     * @memberof DeploySvcDeployment
     */
    definitionId: string;
    /**
     * Description of what this deployment does
     * @type {string}
     * @memberof DeploySvcDeployment
     */
    description?: string;
    /**
     * Details provides additional information about the deployment's current state,
     * including both success and failure conditions (e.g., "Deployment in progress", "Error pulling image").
     * @type {string}
     * @memberof DeploySvcDeployment
     */
    details?: string;
    /**
     * Envars is a map of environment variables that will be passed down to service instances (see Registry Svc Instance)
     * Also see the Registry Svc Definition for required envars.
     * @type {{ [key: string]: string; }}
     * @memberof DeploySvcDeployment
     */
    envars?: { [key: string]: string; };
    /**
     * ID of the deployment (e.g., "depl_dbOdi5eLQK")
     * @type {string}
     * @memberof DeploySvcDeployment
     */
    id: string;
    /**
     * Short name for easy reference (e.g., "user-service-v2")
     * @type {string}
     * @memberof DeploySvcDeployment
     */
    name?: string;
    /**
     * Number of container instances to run
     * @type {number}
     * @memberof DeploySvcDeployment
     */
    replicas?: number;
    /**
     * Resource requirements for each replica
     * @type {DeploySvcResourceLimits}
     * @memberof DeploySvcDeployment
     */
    resources?: DeploySvcResourceLimits;
    /**
     * Current status of the deployment (e.g., "OK", "Error", "Pending")
     * @type {DeploySvcDeploymentStatus}
     * @memberof DeploySvcDeployment
     */
    status?: DeploySvcDeploymentStatus;
    /**
     * Deployment strategy (e.g., rolling update)
     * @type {DeploySvcDeploymentStrategy}
     * @memberof DeploySvcDeployment
     */
    strategy?: DeploySvcDeploymentStrategy;
    /**
     * Target deployment regions or clusters
     * @type {Array<DeploySvcTargetRegion>}
     * @memberof DeploySvcDeployment
     */
    targetRegions?: Array<DeploySvcTargetRegion>;
}



/**
 * Check if a given object implements the DeploySvcDeployment interface.
 */
export function instanceOfDeploySvcDeployment(value: object): value is DeploySvcDeployment {
    if (!('definitionId' in value) || value['definitionId'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    return true;
}

export function DeploySvcDeploymentFromJSON(json: any): DeploySvcDeployment {
    return DeploySvcDeploymentFromJSONTyped(json, false);
}

export function DeploySvcDeploymentFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcDeployment {
    if (json == null) {
        return json;
    }
    return {
        
        'autoScaling': json['autoScaling'] == null ? undefined : DeploySvcAutoScalingConfigFromJSON(json['autoScaling']),
        'definitionId': json['definitionId'],
        'description': json['description'] == null ? undefined : json['description'],
        'details': json['details'] == null ? undefined : json['details'],
        'envars': json['envars'] == null ? undefined : json['envars'],
        'id': json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'replicas': json['replicas'] == null ? undefined : json['replicas'],
        'resources': json['resources'] == null ? undefined : DeploySvcResourceLimitsFromJSON(json['resources']),
        'status': json['status'] == null ? undefined : DeploySvcDeploymentStatusFromJSON(json['status']),
        'strategy': json['strategy'] == null ? undefined : DeploySvcDeploymentStrategyFromJSON(json['strategy']),
        'targetRegions': json['targetRegions'] == null ? undefined : ((json['targetRegions'] as Array<any>).map(DeploySvcTargetRegionFromJSON)),
    };
}

export function DeploySvcDeploymentToJSON(json: any): DeploySvcDeployment {
    return DeploySvcDeploymentToJSONTyped(json, false);
}

export function DeploySvcDeploymentToJSONTyped(value?: DeploySvcDeployment | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'autoScaling': DeploySvcAutoScalingConfigToJSON(value['autoScaling']),
        'definitionId': value['definitionId'],
        'description': value['description'],
        'details': value['details'],
        'envars': value['envars'],
        'id': value['id'],
        'name': value['name'],
        'replicas': value['replicas'],
        'resources': DeploySvcResourceLimitsToJSON(value['resources']),
        'status': DeploySvcDeploymentStatusToJSON(value['status']),
        'strategy': DeploySvcDeploymentStrategyToJSON(value['strategy']),
        'targetRegions': value['targetRegions'] == null ? undefined : ((value['targetRegions'] as Array<any>).map(DeploySvcTargetRegionToJSON)),
    };
}

