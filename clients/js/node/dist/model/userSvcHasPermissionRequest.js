/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcHasPermissionRequest {
    static getAttributeTypeMap() {
        return UserSvcHasPermissionRequest.attributeTypeMap;
    }
}
UserSvcHasPermissionRequest.discriminator = undefined;
UserSvcHasPermissionRequest.attributeTypeMap = [
    {
        "name": "contactsPermited",
        "baseName": "contactsPermited",
        "type": "Array<string>"
    },
    {
        "name": "permittedSlugs",
        "baseName": "permittedSlugs",
        "type": "Array<string>"
    }
];
