/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 */
export const SecretSvcChecksumAlgorithm = {
    ChecksumAlgorithmUnspecified: '',
    ChecksumAlgorithmCRC32: 'CRC32',
    ChecksumAlgorithmBlake2s: 'BLAKE2s',
    ChecksumAlgorithmSha256: 'SHA-256',
    ChecksumAlgorithmSha512: 'SHA-512'
};
export function instanceOfSecretSvcChecksumAlgorithm(value) {
    for (const key in SecretSvcChecksumAlgorithm) {
        if (Object.prototype.hasOwnProperty.call(SecretSvcChecksumAlgorithm, key)) {
            if (SecretSvcChecksumAlgorithm[key] === value) {
                return true;
            }
        }
    }
    return false;
}
export function SecretSvcChecksumAlgorithmFromJSON(json) {
    return SecretSvcChecksumAlgorithmFromJSONTyped(json, false);
}
export function SecretSvcChecksumAlgorithmFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
export function SecretSvcChecksumAlgorithmToJSON(value) {
    return value;
}
export function SecretSvcChecksumAlgorithmToJSONTyped(value, ignoreDiscriminator) {
    return value;
}
