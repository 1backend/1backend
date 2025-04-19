/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { ModelSvcArchitectures } from './modelSvcArchitectures';
import { PromptSvcPromptType } from './promptSvcPromptType';

export class ModelSvcPlatform {
    'architectures'?: ModelSvcArchitectures;
    'id'?: string;
    'name'?: string;
    /**
    * List of prompt types that the AI engine supports.
    */
    'types'?: Array<PromptSvcPromptType>;
    'version'?: number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "architectures",
            "baseName": "architectures",
            "type": "ModelSvcArchitectures"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "name",
            "baseName": "name",
            "type": "string"
        },
        {
            "name": "types",
            "baseName": "types",
            "type": "Array<PromptSvcPromptType>"
        },
        {
            "name": "version",
            "baseName": "version",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return ModelSvcPlatform.attributeTypeMap;
    }
}

