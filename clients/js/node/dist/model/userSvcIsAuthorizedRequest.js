/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class UserSvcIsAuthorizedRequest {
    static getAttributeTypeMap() {
        return UserSvcIsAuthorizedRequest.attributeTypeMap;
    }
}
UserSvcIsAuthorizedRequest.discriminator = undefined;
UserSvcIsAuthorizedRequest.attributeTypeMap = [
    {
        "name": "contactsGranted",
        "baseName": "contactsGranted",
        "type": "Array<string>"
    },
    {
        "name": "grantedSlugs",
        "baseName": "grantedSlugs",
        "type": "Array<string>"
    }
];
