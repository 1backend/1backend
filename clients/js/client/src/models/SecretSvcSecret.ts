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

import { mapValues } from '../runtime';
import type { SecretSvcChecksumAlgorithm } from './SecretSvcChecksumAlgorithm';
import {
    SecretSvcChecksumAlgorithmFromJSON,
    SecretSvcChecksumAlgorithmFromJSONTyped,
    SecretSvcChecksumAlgorithmToJSON,
    SecretSvcChecksumAlgorithmToJSONTyped,
} from './SecretSvcChecksumAlgorithm';

/**
 * 
 * @export
 * @interface SecretSvcSecret
 */
export interface SecretSvcSecret {
    /**
     * Slugs of services/users who can change the deleters list
     * @type {Array<string>}
     * @memberof SecretSvcSecret
     */
    canChangeDeleters?: Array<string>;
    /**
     * Slugs of services/users who can change the readers list
     * @type {Array<string>}
     * @memberof SecretSvcSecret
     */
    canChangeReaders?: Array<string>;
    /**
     * Slugs of services/users who can change the writers list
     * @type {Array<string>}
     * @memberof SecretSvcSecret
     */
    canChangeWriters?: Array<string>;
    /**
     * Checksum of the secret value
     * @type {string}
     * @memberof SecretSvcSecret
     */
    checksum?: string;
    /**
     * Algorithm used for the checksum (e.g., "CRC32")
     * @type {SecretSvcChecksumAlgorithm}
     * @memberof SecretSvcSecret
     */
    checksumAlgorithm?: SecretSvcChecksumAlgorithm;
    /**
     * Slugs of services/users who can delete the secret
     * @type {Array<string>}
     * @memberof SecretSvcSecret
     */
    deleters?: Array<string>;
    /**
     * Whether the secret is encrypted
     * All secrets are encrypted before written to the DB.
     * This really only exists for write requests to know if the secret is already encrypted.
     * Ie: while most `secret save [key] [value]` commands are probably not encrypted,
     * File based saves, eg. `secret save secretA.yaml` are probably encrypted.
     * @type {boolean}
     * @memberof SecretSvcSecret
     */
    encrypted?: boolean;
    /**
     * Id of the secret
     * @type {string}
     * @memberof SecretSvcSecret
     */
    id?: string;
    /**
     * Envar or slug-like key of the secret
     * @type {string}
     * @memberof SecretSvcSecret
     */
    key?: string;
    /**
     * Namespace of the secret
     * @type {string}
     * @memberof SecretSvcSecret
     */
    namespace?: string;
    /**
     * Slugs of services/users who can read the secret
     * @type {Array<string>}
     * @memberof SecretSvcSecret
     */
    readers?: Array<string>;
    /**
     * Secret Value
     * @type {string}
     * @memberof SecretSvcSecret
     */
    value?: string;
    /**
     * Slugs of services/users who can modify the secret
     * @type {Array<string>}
     * @memberof SecretSvcSecret
     */
    writers?: Array<string>;
}



/**
 * Check if a given object implements the SecretSvcSecret interface.
 */
export function instanceOfSecretSvcSecret(value: object): value is SecretSvcSecret {
    return true;
}

export function SecretSvcSecretFromJSON(json: any): SecretSvcSecret {
    return SecretSvcSecretFromJSONTyped(json, false);
}

export function SecretSvcSecretFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcSecret {
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

export function SecretSvcSecretToJSON(json: any): SecretSvcSecret {
    return SecretSvcSecretToJSONTyped(json, false);
}

export function SecretSvcSecretToJSONTyped(value?: SecretSvcSecret | null, ignoreDiscriminator: boolean = false): any {
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

