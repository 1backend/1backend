/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { ContainerSvcKeepFromJSON, ContainerSvcKeepToJSON, } from './ContainerSvcKeep';
import { ContainerSvcAssetFromJSON, ContainerSvcAssetToJSON, } from './ContainerSvcAsset';
import { ContainerSvcNetworkFromJSON, ContainerSvcNetworkToJSON, } from './ContainerSvcNetwork';
import { ContainerSvcCapabilitiesFromJSON, ContainerSvcCapabilitiesToJSON, } from './ContainerSvcCapabilities';
import { ContainerSvcVolumeFromJSON, ContainerSvcVolumeToJSON, } from './ContainerSvcVolume';
import { ContainerSvcLabelFromJSON, ContainerSvcLabelToJSON, } from './ContainerSvcLabel';
import { ContainerSvcEnvVarFromJSON, ContainerSvcEnvVarToJSON, } from './ContainerSvcEnvVar';
import { ContainerSvcResourcesFromJSON, ContainerSvcResourcesToJSON, } from './ContainerSvcResources';
import { ContainerSvcPortMappingFromJSON, ContainerSvcPortMappingToJSON, } from './ContainerSvcPortMapping';
/**
 * Check if a given object implements the ContainerSvcContainer interface.
 */
export function instanceOfContainerSvcContainer(value) {
    return true;
}
export function ContainerSvcContainerFromJSON(json) {
    return ContainerSvcContainerFromJSONTyped(json, false);
}
export function ContainerSvcContainerFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'assets': json['assets'] == null ? undefined : (json['assets'].map(ContainerSvcAssetFromJSON)),
        'capabilities': json['capabilities'] == null ? undefined : ContainerSvcCapabilitiesFromJSON(json['capabilities']),
        'envs': json['envs'] == null ? undefined : (json['envs'].map(ContainerSvcEnvVarFromJSON)),
        'hash': json['hash'] == null ? undefined : json['hash'],
        'id': json['id'] == null ? undefined : json['id'],
        'image': json['image'] == null ? undefined : json['image'],
        'keeps': json['keeps'] == null ? undefined : (json['keeps'].map(ContainerSvcKeepFromJSON)),
        'labels': json['labels'] == null ? undefined : (json['labels'].map(ContainerSvcLabelFromJSON)),
        'names': json['names'] == null ? undefined : json['names'],
        'network': json['network'] == null ? undefined : ContainerSvcNetworkFromJSON(json['network']),
        'nodeId': json['nodeId'] == null ? undefined : json['nodeId'],
        'ports': json['ports'] == null ? undefined : (json['ports'].map(ContainerSvcPortMappingFromJSON)),
        'resources': json['resources'] == null ? undefined : ContainerSvcResourcesFromJSON(json['resources']),
        'runtime': json['runtime'] == null ? undefined : json['runtime'],
        'status': json['status'] == null ? undefined : json['status'],
        'volumes': json['volumes'] == null ? undefined : (json['volumes'].map(ContainerSvcVolumeFromJSON)),
    };
}
export function ContainerSvcContainerToJSON(json) {
    return ContainerSvcContainerToJSONTyped(json, false);
}
export function ContainerSvcContainerToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'assets': value['assets'] == null ? undefined : (value['assets'].map(ContainerSvcAssetToJSON)),
        'capabilities': ContainerSvcCapabilitiesToJSON(value['capabilities']),
        'envs': value['envs'] == null ? undefined : (value['envs'].map(ContainerSvcEnvVarToJSON)),
        'hash': value['hash'],
        'id': value['id'],
        'image': value['image'],
        'keeps': value['keeps'] == null ? undefined : (value['keeps'].map(ContainerSvcKeepToJSON)),
        'labels': value['labels'] == null ? undefined : (value['labels'].map(ContainerSvcLabelToJSON)),
        'names': value['names'],
        'network': ContainerSvcNetworkToJSON(value['network']),
        'nodeId': value['nodeId'],
        'ports': value['ports'] == null ? undefined : (value['ports'].map(ContainerSvcPortMappingToJSON)),
        'resources': ContainerSvcResourcesToJSON(value['resources']),
        'runtime': value['runtime'],
        'status': value['status'],
        'volumes': value['volumes'] == null ? undefined : (value['volumes'].map(ContainerSvcVolumeToJSON)),
    };
}
