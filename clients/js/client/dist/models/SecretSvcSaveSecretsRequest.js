/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { SecretSvcSecretFromJSON, SecretSvcSecretToJSON, } from './SecretSvcSecret';
/**
 * Check if a given object implements the SecretSvcSaveSecretsRequest interface.
 */
export function instanceOfSecretSvcSaveSecretsRequest(value) {
    return true;
}
export function SecretSvcSaveSecretsRequestFromJSON(json) {
    return SecretSvcSaveSecretsRequestFromJSONTyped(json, false);
}
export function SecretSvcSaveSecretsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'secrets': json['secrets'] == null ? undefined : (json['secrets'].map(SecretSvcSecretFromJSON)),
    };
}
export function SecretSvcSaveSecretsRequestToJSON(json) {
    return SecretSvcSaveSecretsRequestToJSONTyped(json, false);
}
export function SecretSvcSaveSecretsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'secrets': value['secrets'] == null ? undefined : (value['secrets'].map(SecretSvcSecretToJSON)),
    };
}
