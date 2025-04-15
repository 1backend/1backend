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
/**
 * Check if a given object implements the UserSvcListPermitsRequest interface.
 */
export function instanceOfUserSvcListPermitsRequest(value) {
    return true;
}
export function UserSvcListPermitsRequestFromJSON(json) {
    return UserSvcListPermitsRequestFromJSONTyped(json, false);
}
export function UserSvcListPermitsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'permission': json['permission'] == null ? undefined : json['permission'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}
export function UserSvcListPermitsRequestToJSON(json) {
    return UserSvcListPermitsRequestToJSONTyped(json, false);
}
export function UserSvcListPermitsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'permission': value['permission'],
        'slug': value['slug'],
    };
}
