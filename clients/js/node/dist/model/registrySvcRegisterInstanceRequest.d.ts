/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export declare class RegistrySvcRegisterInstanceRequest {
    /**
    * The ID of the deployment that this instance is an instance of.
    */
    'deploymentId'?: string;
    /**
    * Host of the instance address. Required if URL is not provided
    */
    'host'?: string;
    'id'?: string;
    /**
    * IP of the instance address. Optional: to register by IP instead of host
    */
    'ip'?: string;
    /**
    * Path of the instance address. Optional (e.g., \"/api\")
    */
    'path'?: string;
    /**
    * Port of the instance address. Required if URL is not provided
    */
    'port'?: number;
    /**
    * Scheme of the instance address. Required if URL is not provided.
    */
    'scheme'?: string;
    /**
    * Full address URL of the instance.
    */
    'url': string;
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
