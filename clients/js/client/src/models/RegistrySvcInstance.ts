/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { RegistrySvcInstanceStatus } from './RegistrySvcInstanceStatus';
import {
    RegistrySvcInstanceStatusFromJSON,
    RegistrySvcInstanceStatusFromJSONTyped,
    RegistrySvcInstanceStatusToJSON,
    RegistrySvcInstanceStatusToJSONTyped,
} from './RegistrySvcInstanceStatus';

/**
 * 
 * @export
 * @interface RegistrySvcInstance
 */
export interface RegistrySvcInstance {
    /**
     * The ID of the deployment that this instance is an instance of.
     * Only instances deployed by 1Backend have a DeploymentId.
     * Services can be deployed through other means (Docker Compose, K8s, anything),
     * in that case they self-register and will not have a DeploymentId.
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    deploymentId?: string;
    /**
     * Details
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    details?: string;
    /**
     * Host of the instance address. Required if URL is not provided
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    host?: string;
    /**
     * Required: ID of the instance
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    id: string;
    /**
     * IP of the instance address. Optional: to register by IP instead of host
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    ip?: string;
    /**
     * Last time the instance gave a sign of life
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    lastHeartbeat?: string;
    /**
     * NodeURL is the URL of the 1Backend server the instance is running on.
     * To have a NodeURL the instance must either:
     * - Be deployed by 1Backend
     * - Declare the 1Backend server URL when registering its instance
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    nodeUrl?: string;
    /**
     * Path of the instance address. Optional (e.g., "/api")
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    path?: string;
    /**
     * Port of the instance address. Required if URL is not provided
     * @type {number}
     * @memberof RegistrySvcInstance
     */
    port?: number;
    /**
     * Scheme of the instance address. Required if URL is not provided.
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    scheme?: string;
    /**
     * Slug of the account that owns this instance
     * Services that want to be proxied by their slug are advised to self register
     * their instance at startup.
     * Keep in mind, instances might be deployed by 1Backend yet they still won't be 1Backend services
     * and they won't have slugs. Think NGINX, MySQL, etc.
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    slug?: string;
    /**
     * Status
     * @type {RegistrySvcInstanceStatus}
     * @memberof RegistrySvcInstance
     */
    status: RegistrySvcInstanceStatus;
    /**
     * Tags are used to filter instances
     * @type {Array<string>}
     * @memberof RegistrySvcInstance
     */
    tags?: Array<string>;
    /**
     * Full address URL of the instance.
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    url: string;
}



/**
 * Check if a given object implements the RegistrySvcInstance interface.
 */
export function instanceOfRegistrySvcInstance(value: object): value is RegistrySvcInstance {
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('status' in value) || value['status'] === undefined) return false;
    if (!('url' in value) || value['url'] === undefined) return false;
    return true;
}

export function RegistrySvcInstanceFromJSON(json: any): RegistrySvcInstance {
    return RegistrySvcInstanceFromJSONTyped(json, false);
}

export function RegistrySvcInstanceFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcInstance {
    if (json == null) {
        return json;
    }
    return {
        
        'deploymentId': json['deploymentId'] == null ? undefined : json['deploymentId'],
        'details': json['details'] == null ? undefined : json['details'],
        'host': json['host'] == null ? undefined : json['host'],
        'id': json['id'],
        'ip': json['ip'] == null ? undefined : json['ip'],
        'lastHeartbeat': json['lastHeartbeat'] == null ? undefined : json['lastHeartbeat'],
        'nodeUrl': json['nodeUrl'] == null ? undefined : json['nodeUrl'],
        'path': json['path'] == null ? undefined : json['path'],
        'port': json['port'] == null ? undefined : json['port'],
        'scheme': json['scheme'] == null ? undefined : json['scheme'],
        'slug': json['slug'] == null ? undefined : json['slug'],
        'status': RegistrySvcInstanceStatusFromJSON(json['status']),
        'tags': json['tags'] == null ? undefined : json['tags'],
        'url': json['url'],
    };
}

export function RegistrySvcInstanceToJSON(json: any): RegistrySvcInstance {
    return RegistrySvcInstanceToJSONTyped(json, false);
}

export function RegistrySvcInstanceToJSONTyped(value?: RegistrySvcInstance | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'deploymentId': value['deploymentId'],
        'details': value['details'],
        'host': value['host'],
        'id': value['id'],
        'ip': value['ip'],
        'lastHeartbeat': value['lastHeartbeat'],
        'nodeUrl': value['nodeUrl'],
        'path': value['path'],
        'port': value['port'],
        'scheme': value['scheme'],
        'slug': value['slug'],
        'status': RegistrySvcInstanceStatusToJSON(value['status']),
        'tags': value['tags'],
        'url': value['url'],
    };
}

