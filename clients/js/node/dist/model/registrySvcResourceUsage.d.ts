/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { RegistrySvcUsage } from './registrySvcUsage';
export declare class RegistrySvcResourceUsage {
    /**
    * CPU usage metrics.
    */
    'cpu'?: RegistrySvcUsage;
    /**
    * Disk usage metrics.
    */
    'disk'?: RegistrySvcUsage;
    /**
    * Memory usage metrics.
    */
    'memory'?: RegistrySvcUsage;
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
