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
 * Check if a given object implements the SecretSvcListSecretsRequest interface.
 */
export function instanceOfSecretSvcListSecretsRequest(value) {
    return true;
}
export function SecretSvcListSecretsRequestFromJSON(json) {
    return SecretSvcListSecretsRequestFromJSONTyped(json, false);
}
export function SecretSvcListSecretsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'key': json['key'] == null ? undefined : json['key'],
        'keys': json['keys'] == null ? undefined : json['keys'],
        'namespace': json['namespace'] == null ? undefined : json['namespace'],
    };
}
export function SecretSvcListSecretsRequestToJSON(json) {
    return SecretSvcListSecretsRequestToJSONTyped(json, false);
}
export function SecretSvcListSecretsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'key': value['key'],
        'keys': value['keys'],
        'namespace': value['namespace'],
    };
}
