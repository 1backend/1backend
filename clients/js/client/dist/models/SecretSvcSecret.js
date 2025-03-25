/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { SecretSvcChecksumAlgorithmFromJSON, SecretSvcChecksumAlgorithmToJSON, } from './SecretSvcChecksumAlgorithm';
/**
 * Check if a given object implements the SecretSvcSecret interface.
 */
export function instanceOfSecretSvcSecret(value) {
    return true;
}
export function SecretSvcSecretFromJSON(json) {
    return SecretSvcSecretFromJSONTyped(json, false);
}
export function SecretSvcSecretFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'canChangeDeleters': json['canChangeDeleters'] == null ? undefined : json['canChangeDeleters'],
        'canChangeReaders': json['canChangeReaders'] == null ? undefined : json['canChangeReaders'],
        'canChangeWriters': json['canChangeWriters'] == null ? undefined : json['canChangeWriters'],
        'checksum': json['checksum'] == null ? undefined : json['checksum'],
        'checksumAlgorithm': json['checksumAlgorithm'] == null ? undefined : SecretSvcChecksumAlgorithmFromJSON(json['checksumAlgorithm']),
        'deleters': json['deleters'] == null ? undefined : json['deleters'],
        'encrypted': json['encrypted'] == null ? undefined : json['encrypted'],
        'id': json['id'] == null ? undefined : json['id'],
        'key': json['key'] == null ? undefined : json['key'],
        'namespace': json['namespace'] == null ? undefined : json['namespace'],
        'readers': json['readers'] == null ? undefined : json['readers'],
        'value': json['value'] == null ? undefined : json['value'],
        'writers': json['writers'] == null ? undefined : json['writers'],
    };
}
export function SecretSvcSecretToJSON(json) {
    return SecretSvcSecretToJSONTyped(json, false);
}
export function SecretSvcSecretToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'canChangeDeleters': value['canChangeDeleters'],
        'canChangeReaders': value['canChangeReaders'],
        'canChangeWriters': value['canChangeWriters'],
        'checksum': value['checksum'],
        'checksumAlgorithm': SecretSvcChecksumAlgorithmToJSON(value['checksumAlgorithm']),
        'deleters': value['deleters'],
        'encrypted': value['encrypted'],
        'id': value['id'],
        'key': value['key'],
        'namespace': value['namespace'],
        'readers': value['readers'],
        'value': value['value'],
        'writers': value['writers'],
    };
}
