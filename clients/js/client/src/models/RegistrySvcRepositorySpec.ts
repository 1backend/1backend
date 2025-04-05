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
 * @interface RegistrySvcRepositorySpec
 */
export interface RegistrySvcRepositorySpec {
    /**
     * Context is the path to the image build context
     * @type {string}
     * @memberof RegistrySvcRepositorySpec
     */
    buildContext?: string;
    /**
     * ContainerFile is the path to the file that contains the container build instructions
     * Relative from the build context. By default, it is assumed to be a Dockerfile.
     * @type {string}
     * @memberof RegistrySvcRepositorySpec
     */
    containerFile?: string;
    /**
     * Ports the container will listen on internally
     * @type {Array<number>}
     * @memberof RegistrySvcRepositorySpec
     */
    internalPorts: Array<number>;
    /**
     * URL is the URL to the repository
     * @type {string}
     * @memberof RegistrySvcRepositorySpec
     */
    url: string;
    /**
     * Version of the code to use
     * @type {string}
     * @memberof RegistrySvcRepositorySpec
     */
    version?: string;
}

/**
 * Check if a given object implements the RegistrySvcRepositorySpec interface.
 */
export function instanceOfRegistrySvcRepositorySpec(value: object): value is RegistrySvcRepositorySpec {
    if (!('internalPorts' in value) || value['internalPorts'] === undefined) return false;
    if (!('url' in value) || value['url'] === undefined) return false;
    return true;
}

export function RegistrySvcRepositorySpecFromJSON(json: any): RegistrySvcRepositorySpec {
    return RegistrySvcRepositorySpecFromJSONTyped(json, false);
}

export function RegistrySvcRepositorySpecFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcRepositorySpec {
    if (json == null) {
        return json;
    }
    return {
        
        'buildContext': json['buildContext'] == null ? undefined : json['buildContext'],
        'containerFile': json['containerFile'] == null ? undefined : json['containerFile'],
        'internalPorts': json['internalPorts'],
        'url': json['url'],
        'version': json['version'] == null ? undefined : json['version'],
    };
}

export function RegistrySvcRepositorySpecToJSON(json: any): RegistrySvcRepositorySpec {
    return RegistrySvcRepositorySpecToJSONTyped(json, false);
}

export function RegistrySvcRepositorySpecToJSONTyped(value?: RegistrySvcRepositorySpec | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'buildContext': value['buildContext'],
        'containerFile': value['containerFile'],
        'internalPorts': value['internalPorts'],
        'url': value['url'],
        'version': value['version'],
    };
}

