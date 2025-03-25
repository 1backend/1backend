/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { UserSvcAuthToken } from './UserSvcAuthToken';
import type { UserSvcOrganization } from './UserSvcOrganization';
/**
 *
 * @export
 * @interface UserSvcCreateOrganizationResponse
 */
export interface UserSvcCreateOrganizationResponse {
    /**
     *
     * @type {UserSvcOrganization}
     * @memberof UserSvcCreateOrganizationResponse
     */
    organization: UserSvcOrganization;
    /**
     * Due to the nature of JWT tokens, the token must be refreshed after
     * creating an organization, as dynamic organization roles are embedded in it.
     * @type {UserSvcAuthToken}
     * @memberof UserSvcCreateOrganizationResponse
     */
    token: UserSvcAuthToken;
}
/**
 * Check if a given object implements the UserSvcCreateOrganizationResponse interface.
 */
export declare function instanceOfUserSvcCreateOrganizationResponse(value: object): value is UserSvcCreateOrganizationResponse;
export declare function UserSvcCreateOrganizationResponseFromJSON(json: any): UserSvcCreateOrganizationResponse;
export declare function UserSvcCreateOrganizationResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcCreateOrganizationResponse;
export declare function UserSvcCreateOrganizationResponseToJSON(json: any): UserSvcCreateOrganizationResponse;
export declare function UserSvcCreateOrganizationResponseToJSONTyped(value?: UserSvcCreateOrganizationResponse | null, ignoreDiscriminator?: boolean): any;
