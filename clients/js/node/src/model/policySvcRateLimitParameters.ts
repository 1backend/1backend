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
import { PolicySvcEntity } from './policySvcEntity';
import { PolicySvcScope } from './policySvcScope';

export class PolicySvcRateLimitParameters {
    'entity'?: PolicySvcEntity;
    'maxRequests'?: number;
    'scope'?: PolicySvcScope;
    'timeWindow'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "entity",
            "baseName": "entity",
            "type": "PolicySvcEntity"
        },
        {
            "name": "maxRequests",
            "baseName": "maxRequests",
            "type": "number"
        },
        {
            "name": "scope",
            "baseName": "scope",
            "type": "PolicySvcScope"
        },
        {
            "name": "timeWindow",
            "baseName": "timeWindow",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return PolicySvcRateLimitParameters.attributeTypeMap;
    }
}

export namespace PolicySvcRateLimitParameters {
}
