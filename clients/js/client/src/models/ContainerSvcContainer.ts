/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
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
import type { ContainerSvcNetwork } from './ContainerSvcNetwork';
import {
    ContainerSvcNetworkFromJSON,
    ContainerSvcNetworkFromJSONTyped,
    ContainerSvcNetworkToJSON,
    ContainerSvcNetworkToJSONTyped,
} from './ContainerSvcNetwork';
import type { ContainerSvcCapabilities } from './ContainerSvcCapabilities';
import {
    ContainerSvcCapabilitiesFromJSON,
    ContainerSvcCapabilitiesFromJSONTyped,
    ContainerSvcCapabilitiesToJSON,
    ContainerSvcCapabilitiesToJSONTyped,
} from './ContainerSvcCapabilities';
import type { ContainerSvcVolume } from './ContainerSvcVolume';
import {
    ContainerSvcVolumeFromJSON,
    ContainerSvcVolumeFromJSONTyped,
    ContainerSvcVolumeToJSON,
    ContainerSvcVolumeToJSONTyped,
} from './ContainerSvcVolume';
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
import type { ContainerSvcResources } from './ContainerSvcResources';
import {
    ContainerSvcResourcesFromJSON,
    ContainerSvcResourcesFromJSONTyped,
    ContainerSvcResourcesToJSON,
    ContainerSvcResourcesToJSONTyped,
} from './ContainerSvcResources';
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
 * @interface ContainerSvcContainer
 */
export interface ContainerSvcContainer {
    /**
     * Assets maps environment variable names to file URLs.
     * Example: {"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"}
     * These files are downloaded by the File Svc and mounted in the container.
     * The environment variable `MODEL` will point to the local file path in the container.
     * @type {Array<ContainerSvcAsset>}
     * @memberof ContainerSvcContainer
     */
    assets?: Array<ContainerSvcAsset>;
    /**
     * Capabilities define additional runtime features, such as GPU support.
     * @type {ContainerSvcCapabilities}
     * @memberof ContainerSvcContainer
     */
    capabilities?: ContainerSvcCapabilities;
    /**
     * Envs are environment variables set within the container.
     * @type {Array<ContainerSvcEnvVar>}
     * @memberof ContainerSvcContainer
     */
    envs?: Array<ContainerSvcEnvVar>;
    /**
     * Hash is a unique identifier associated with the container.
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    hash?: string;
    /**
     * Id is the unique identifier for the container instance.
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    id?: string;
    /**
     * Image is the Docker image used to create the container.
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    image?: string;
    /**
     * Keeps are paths that persist across container restarts.
     * They function like mounts or volumes, but their external storage location is irrelevant.
     * @type {Array<ContainerSvcKeep>}
     * @memberof ContainerSvcContainer
     */
    keeps?: Array<ContainerSvcKeep>;
    /**
     * Labels are metadata tags assigned to the container.
     * @type {Array<ContainerSvcLabel>}
     * @memberof ContainerSvcContainer
     */
    labels?: Array<ContainerSvcLabel>;
    /**
     * Names are the human-readable aliases assigned to the container.
     * @type {Array<string>}
     * @memberof ContainerSvcContainer
     */
    names?: Array<string>;
    /**
     * Network contains networking-related information for the container.
     * @type {ContainerSvcNetwork}
     * @memberof ContainerSvcContainer
     */
    network?: ContainerSvcNetwork;
    /**
     * Node Id
     * Please see the documentation for the envar OB_NODE_ID
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    nodeId?: string;
    /**
     * Ports maps host ports (keys) to container ports (values).
     * @type {Array<ContainerSvcPortMapping>}
     * @memberof ContainerSvcContainer
     */
    ports?: Array<ContainerSvcPortMapping>;
    /**
     * Resources defines CPU, memory, and disk constraints for the container.
     * @type {ContainerSvcResources}
     * @memberof ContainerSvcContainer
     */
    resources?: ContainerSvcResources;
    /**
     * Runtime specifies the container runtime (e.g., Docker, containerd, etc.).
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    runtime?: string;
    /**
     * Status indicates the current state of the container (e.g., running, stopped).
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    status?: string;
    /**
     * Volumes mounted by the container.
     * @type {Array<ContainerSvcVolume>}
     * @memberof ContainerSvcContainer
     */
    volumes?: Array<ContainerSvcVolume>;
}

