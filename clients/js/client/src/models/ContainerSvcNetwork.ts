/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
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
 * @interface ContainerSvcNetwork
 */
export interface ContainerSvcNetwork {
    /**
     * IPAddress is the assigned IP address of the container.
     * @type {string}
     * @memberof ContainerSvcNetwork
     */
    ipAddress?: string;
    /**
     * MacAddress is the container's MAC address if applicable.
     * @type {string}
     * @memberof ContainerSvcNetwork
     */
    macAddress?: string;
    /**
     * Mode specifies the container's network mode (e.g., bridge, host, none, custom).
     * @type {string}
     * @memberof ContainerSvcNetwork
     */
    mode?: string;
}

/**
 * Check if a given object implements the ContainerSvcNetwork interface.
 */
export function instanceOfContainerSvcNetwork(value: object): value is ContainerSvcNetwork {
    return true;
}

export function ContainerSvcNetworkFromJSON(json: any): ContainerSvcNetwork {
    return ContainerSvcNetworkFromJSONTyped(json, false);
}

export function ContainerSvcNetworkFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcNetwork {
    if (json == null) {
        return json;
    }
    return {
        
        'ipAddress': json['ipAddress'] == null ? undefined : json['ipAddress'],
        'macAddress': json['macAddress'] == null ? undefined : json['macAddress'],
        'mode': json['mode'] == null ? undefined : json['mode'],
    };
}

export function ContainerSvcNetworkToJSON(json: any): ContainerSvcNetwork {
    return ContainerSvcNetworkToJSONTyped(json, false);
}

export function ContainerSvcNetworkToJSONTyped(value?: ContainerSvcNetwork | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ipAddress': value['ipAddress'],
        'macAddress': value['macAddress'],
        'mode': value['mode'],
    };
}

