/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
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
 * @interface RegistrySvcRegisterInstanceRequest
 */
export interface RegistrySvcRegisterInstanceRequest {
    /**
     * The ID of the deployment that this instance is an instance of.
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    deploymentId?: string;
    /**
     * Host of the instance address. Required if URL is not provided
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    host?: string;
    /**
     * 
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    id?: string;
    /**
     * IP of the instance address. Optional: to register by IP instead of host
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    ip?: string;
    /**
     * Path of the instance address. Optional (e.g., "/api")
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    path?: string;
    /**
     * Port of the instance address. Required if URL is not provided
     * @type {number}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    port?: number;
    /**
     * Scheme of the instance address. Required if URL is not provided.
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    scheme?: string;
    /**
     * Full address URL of the instance.
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    url: string;
}

/**
 * Check if a given object implements the RegistrySvcRegisterInstanceRequest interface.
 */
export function instanceOfRegistrySvcRegisterInstanceRequest(value: object): value is RegistrySvcRegisterInstanceRequest {
    if (!('url' in value) || value['url'] === undefined) return false;
    return true;
}

export function RegistrySvcRegisterInstanceRequestFromJSON(json: any): RegistrySvcRegisterInstanceRequest {
    return RegistrySvcRegisterInstanceRequestFromJSONTyped(json, false);
}

export function RegistrySvcRegisterInstanceRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcRegisterInstanceRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'deploymentId': json['deploymentId'] == null ? undefined : json['deploymentId'],
        'host': json['host'] == null ? undefined : json['host'],
        'id': json['id'] == null ? undefined : json['id'],
        'ip': json['ip'] == null ? undefined : json['ip'],
        'path': json['path'] == null ? undefined : json['path'],
        'port': json['port'] == null ? undefined : json['port'],
        'scheme': json['scheme'] == null ? undefined : json['scheme'],
        'url': json['url'],
    };
}

export function RegistrySvcRegisterInstanceRequestToJSON(json: any): RegistrySvcRegisterInstanceRequest {
    return RegistrySvcRegisterInstanceRequestToJSONTyped(json, false);
}

export function RegistrySvcRegisterInstanceRequestToJSONTyped(value?: RegistrySvcRegisterInstanceRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'deploymentId': value['deploymentId'],
        'host': value['host'],
        'id': value['id'],
        'ip': value['ip'],
        'path': value['path'],
        'port': value['port'],
        'scheme': value['scheme'],
        'url': value['url'],
    };
}

