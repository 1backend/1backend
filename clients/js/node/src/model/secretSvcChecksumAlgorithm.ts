/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export enum SecretSvcChecksumAlgorithm {
    ChecksumAlgorithmUnspecified = <any> '',
    ChecksumAlgorithmCRC32 = <any> 'CRC32',
    ChecksumAlgorithmBlake2s = <any> 'BLAKE2s',
    ChecksumAlgorithmSha256 = <any> 'SHA-256',
    ChecksumAlgorithmSha512 = <any> 'SHA-512'
}
