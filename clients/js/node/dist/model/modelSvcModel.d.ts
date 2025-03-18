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
import { ModelSvcAsset } from './modelSvcAsset';
export declare class ModelSvcModel {
    'assets'?: Array<ModelSvcAsset>;
    'bits'?: number;
    'description'?: string;
    'extension'?: string;
    'flavour'?: string;
    'fullName'?: string;
    'id': string;
    'maxBits'?: number;
    'maxRam'?: number;
    'mirrors'?: Array<string>;
    'name': string;
    'parameters'?: string;
    'platformId': string;
    'promptTemplate'?: string;
    'quality'?: string;
    'quantComment'?: string;
    'size'?: number;
    'tags'?: Array<string>;
    'uncensored'?: boolean;
    'version'?: string;
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
