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

import { RequestFile } from './models';
import { SecretSvcChecksumAlgorithm } from './secretSvcChecksumAlgorithm';

export class SecretSvcSecret {
    /**
    * Slugs of services/users who can change the deleters list
    */
    'canChangeDeleters'?: Array<string>;
    /**
    * Slugs of services/users who can change the readers list
    */
    'canChangeReaders'?: Array<string>;
    /**
    * Slugs of services/users who can change the writers list
    */
    'canChangeWriters'?: Array<string>;
    /**
    * Checksum of the secret value
    */
    'checksum'?: string;
    /**
    * Algorithm used for the checksum (e.g., \"CRC32\")
    */
    'checksumAlgorithm'?: SecretSvcChecksumAlgorithm;
    /**
    * Slugs of services/users who can delete the secret
    */
    'deleters'?: Array<string>;
    /**
    * Whether the secret is encrypted All secrets are encrypted before written to the DB. This really only exists for write requests to know if the secret is already encrypted. Ie: while most `secret save [key] [value]` commands are probably not encrypted, File based saves, eg. `secret save secretA.yaml` are probably encrypted.
    */
    'encrypted'?: boolean;
    /**
    * Id of the secret
    */
    'id'?: string;
    /**
    * Envar or slug-like key of the secret
    */
    'key'?: string;
    /**
    * Namespace of the secret
    */
    'namespace'?: string;
    /**
    * Slugs of services/users who can read the secret
    */
    'readers'?: Array<string>;
    /**
    * Secret Value
    */
    'value'?: string;
    /**
    * Slugs of services/users who can modify the secret
    */
    'writers'?: Array<string>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "canChangeDeleters",
            "baseName": "canChangeDeleters",
            "type": "Array<string>"
        },
        {
            "name": "canChangeReaders",
            "baseName": "canChangeReaders",
            "type": "Array<string>"
        },
        {
            "name": "canChangeWriters",
            "baseName": "canChangeWriters",
            "type": "Array<string>"
        },
        {
            "name": "checksum",
            "baseName": "checksum",
            "type": "string"
        },
        {
            "name": "checksumAlgorithm",
            "baseName": "checksumAlgorithm",
            "type": "SecretSvcChecksumAlgorithm"
        },
        {
            "name": "deleters",
            "baseName": "deleters",
            "type": "Array<string>"
        },
        {
            "name": "encrypted",
            "baseName": "encrypted",
            "type": "boolean"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "key",
            "baseName": "key",
            "type": "string"
        },
        {
            "name": "namespace",
            "baseName": "namespace",
            "type": "string"
        },
        {
            "name": "readers",
            "baseName": "readers",
            "type": "Array<string>"
        },
        {
            "name": "value",
            "baseName": "value",
            "type": "string"
        },
        {
            "name": "writers",
            "baseName": "writers",
            "type": "Array<string>"
        }    ];

    static getAttributeTypeMap() {
        return SecretSvcSecret.attributeTypeMap;
    }
}

export namespace SecretSvcSecret {
}
