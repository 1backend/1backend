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
/**
 * Check if a given object implements the SecretSvcRemoveSecretsRequest interface.
 */
export function instanceOfSecretSvcRemoveSecretsRequest(value) {
    return true;
}
export function SecretSvcRemoveSecretsRequestFromJSON(json) {
    return SecretSvcRemoveSecretsRequestFromJSONTyped(json, false);
}
export function SecretSvcRemoveSecretsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'id': json['id'] == null ? undefined : json['id'],
        'ids': json['ids'] == null ? undefined : json['ids'],
        'key': json['key'] == null ? undefined : json['key'],
        'keys': json['keys'] == null ? undefined : json['keys'],
    };
}
export function SecretSvcRemoveSecretsRequestToJSON(json) {
    return SecretSvcRemoveSecretsRequestToJSONTyped(json, false);
}
export function SecretSvcRemoveSecretsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'id': value['id'],
        'ids': value['ids'],
        'key': value['key'],
        'keys': value['keys'],
    };
}
