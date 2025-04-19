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
export declare class DeploySvcResourceLimits {
    /**
    * CPU limit, e.g., \"500m\" for 0.5 cores
    */
    'cpu'?: string;
    /**
    * Memory limit, e.g., \"128Mi\"
    */
    'memory'?: string;
    /**
    * Optional: GPU VRAM requirement, e.g., \"48GB\"
    */
    'vram'?: string;
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
