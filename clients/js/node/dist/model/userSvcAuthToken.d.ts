/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export declare class UserSvcAuthToken {
    /**
    * Active tokens contain the most up-to-date information. When a user\'s role changes—due to role assignment, organization creation/assignment, etc.—all existing tokens are marked inactive. Active tokens are reused during login, while inactive tokens are retained for historical reference.
    */
    'active'?: boolean;
    'createdAt'?: string;
    'deletedAt'?: string;
    'id'?: string;
    'token': string;
    'updatedAt'?: string;
    'userId': string;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
