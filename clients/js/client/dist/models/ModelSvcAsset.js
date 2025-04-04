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
/**
 * Check if a given object implements the ModelSvcAsset interface.
 */
export function instanceOfModelSvcAsset(value) {
    if (!('envVarKey' in value) || value['envVarKey'] === undefined)
        return false;
    if (!('url' in value) || value['url'] === undefined)
        return false;
    return true;
}
export function ModelSvcAssetFromJSON(json) {
    return ModelSvcAssetFromJSONTyped(json, false);
}
export function ModelSvcAssetFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'envVarKey': json['envVarKey'],
        'url': json['url'],
    };
}
export function ModelSvcAssetToJSON(json) {
    return ModelSvcAssetToJSONTyped(json, false);
}
export function ModelSvcAssetToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'envVarKey': value['envVarKey'],
        'url': value['url'],
    };
}