/**
 * Check if a given object implements the ContainerSvcContainer interface.
 */
export function instanceOfContainerSvcContainer(value: object): value is ContainerSvcContainer {
    return true;
}

export function ContainerSvcContainerFromJSON(json: any): ContainerSvcContainer {
    return ContainerSvcContainerFromJSONTyped(json, false);
}

export function ContainerSvcContainerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcContainer {
    if (json == null) {
        return json;
    }
    return {
        
        'assets': json['assets'] == null ? undefined : ((json['assets'] as Array<any>).map(ContainerSvcAssetFromJSON)),
        'capabilities': json['capabilities'] == null ? undefined : ContainerSvcCapabilitiesFromJSON(json['capabilities']),
        'envs': json['envs'] == null ? undefined : ((json['envs'] as Array<any>).map(ContainerSvcEnvVarFromJSON)),
        'hash': json['hash'] == null ? undefined : json['hash'],
        'id': json['id'] == null ? undefined : json['id'],
        'image': json['image'] == null ? undefined : json['image'],
        'keeps': json['keeps'] == null ? undefined : ((json['keeps'] as Array<any>).map(ContainerSvcKeepFromJSON)),
        'labels': json['labels'] == null ? undefined : ((json['labels'] as Array<any>).map(ContainerSvcLabelFromJSON)),
        'names': json['names'] == null ? undefined : json['names'],
        'network': json['network'] == null ? undefined : ContainerSvcNetworkFromJSON(json['network']),
        'nodeId': json['nodeId'] == null ? undefined : json['nodeId'],
        'ports': json['ports'] == null ? undefined : ((json['ports'] as Array<any>).map(ContainerSvcPortMappingFromJSON)),
        'resources': json['resources'] == null ? undefined : ContainerSvcResourcesFromJSON(json['resources']),
        'runtime': json['runtime'] == null ? undefined : json['runtime'],
        'status': json['status'] == null ? undefined : json['status'],
        'volumes': json['volumes'] == null ? undefined : ((json['volumes'] as Array<any>).map(ContainerSvcVolumeFromJSON)),
    };
}

export function ContainerSvcContainerToJSON(json: any): ContainerSvcContainer {
    return ContainerSvcContainerToJSONTyped(json, false);
}

export function ContainerSvcContainerToJSONTyped(value?: ContainerSvcContainer | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'assets': value['assets'] == null ? undefined : ((value['assets'] as Array<any>).map(ContainerSvcAssetToJSON)),
        'capabilities': ContainerSvcCapabilitiesToJSON(value['capabilities']),
        'envs': value['envs'] == null ? undefined : ((value['envs'] as Array<any>).map(ContainerSvcEnvVarToJSON)),
        'hash': value['hash'],
        'id': value['id'],
        'image': value['image'],
        'keeps': value['keeps'] == null ? undefined : ((value['keeps'] as Array<any>).map(ContainerSvcKeepToJSON)),
        'labels': value['labels'] == null ? undefined : ((value['labels'] as Array<any>).map(ContainerSvcLabelToJSON)),
        'names': value['names'],
        'network': ContainerSvcNetworkToJSON(value['network']),
        'nodeId': value['nodeId'],
        'ports': value['ports'] == null ? undefined : ((value['ports'] as Array<any>).map(ContainerSvcPortMappingToJSON)),
        'resources': ContainerSvcResourcesToJSON(value['resources']),
        'runtime': value['runtime'],
        'status': value['status'],
        'volumes': value['volumes'] == null ? undefined : ((value['volumes'] as Array<any>).map(ContainerSvcVolumeToJSON)),
    };
}

