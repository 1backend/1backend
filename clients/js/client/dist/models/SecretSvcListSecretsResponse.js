/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { SecretSvcSecretFromJSON, SecretSvcSecretToJSON, } from './SecretSvcSecret';
/**
 * Check if a given object implements the SecretSvcListSecretsResponse interface.
 */
export function instanceOfSecretSvcListSecretsResponse(value) {
    return true;
}
export function SecretSvcListSecretsResponseFromJSON(json) {
    return SecretSvcListSecretsResponseFromJSONTyped(json, false);
}
export function SecretSvcListSecretsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'secrets': json['secrets'] == null ? undefined : (json['secrets'].map(SecretSvcSecretFromJSON)),
    };
}
export function SecretSvcListSecretsResponseToJSON(json) {
    return SecretSvcListSecretsResponseToJSONTyped(json, false);
}
export function SecretSvcListSecretsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'secrets': value['secrets'] == null ? undefined : (value['secrets'].map(SecretSvcSecretToJSON)),
    };
}
