/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ContainerSvcAsset
 */
export interface ContainerSvcAsset {
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcAsset
     */
    envVarKey: string;
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcAsset
     */
    url: string;
}

/**
 * Check if a given object implements the ContainerSvcAsset interface.
 */
export function instanceOfContainerSvcAsset(value: object): value is ContainerSvcAsset {
    if (!('envVarKey' in value) || value['envVarKey'] === undefined) return false;
    if (!('url' in value) || value['url'] === undefined) return false;
    return true;
}

export function ContainerSvcAssetFromJSON(json: any): ContainerSvcAsset {
    return ContainerSvcAssetFromJSONTyped(json, false);
}

export function ContainerSvcAssetFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcAsset {
    if (json == null) {
        return json;
    }
    return {
        
        'envVarKey': json['envVarKey'],
        'url': json['url'],
    };
}

export function ContainerSvcAssetToJSON(json: any): ContainerSvcAsset {
    return ContainerSvcAssetToJSONTyped(json, false);
}

export function ContainerSvcAssetToJSONTyped(value?: ContainerSvcAsset | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'envVarKey': value['envVarKey'],
        'url': value['url'],
    };
}

