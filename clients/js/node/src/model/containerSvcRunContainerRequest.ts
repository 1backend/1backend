/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { ContainerSvcAsset } from './containerSvcAsset';
import { ContainerSvcCapabilities } from './containerSvcCapabilities';
import { ContainerSvcEnvVar } from './containerSvcEnvVar';
import { ContainerSvcKeep } from './containerSvcKeep';
import { ContainerSvcLabel } from './containerSvcLabel';
import { ContainerSvcPortMapping } from './containerSvcPortMapping';

export class ContainerSvcRunContainerRequest {
    /**
    * Assets maps environment variable names to file URLs. Example: {\"MODEL\": \"https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf\"} These files are downloaded by the File Svc and mounted in the container. The environment variable `MODEL` will point to the local file path in the container.
    */
    'assets'?: Array<ContainerSvcAsset>;
    /**
    * Capabilities define additional runtime features, such as GPU support.
    */
    'capabilities'?: ContainerSvcCapabilities;
    /**
    * Envs are environment variables set within the container.
    */
    'envs'?: Array<ContainerSvcEnvVar>;
    /**
    * Hash is a unique identifier for the container
    */
    'hash'?: string;
    /**
    * Image is the Docker image to use for the container
    */
    'image': string;
    /**
    * Keeps are paths that persist across container restarts. They function like mounts or volumes, but their external storage location is irrelevant.
    */
    'keeps'?: Array<ContainerSvcKeep>;
    /**
    * Labels are metadata tags assigned to the container.
    */
    'labels'?: Array<ContainerSvcLabel>;
    /**
    * Names are the human-readable aliases assigned to the container.
    */
    'names'?: Array<string>;
    /**
    * Ports maps host ports (keys) to container ports (values).
    */
    'ports'?: Array<ContainerSvcPortMapping>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "assets",
            "baseName": "assets",
            "type": "Array<ContainerSvcAsset>"
        },
        {
            "name": "capabilities",
            "baseName": "capabilities",
            "type": "ContainerSvcCapabilities"
        },
        {
            "name": "envs",
            "baseName": "envs",
            "type": "Array<ContainerSvcEnvVar>"
        },
        {
            "name": "hash",
            "baseName": "hash",
            "type": "string"
        },
        {
            "name": "image",
            "baseName": "image",
            "type": "string"
        },
        {
            "name": "keeps",
            "baseName": "keeps",
            "type": "Array<ContainerSvcKeep>"
        },
        {
            "name": "labels",
            "baseName": "labels",
            "type": "Array<ContainerSvcLabel>"
        },
        {
            "name": "names",
            "baseName": "names",
            "type": "Array<string>"
        },
        {
            "name": "ports",
            "baseName": "ports",
            "type": "Array<ContainerSvcPortMapping>"
        }    ];

    static getAttributeTypeMap() {
        return ContainerSvcRunContainerRequest.attributeTypeMap;
    }
}

