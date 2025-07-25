/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the SecretSvcDecryptValueResponse interface.
 */
export function instanceOfSecretSvcDecryptValueResponse(value) {
    return true;
}
export function SecretSvcDecryptValueResponseFromJSON(json) {
    return SecretSvcDecryptValueResponseFromJSONTyped(json, false);
}
export function SecretSvcDecryptValueResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'value': json['value'] == null ? undefined : json['value'],
        'values': json['values'] == null ? undefined : json['values'],
    };
}
export function SecretSvcDecryptValueResponseToJSON(json) {
    return SecretSvcDecryptValueResponseToJSONTyped(json, false);
}
export function SecretSvcDecryptValueResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'value': value['value'],
        'values': value['values'],
    };
}
