/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { ModelSvcEnvVar } from './modelSvcEnvVar';
import { ModelSvcKeep } from './modelSvcKeep';

export class ModelSvcContainer {
    /**
    * Environment variables to be passed to the container (e.g., \"DEVICES=all\").
    */
    'envars'?: Array<ModelSvcEnvVar>;
    /**
    * Template for constructing the container image name.
    */
    'imageTemplate'?: string;
    /**
    * List of container paths that should persist across restarts.
    */
    'keeps'?: Array<ModelSvcKeep>;
    /**
    * Internal port exposed by the container.
    */
    'port'?: number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "envars",
            "baseName": "envars",
            "type": "Array<ModelSvcEnvVar>"
        },
        {
            "name": "imageTemplate",
            "baseName": "imageTemplate",
            "type": "string"
        },
        {
            "name": "keeps",
            "baseName": "keeps",
            "type": "Array<ModelSvcKeep>"
        },
        {
            "name": "port",
            "baseName": "port",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return ModelSvcContainer.attributeTypeMap;
    }
}

