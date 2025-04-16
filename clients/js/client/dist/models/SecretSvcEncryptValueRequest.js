/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the SecretSvcEncryptValueRequest interface.
 */
export function instanceOfSecretSvcEncryptValueRequest(value) {
    return true;
}
export function SecretSvcEncryptValueRequestFromJSON(json) {
    return SecretSvcEncryptValueRequestFromJSONTyped(json, false);
}
export function SecretSvcEncryptValueRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'value': json['value'] == null ? undefined : json['value'],
        'values': json['values'] == null ? undefined : json['values'],
    };
}
export function SecretSvcEncryptValueRequestToJSON(json) {
    return SecretSvcEncryptValueRequestToJSONTyped(json, false);
}
export function SecretSvcEncryptValueRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'value': value['value'],
        'values': value['values'],
    };
}
