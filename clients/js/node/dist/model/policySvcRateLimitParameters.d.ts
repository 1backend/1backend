/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { PolicySvcEntity } from './policySvcEntity';
import { PolicySvcScope } from './policySvcScope';
export declare class PolicySvcRateLimitParameters {
    'entity'?: PolicySvcEntity;
    'maxRequests'?: number;
    'scope'?: PolicySvcScope;
    'timeWindow'?: string;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
export declare namespace PolicySvcRateLimitParameters {
}
