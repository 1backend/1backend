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
import { DatastoreQueryFromJSON, DatastoreQueryToJSON, } from './DatastoreQuery';
/**
 * Check if a given object implements the UserSvcListUsersRequest interface.
 */
export function instanceOfUserSvcListUsersRequest(value) {
    return true;
}
export function UserSvcListUsersRequestFromJSON(json) {
    return UserSvcListUsersRequestFromJSONTyped(json, false);
}
export function UserSvcListUsersRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contactId': json['contactId'] == null ? undefined : json['contactId'],
        'query': json['query'] == null ? undefined : DatastoreQueryFromJSON(json['query']),
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}
export function UserSvcListUsersRequestToJSON(json) {
    return UserSvcListUsersRequestToJSONTyped(json, false);
}
export function UserSvcListUsersRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'contactId': value['contactId'],
        'query': DatastoreQueryToJSON(value['query']),
        'userId': value['userId'],
    };
}
