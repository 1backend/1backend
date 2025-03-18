/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
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
 * @interface ContainerSvcVolume
 */
export interface ContainerSvcVolume {
    /**
     * Destination is the path inside the container.
     * @type {string}
     * @memberof ContainerSvcVolume
     */
    destination?: string;
    /**
     * ReadOnly indicates whether the mount is read-only.
     * @type {boolean}
     * @memberof ContainerSvcVolume
     */
    readOnly?: boolean;
    /**
     * Source is the host path or volume name.
     * @type {string}
     * @memberof ContainerSvcVolume
     */
    source?: string;
}

/**
 * Check if a given object implements the ContainerSvcVolume interface.
 */
export function instanceOfContainerSvcVolume(value: object): value is ContainerSvcVolume {
    return true;
}

export function ContainerSvcVolumeFromJSON(json: any): ContainerSvcVolume {
    return ContainerSvcVolumeFromJSONTyped(json, false);
}

export function ContainerSvcVolumeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcVolume {
    if (json == null) {
        return json;
    }
    return {
        
        'destination': json['destination'] == null ? undefined : json['destination'],
        'readOnly': json['readOnly'] == null ? undefined : json['readOnly'],
        'source': json['source'] == null ? undefined : json['source'],
    };
}

export function ContainerSvcVolumeToJSON(json: any): ContainerSvcVolume {
    return ContainerSvcVolumeToJSONTyped(json, false);
}

export function ContainerSvcVolumeToJSONTyped(value?: ContainerSvcVolume | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'destination': value['destination'],
        'readOnly': value['readOnly'],
        'source': value['source'],
    };
}

