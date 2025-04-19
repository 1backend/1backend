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
import { PolicySvcParameters } from './policySvcParameters';
import { PolicySvcTemplateId } from './policySvcTemplateId';

export class PolicySvcInstance {
    'endpoint'?: string;
    'id'?: string;
    'parameters': PolicySvcParameters;
    'templateId': PolicySvcTemplateId;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "endpoint",
            "baseName": "endpoint",
            "type": "string"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "parameters",
            "baseName": "parameters",
            "type": "PolicySvcParameters"
        },
        {
            "name": "templateId",
            "baseName": "templateId",
            "type": "PolicySvcTemplateId"
        }    ];

    static getAttributeTypeMap() {
        return PolicySvcInstance.attributeTypeMap;
    }
}

export namespace PolicySvcInstance {
}
