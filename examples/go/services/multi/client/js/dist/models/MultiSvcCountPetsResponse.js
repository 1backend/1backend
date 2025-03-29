/* tslint:disable */
/* eslint-disable */
/**
 * Multi Svc
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
 * Check if a given object implements the MultiSvcCountPetsResponse interface.
 */
export function instanceOfMultiSvcCountPetsResponse(value) {
    if (!('petCount' in value) || value['petCount'] === undefined)
        return false;
    return true;
}
export function MultiSvcCountPetsResponseFromJSON(json) {
    return MultiSvcCountPetsResponseFromJSONTyped(json, false);
}
export function MultiSvcCountPetsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'petCount': json['petCount'],
    };
}
export function MultiSvcCountPetsResponseToJSON(json) {
    return MultiSvcCountPetsResponseToJSONTyped(json, false);
}
export function MultiSvcCountPetsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'petCount': value['petCount'],
    };
}
