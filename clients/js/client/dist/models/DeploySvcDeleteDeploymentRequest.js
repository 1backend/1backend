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
/**
 * Check if a given object implements the DeploySvcDeleteDeploymentRequest interface.
 */
export function instanceOfDeploySvcDeleteDeploymentRequest(value) {
    if (!('deploymentId' in value) || value['deploymentId'] === undefined)
        return false;
    return true;
}
export function DeploySvcDeleteDeploymentRequestFromJSON(json) {
    return DeploySvcDeleteDeploymentRequestFromJSONTyped(json, false);
}
export function DeploySvcDeleteDeploymentRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'deploymentId': json['deploymentId'],
    };
}
export function DeploySvcDeleteDeploymentRequestToJSON(json) {
    return DeploySvcDeleteDeploymentRequestToJSONTyped(json, false);
}
export function DeploySvcDeleteDeploymentRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'deploymentId': value['deploymentId'],
    };
}
