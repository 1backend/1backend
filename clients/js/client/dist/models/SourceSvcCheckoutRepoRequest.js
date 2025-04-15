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
 * Check if a given object implements the SourceSvcCheckoutRepoRequest interface.
 */
export function instanceOfSourceSvcCheckoutRepoRequest(value) {
    return true;
}
export function SourceSvcCheckoutRepoRequestFromJSON(json) {
    return SourceSvcCheckoutRepoRequestFromJSONTyped(json, false);
}
export function SourceSvcCheckoutRepoRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'password': json['password'] == null ? undefined : json['password'],
        'ssh_key': json['ssh_key'] == null ? undefined : json['ssh_key'],
        'ssh_key_pwd': json['ssh_key_pwd'] == null ? undefined : json['ssh_key_pwd'],
        'token': json['token'] == null ? undefined : json['token'],
        'url': json['url'] == null ? undefined : json['url'],
        'username': json['username'] == null ? undefined : json['username'],
        'version': json['version'] == null ? undefined : json['version'],
    };
}
export function SourceSvcCheckoutRepoRequestToJSON(json) {
    return SourceSvcCheckoutRepoRequestToJSONTyped(json, false);
}
export function SourceSvcCheckoutRepoRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'password': value['password'],
        'ssh_key': value['ssh_key'],
        'ssh_key_pwd': value['ssh_key_pwd'],
        'token': value['token'],
        'url': value['url'],
        'username': value['username'],
        'version': value['version'],
    };
}
