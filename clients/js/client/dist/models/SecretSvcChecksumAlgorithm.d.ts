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
 *
 * @export
 */
export declare const SecretSvcChecksumAlgorithm: {
    readonly ChecksumAlgorithmUnspecified: "";
    readonly ChecksumAlgorithmCRC32: "CRC32";
    readonly ChecksumAlgorithmBlake2s: "BLAKE2s";
    readonly ChecksumAlgorithmSha256: "SHA-256";
    readonly ChecksumAlgorithmSha512: "SHA-512";
};
export type SecretSvcChecksumAlgorithm = typeof SecretSvcChecksumAlgorithm[keyof typeof SecretSvcChecksumAlgorithm];
export declare function instanceOfSecretSvcChecksumAlgorithm(value: any): boolean;
export declare function SecretSvcChecksumAlgorithmFromJSON(json: any): SecretSvcChecksumAlgorithm;
export declare function SecretSvcChecksumAlgorithmFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcChecksumAlgorithm;
export declare function SecretSvcChecksumAlgorithmToJSON(value?: SecretSvcChecksumAlgorithm | null): any;
export declare function SecretSvcChecksumAlgorithmToJSONTyped(value: any, ignoreDiscriminator: boolean): SecretSvcChecksumAlgorithm;
