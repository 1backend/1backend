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

export class ContainerSvcBuildImageRequest {
    /**
    * ContextPath is the local path to the build context
    */
    'contextPath': string;
    /**
    * DockerfilePath is the local path to the Dockerfile
    */
    'dockerfilePath'?: string;
    /**
    * Name is the name of the image to build
    */
    'name': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "contextPath",
            "baseName": "contextPath",
            "type": "string"
        },
        {
            "name": "dockerfilePath",
            "baseName": "dockerfilePath",
            "type": "string"
        },
        {
            "name": "name",
            "baseName": "name",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcBuildImageRequest.attributeTypeMap;
    }
}

