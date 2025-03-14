/* tslint:disable */
/* eslint-disable */
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
/**
 * Check if a given object implements the SecretSvcDecryptValueRequest interface.
 */
export function instanceOfSecretSvcDecryptValueRequest(value) {
    return true;
}
export function SecretSvcDecryptValueRequestFromJSON(json) {
    return SecretSvcDecryptValueRequestFromJSONTyped(json, false);
}
export function SecretSvcDecryptValueRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'value': json['value'] == null ? undefined : json['value'],
        'values': json['values'] == null ? undefined : json['values'],
    };
}
export function SecretSvcDecryptValueRequestToJSON(json) {
    return SecretSvcDecryptValueRequestToJSONTyped(json, false);
}
export function SecretSvcDecryptValueRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'value': value['value'],
        'values': value['values'],
    };
}
