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
import { ModelSvcCudaParameters } from './modelSvcCudaParameters';
import { ModelSvcDefaultParameters } from './modelSvcDefaultParameters';
export declare class ModelSvcArchitectures {
    /**
    * CUDA-specific container parameters, if applicable.
    */
    'cuda'?: ModelSvcCudaParameters;
    /**
    * Default container configuration for non-GPU environments.
    */
    '_default'?: ModelSvcDefaultParameters;
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
