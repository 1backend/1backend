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
import { UserSvcAuthToken } from './userSvcAuthToken';
import { UserSvcOrganization } from './userSvcOrganization';
export declare class UserSvcSaveOrganizationResponse {
    'organization': UserSvcOrganization;
    /**
    * Due to the nature of JWT tokens, the token must be refreshed after creating an organization, as dynamic organization roles are embedded in it.
    */
    'token': UserSvcAuthToken;
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
