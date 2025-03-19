/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { RegistrySvcAPISpecFromJSON, RegistrySvcAPISpecToJSON, } from './RegistrySvcAPISpec';
import { RegistrySvcPortMappingFromJSON, RegistrySvcPortMappingToJSON, } from './RegistrySvcPortMapping';
import { RegistrySvcImageSpecFromJSON, RegistrySvcImageSpecToJSON, } from './RegistrySvcImageSpec';
import { RegistrySvcRepositorySpecFromJSON, RegistrySvcRepositorySpecToJSON, } from './RegistrySvcRepositorySpec';
import { RegistrySvcClientFromJSON, RegistrySvcClientToJSON, } from './RegistrySvcClient';
import { RegistrySvcEnvVarFromJSON, RegistrySvcEnvVarToJSON, } from './RegistrySvcEnvVar';
/**
 * Check if a given object implements the RegistrySvcDefinition interface.
 */
export function instanceOfRegistrySvcDefinition(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    return true;
}
export function RegistrySvcDefinitionFromJSON(json) {
    return RegistrySvcDefinitionFromJSONTyped(json, false);
}
export function RegistrySvcDefinitionFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'apiSpecs': json['apiSpecs'] == null ? undefined : (json['apiSpecs'].map(RegistrySvcAPISpecFromJSON)),
        'clients': json['clients'] == null ? undefined : (json['clients'].map(RegistrySvcClientFromJSON)),
        'envars': json['envars'] == null ? undefined : (json['envars'].map(RegistrySvcEnvVarFromJSON)),
        'id': json['id'],
        'image': json['image'] == null ? undefined : RegistrySvcImageSpecFromJSON(json['image']),
        'ports': json['ports'] == null ? undefined : (json['ports'].map(RegistrySvcPortMappingFromJSON)),
        'repository': json['repository'] == null ? undefined : RegistrySvcRepositorySpecFromJSON(json['repository']),
    };
}
export function RegistrySvcDefinitionToJSON(json) {
    return RegistrySvcDefinitionToJSONTyped(json, false);
}
export function RegistrySvcDefinitionToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'apiSpecs': value['apiSpecs'] == null ? undefined : (value['apiSpecs'].map(RegistrySvcAPISpecToJSON)),
        'clients': value['clients'] == null ? undefined : (value['clients'].map(RegistrySvcClientToJSON)),
        'envars': value['envars'] == null ? undefined : (value['envars'].map(RegistrySvcEnvVarToJSON)),
        'id': value['id'],
        'image': RegistrySvcImageSpecToJSON(value['image']),
        'ports': value['ports'] == null ? undefined : (value['ports'].map(RegistrySvcPortMappingToJSON)),
        'repository': RegistrySvcRepositorySpecToJSON(value['repository']),
    };
}
