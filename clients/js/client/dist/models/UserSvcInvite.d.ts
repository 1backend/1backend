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
/**
 *
 * @export
 * @interface UserSvcInvite
 */
export interface UserSvcInvite {
    /**
     *
     * @type {string}
     * @memberof UserSvcInvite
     */
    appliedAt?: string;
    /**
     * ContactId represents the recipient of the invite.
     * If the user is already registered, the role is assigned immediately;
     * otherwise, it is applied upon registration.
     * @type {string}
     * @memberof UserSvcInvite
     */
    contactId: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcInvite
     */
    createdAt: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcInvite
     */
    deletedAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcInvite
     */
    expiresAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcInvite
     */
    id: string;
    /**
     * OwnerIds specifies the users who created the invite.
     * If you create an invite that already exists for a given role and contact ID,
     * you get added to the list of owners.
     * @type {Array<string>}
     * @memberof UserSvcInvite
     */
    ownerIds: Array<string>;
    /**
     * RoleId specifies the role to be assigned to the ContactId.
     * Callers can only assign roles they own, identified by their service slug
     * (e.g., if "my-service" creates an invite, the role must be "my-service:admin").
     * Dynamic organization roles can also be assigned
     * (e.g., "user-svc:org:{%orgId}:admin" or "user-svc:org:{%orgId}:user"),
     * but in this case, the caller must be an admin of the target organization.
     * @type {string}
     * @memberof UserSvcInvite
     */
    roleId: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcInvite
     */
    updatedAt?: string;
}
/**
 * Check if a given object implements the UserSvcInvite interface.
 */
export declare function instanceOfUserSvcInvite(value: object): value is UserSvcInvite;
export declare function UserSvcInviteFromJSON(json: any): UserSvcInvite;
export declare function UserSvcInviteFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcInvite;
export declare function UserSvcInviteToJSON(json: any): UserSvcInvite;
export declare function UserSvcInviteToJSONTyped(value?: UserSvcInvite | null, ignoreDiscriminator?: boolean): any;
