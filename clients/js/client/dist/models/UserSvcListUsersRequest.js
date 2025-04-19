/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcListUsersOrderByFieldFromJSON, UserSvcListUsersOrderByFieldToJSON, } from './UserSvcListUsersOrderByField';
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
        'afterTime': json['afterTime'] == null ? undefined : json['afterTime'],
        'contactId': json['contactId'] == null ? undefined : json['contactId'],
        'count': json['count'] == null ? undefined : json['count'],
        'ids': json['ids'] == null ? undefined : json['ids'],
        'limit': json['limit'] == null ? undefined : json['limit'],
        'offset': json['offset'] == null ? undefined : json['offset'],
        'orderByDesc': json['orderByDesc'] == null ? undefined : json['orderByDesc'],
        'orderByField': json['orderByField'] == null ? undefined : UserSvcListUsersOrderByFieldFromJSON(json['orderByField']),
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
        'afterTime': value['afterTime'],
        'contactId': value['contactId'],
        'count': value['count'],
        'ids': value['ids'],
        'limit': value['limit'],
        'offset': value['offset'],
        'orderByDesc': value['orderByDesc'],
        'orderByField': UserSvcListUsersOrderByFieldToJSON(value['orderByField']),
    };
}
