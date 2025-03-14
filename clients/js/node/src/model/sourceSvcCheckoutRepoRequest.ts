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

import { RequestFile } from './models';

export class SourceSvcCheckoutRepoRequest {
    /**
    * Password or token for HTTPS auth
    */
    'password'?: string;
    /**
    * SSH private key (optional for SSH connection)
    */
    'ssh_key'?: string;
    /**
    * Password for SSH private key if encrypted (optional)
    */
    'ssh_key_pwd'?: string;
    /**
    * Token for HTTPS auth (optional for SSH)
    */
    'token'?: string;
    /**
    * Full repository URL (e.g., https://github.com/user/repo)
    */
    'url'?: string;
    /**
    * Username for HTTPS or SSH user (optional for SSH)
    */
    'username'?: string;
    /**
    * Branch, tag, or commit SHA
    */
    'version'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "password",
            "baseName": "password",
            "type": "string"
        },
        {
            "name": "ssh_key",
            "baseName": "ssh_key",
            "type": "string"
        },
        {
            "name": "ssh_key_pwd",
            "baseName": "ssh_key_pwd",
            "type": "string"
        },
        {
            "name": "token",
            "baseName": "token",
            "type": "string"
        },
        {
            "name": "url",
            "baseName": "url",
            "type": "string"
        },
        {
            "name": "username",
            "baseName": "username",
            "type": "string"
        },
        {
            "name": "version",
            "baseName": "version",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return SourceSvcCheckoutRepoRequest.attributeTypeMap;
    }
}

