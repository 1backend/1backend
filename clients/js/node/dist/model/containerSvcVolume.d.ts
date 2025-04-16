/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export declare class ContainerSvcVolume {
    /**
    * Destination is the path inside the container.
    */
    'destination'?: string;
    /**
    * ReadOnly indicates whether the mount is read-only.
    */
    'readOnly'?: boolean;
    /**
    * Source is the host path or volume name.
    */
    'source'?: string;
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
