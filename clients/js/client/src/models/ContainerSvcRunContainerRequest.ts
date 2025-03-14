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
import type { ContainerSvcKeep } from './ContainerSvcKeep';
import {
    ContainerSvcKeepFromJSON,
    ContainerSvcKeepFromJSONTyped,
    ContainerSvcKeepToJSON,
    ContainerSvcKeepToJSONTyped,
} from './ContainerSvcKeep';
import type { ContainerSvcAsset } from './ContainerSvcAsset';
import {
    ContainerSvcAssetFromJSON,
    ContainerSvcAssetFromJSONTyped,
    ContainerSvcAssetToJSON,
    ContainerSvcAssetToJSONTyped,
} from './ContainerSvcAsset';
import type { ContainerSvcCapabilities } from './ContainerSvcCapabilities';
import {
    ContainerSvcCapabilitiesFromJSON,
    ContainerSvcCapabilitiesFromJSONTyped,
    ContainerSvcCapabilitiesToJSON,
    ContainerSvcCapabilitiesToJSONTyped,
} from './ContainerSvcCapabilities';
import type { ContainerSvcLabel } from './ContainerSvcLabel';
import {
    ContainerSvcLabelFromJSON,
    ContainerSvcLabelFromJSONTyped,
    ContainerSvcLabelToJSON,
    ContainerSvcLabelToJSONTyped,
} from './ContainerSvcLabel';
import type { ContainerSvcEnvVar } from './ContainerSvcEnvVar';
import {
    ContainerSvcEnvVarFromJSON,
    ContainerSvcEnvVarFromJSONTyped,
    ContainerSvcEnvVarToJSON,
    ContainerSvcEnvVarToJSONTyped,
} from './ContainerSvcEnvVar';
import type { ContainerSvcPortMapping } from './ContainerSvcPortMapping';
import {
    ContainerSvcPortMappingFromJSON,
    ContainerSvcPortMappingFromJSONTyped,
    ContainerSvcPortMappingToJSON,
    ContainerSvcPortMappingToJSONTyped,
} from './ContainerSvcPortMapping';

/**
 * 
 * @export
 * @interface ContainerSvcRunContainerRequest
 */
export interface ContainerSvcRunContainerRequest {
    /**
     * Assets maps environment variable names to file URLs.
     * Example: {"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"}
     * These files are downloaded by the File Svc and mounted in the container.
     * The environment variable `MODEL` will point to the local file path in the container.
     * @type {Array<ContainerSvcAsset>}
     * @memberof ContainerSvcRunContainerRequest
     */
    assets?: Array<ContainerSvcAsset>;
    /**
     * Capabilities define additional runtime features, such as GPU support.
     * @type {ContainerSvcCapabilities}
     * @memberof ContainerSvcRunContainerRequest
     */
    capabilities?: ContainerSvcCapabilities;
    /**
     * Envs are environment variables set within the container.
     * @type {Array<ContainerSvcEnvVar>}
     * @memberof ContainerSvcRunContainerRequest
     */
    envs?: Array<ContainerSvcEnvVar>;
    /**
     * Hash is a unique identifier for the container
     * @type {string}
     * @memberof ContainerSvcRunContainerRequest
     */
    hash?: string;
    /**
     * Image is the Docker image to use for the container
     * @type {string}
     * @memberof ContainerSvcRunContainerRequest
     */
    image: string;
    /**
     * Keeps are paths that persist across container restarts.
     * They function like mounts or volumes, but their external storage location is irrelevant.
     * @type {Array<ContainerSvcKeep>}
     * @memberof ContainerSvcRunContainerRequest
     */
    keeps?: Array<ContainerSvcKeep>;
    /**
     * Labels are metadata tags assigned to the container.
     * @type {Array<ContainerSvcLabel>}
     * @memberof ContainerSvcRunContainerRequest
     */
    labels?: Array<ContainerSvcLabel>;
    /**
     * Names are the human-readable aliases assigned to the container.
     * @type {Array<string>}
     * @memberof ContainerSvcRunContainerRequest
     */
    names?: Array<string>;
    /**
     * Ports maps host ports (keys) to container ports (values).
     * @type {Array<ContainerSvcPortMapping>}
     * @memberof ContainerSvcRunContainerRequest
     */
    ports?: Array<ContainerSvcPortMapping>;
}

/**
 * Check if a given object implements the ContainerSvcRunContainerRequest interface.
 */
export function instanceOfContainerSvcRunContainerRequest(value: object): value is ContainerSvcRunContainerRequest {
    if (!('image' in value) || value['image'] === undefined) return false;
    return true;
}

export function ContainerSvcRunContainerRequestFromJSON(json: any): ContainerSvcRunContainerRequest {
    return ContainerSvcRunContainerRequestFromJSONTyped(json, false);
}

export function ContainerSvcRunContainerRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcRunContainerRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'assets': json['assets'] == null ? undefined : ((json['assets'] as Array<any>).map(ContainerSvcAssetFromJSON)),
        'capabilities': json['capabilities'] == null ? undefined : ContainerSvcCapabilitiesFromJSON(json['capabilities']),
        'envs': json['envs'] == null ? undefined : ((json['envs'] as Array<any>).map(ContainerSvcEnvVarFromJSON)),
        'hash': json['hash'] == null ? undefined : json['hash'],
        'image': json['image'],
        'keeps': json['keeps'] == null ? undefined : ((json['keeps'] as Array<any>).map(ContainerSvcKeepFromJSON)),
        'labels': json['labels'] == null ? undefined : ((json['labels'] as Array<any>).map(ContainerSvcLabelFromJSON)),
        'names': json['names'] == null ? undefined : json['names'],
        'ports': json['ports'] == null ? undefined : ((json['ports'] as Array<any>).map(ContainerSvcPortMappingFromJSON)),
    };
}

export function ContainerSvcRunContainerRequestToJSON(json: any): ContainerSvcRunContainerRequest {
    return ContainerSvcRunContainerRequestToJSONTyped(json, false);
}

export function ContainerSvcRunContainerRequestToJSONTyped(value?: ContainerSvcRunContainerRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'assets': value['assets'] == null ? undefined : ((value['assets'] as Array<any>).map(ContainerSvcAssetToJSON)),
        'capabilities': ContainerSvcCapabilitiesToJSON(value['capabilities']),
        'envs': value['envs'] == null ? undefined : ((value['envs'] as Array<any>).map(ContainerSvcEnvVarToJSON)),
        'hash': value['hash'],
        'image': value['image'],
        'keeps': value['keeps'] == null ? undefined : ((value['keeps'] as Array<any>).map(ContainerSvcKeepToJSON)),
        'labels': value['labels'] == null ? undefined : ((value['labels'] as Array<any>).map(ContainerSvcLabelToJSON)),
        'names': value['names'],
        'ports': value['ports'] == null ? undefined : ((value['ports'] as Array<any>).map(ContainerSvcPortMappingToJSON)),
    };
}

