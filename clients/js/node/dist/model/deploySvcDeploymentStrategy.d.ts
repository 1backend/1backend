/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DeploySvcStrategyType } from './deploySvcStrategyType';
export declare class DeploySvcDeploymentStrategy {
    /**
    * Max extra replicas during update
    */
    'maxSurge'?: number;
    /**
    * Max unavailable replicas during update
    */
    'maxUnavailable'?: number;
    /**
    * Deployment strategy type (RollingUpdate, Recreate, etc.)
    */
    'type'?: DeploySvcStrategyType;
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
export declare namespace DeploySvcDeploymentStrategy {
}
