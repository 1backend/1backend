/* tslint:disable */
/* eslint-disable */
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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ContainerSvcImagePullableResponse
 */
export interface ContainerSvcImagePullableResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ContainerSvcImagePullableResponse
     */
    pullable: boolean;
}

/**
 * Check if a given object implements the ContainerSvcImagePullableResponse interface.
 */
export function instanceOfContainerSvcImagePullableResponse(value: object): value is ContainerSvcImagePullableResponse {
    if (!('pullable' in value) || value['pullable'] === undefined) return false;
    return true;
}

export function ContainerSvcImagePullableResponseFromJSON(json: any): ContainerSvcImagePullableResponse {
    return ContainerSvcImagePullableResponseFromJSONTyped(json, false);
}

export function ContainerSvcImagePullableResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcImagePullableResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'pullable': json['pullable'],
    };
}

export function ContainerSvcImagePullableResponseToJSON(json: any): ContainerSvcImagePullableResponse {
    return ContainerSvcImagePullableResponseToJSONTyped(json, false);
}

export function ContainerSvcImagePullableResponseToJSONTyped(value?: ContainerSvcImagePullableResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'pullable': value['pullable'],
    };
}

