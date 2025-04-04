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

import { RequestFile } from './models';

export class DeploySvcAutoScalingConfig {
    /**
    * CPU usage threshold for scaling (as a percentage)
    */
    'cpuThreshold'?: number;
    /**
    * Maximum number of replicas to run
    */
    'maxReplicas'?: number;
    /**
    * Minimum number of replicas to run
    */
    'minReplicas'?: number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "cpuThreshold",
            "baseName": "cpuThreshold",
            "type": "number"
        },
        {
            "name": "maxReplicas",
            "baseName": "maxReplicas",
            "type": "number"
        },
        {
            "name": "minReplicas",
            "baseName": "minReplicas",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return DeploySvcAutoScalingConfig.attributeTypeMap;
    }
}

