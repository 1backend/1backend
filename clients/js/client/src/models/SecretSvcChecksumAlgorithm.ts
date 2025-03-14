/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
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
 * 
 * @export
 */
export const SecretSvcChecksumAlgorithm = {
    ChecksumAlgorithmUnspecified: '',
    ChecksumAlgorithmCRC32: 'CRC32',
    ChecksumAlgorithmBlake2s: 'BLAKE2s',
    ChecksumAlgorithmSha256: 'SHA-256',
    ChecksumAlgorithmSha512: 'SHA-512'
} as const;
export type SecretSvcChecksumAlgorithm = typeof SecretSvcChecksumAlgorithm[keyof typeof SecretSvcChecksumAlgorithm];


export function instanceOfSecretSvcChecksumAlgorithm(value: any): boolean {
    for (const key in SecretSvcChecksumAlgorithm) {
        if (Object.prototype.hasOwnProperty.call(SecretSvcChecksumAlgorithm, key)) {
            if (SecretSvcChecksumAlgorithm[key as keyof typeof SecretSvcChecksumAlgorithm] === value) {
                return true;
            }
        }
    }
    return false;
}

export function SecretSvcChecksumAlgorithmFromJSON(json: any): SecretSvcChecksumAlgorithm {
    return SecretSvcChecksumAlgorithmFromJSONTyped(json, false);
}

export function SecretSvcChecksumAlgorithmFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcChecksumAlgorithm {
    return json as SecretSvcChecksumAlgorithm;
}

export function SecretSvcChecksumAlgorithmToJSON(value?: SecretSvcChecksumAlgorithm | null): any {
    return value as any;
}

export function SecretSvcChecksumAlgorithmToJSONTyped(value: any, ignoreDiscriminator: boolean): SecretSvcChecksumAlgorithm {
    return value as SecretSvcChecksumAlgorithm;
}

