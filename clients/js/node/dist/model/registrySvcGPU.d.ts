/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { RegistrySvcProcess } from './registrySvcProcess';
export declare class RegistrySvcGPU {
    'busId'?: string;
    'computeMode'?: string;
    'cudaVersion'?: string;
    'gpuUtilization'?: number;
    /**
    * Id Node.URL + IntraNodeId
    */
    'id'?: string;
    'intraNodeId'?: number;
    'memoryTotal'?: number;
    'memoryUsage'?: number;
    'name'?: string;
    'performanceState'?: string;
    'powerCap'?: number;
    'powerUsage'?: number;
    'processDetails'?: Array<RegistrySvcProcess>;
    'temperature'?: number;
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
