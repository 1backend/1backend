/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DeploySvcDeploymentStrategyFromJSON, DeploySvcDeploymentStrategyToJSON, } from './DeploySvcDeploymentStrategy';
import { DeploySvcAutoScalingConfigFromJSON, DeploySvcAutoScalingConfigToJSON, } from './DeploySvcAutoScalingConfig';
import { DeploySvcDeploymentStatusFromJSON, DeploySvcDeploymentStatusToJSON, } from './DeploySvcDeploymentStatus';
import { DeploySvcTargetRegionFromJSON, DeploySvcTargetRegionToJSON, } from './DeploySvcTargetRegion';
import { DeploySvcResourceLimitsFromJSON, DeploySvcResourceLimitsToJSON, } from './DeploySvcResourceLimits';
/**
 * Check if a given object implements the DeploySvcDeployment interface.
 */
export function instanceOfDeploySvcDeployment(value) {
    if (!('definitionId' in value) || value['definitionId'] === undefined)
        return false;
    if (!('id' in value) || value['id'] === undefined)
        return false;
    return true;
}
export function DeploySvcDeploymentFromJSON(json) {
    return DeploySvcDeploymentFromJSONTyped(json, false);
}
export function DeploySvcDeploymentFromJSONTyped(json, ignoreDiscriminator) {
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
        'targetRegions': json['targetRegions'] == null ? undefined : (json['targetRegions'].map(DeploySvcTargetRegionFromJSON)),
    };
}
export function DeploySvcDeploymentToJSON(json) {
    return DeploySvcDeploymentToJSONTyped(json, false);
}
export function DeploySvcDeploymentToJSONTyped(value, ignoreDiscriminator = false) {
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
        'targetRegions': value['targetRegions'] == null ? undefined : (value['targetRegions'].map(DeploySvcTargetRegionToJSON)),
    };
}
