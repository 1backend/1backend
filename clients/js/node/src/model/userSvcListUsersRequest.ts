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

import { RequestFile } from './models';
import { DatastoreQuery } from './datastoreQuery';

export class UserSvcListUsersRequest {
    'contactId'?: string;
    /**
    * This should be used sparingly or not at all. Might deprecate.
    */
    'query'?: DatastoreQuery;
    'userId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "contactId",
            "baseName": "contactId",
            "type": "string"
        },
        {
            "name": "query",
            "baseName": "query",
            "type": "DatastoreQuery"
        },
        {
            "name": "userId",
            "baseName": "userId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcListUsersRequest.attributeTypeMap;
    }
}

