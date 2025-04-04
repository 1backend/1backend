/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DeploySvcDeploymentFromJSON, DeploySvcDeploymentToJSON, } from './DeploySvcDeployment';
/**
 * Check if a given object implements the DeploySvcListDeploymentsResponse interface.
 */
export function instanceOfDeploySvcListDeploymentsResponse(value) {
    return true;
}
export function DeploySvcListDeploymentsResponseFromJSON(json) {
    return DeploySvcListDeploymentsResponseFromJSONTyped(json, false);
}
export function DeploySvcListDeploymentsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'deployments': json['deployments'] == null ? undefined : (json['deployments'].map(DeploySvcDeploymentFromJSON)),
    };
}
export function DeploySvcListDeploymentsResponseToJSON(json) {
    return DeploySvcListDeploymentsResponseToJSONTyped(json, false);
}
export function DeploySvcListDeploymentsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'deployments': value['deployments'] == null ? undefined : (value['deployments'].map(DeploySvcDeploymentToJSON)),
    };
}
