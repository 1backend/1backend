/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { ModelSvcContainer } from './modelSvcContainer';
export declare class ModelSvcCudaParameters {
    /**
    * Container configuration related to CUDA usage.
    */
    'container'?: ModelSvcContainer;
    /**
    * Level of precision for selecting the CUDA version when resolving the container image. - 2 -> Use \"major.minor\" (e.g., \"12.2\") - 3 -> Use \"major.minor.patch\" (e.g., \"12.2.0\")
    */
    'cudaVersionPrecision'?: number;
    /**
    * Default CUDA version to use (e.g., \"12.2\" or \"12.2.0\").
    */
    'defaultCudaVersion'?: string;
    /**
    * Default cuDNN version to use alongside CUDA.
    */
    'defaultCudnnVersion'?: string;
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
