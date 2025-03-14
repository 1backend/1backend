/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { ModelSvcEnvVar } from './modelSvcEnvVar';
import { ModelSvcKeep } from './modelSvcKeep';
export declare class ModelSvcContainer {
    /**
    * Environment variables to be passed to the container (e.g., \"DEVICES=all\").
    */
    'envars'?: Array<ModelSvcEnvVar>;
    /**
    * Template for constructing the container image name.
    */
    'imageTemplate'?: string;
    /**
    * List of container paths that should persist across restarts.
    */
    'keeps'?: Array<ModelSvcKeep>;
    /**
    * Internal port exposed by the container.
    */
    'port'?: number;
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
