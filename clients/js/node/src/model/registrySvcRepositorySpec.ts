/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class RegistrySvcRepositorySpec {
    /**
    * Context is the path to the image build context
    */
    'buildContext'?: string;
    /**
    * ContainerFile is the path to the file that contains the container build instructions Relative from the build context. By default, it is assumed to be a Dockerfile.
    */
    'containerFile'?: string;
    /**
    * Ports the container will listen on internally
    */
    'internalPorts': Array<number>;
    /**
    * URL is the URL to the repository
    */
    'url': string;
    /**
    * Version of the code to use
    */
    'version'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "buildContext",
            "baseName": "buildContext",
            "type": "string"
        },
        {
            "name": "containerFile",
            "baseName": "containerFile",
            "type": "string"
        },
        {
            "name": "internalPorts",
            "baseName": "internalPorts",
            "type": "Array<number>"
        },
        {
            "name": "url",
            "baseName": "url",
            "type": "string"
        },
        {
            "name": "version",
            "baseName": "version",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcRepositorySpec.attributeTypeMap;
    }
}

