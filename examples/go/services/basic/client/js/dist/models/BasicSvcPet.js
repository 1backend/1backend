/* tslint:disable */
/* eslint-disable */
/**
 * Basic Svc
 * An example service for bootstrapping.
 *
 * The version of the OpenAPI document: 0.3.0-rc.8
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the BasicSvcPet interface.
 */
export function instanceOfBasicSvcPet(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    return true;
}
export function BasicSvcPetFromJSON(json) {
    return BasicSvcPetFromJSONTyped(json, false);
}
export function BasicSvcPetFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'id': json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}
export function BasicSvcPetToJSON(json) {
    return BasicSvcPetToJSONTyped(json, false);
}
export function BasicSvcPetToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'createdAt': value['createdAt'],
        'id': value['id'],
        'name': value['name'],
        'updatedAt': value['updatedAt'],
    };
}
