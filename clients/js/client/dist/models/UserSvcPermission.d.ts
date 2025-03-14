/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface UserSvcPermission
 */
export interface UserSvcPermission {
    /**
     *
     * @type {string}
     * @memberof UserSvcPermission
     */
    createdAt?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcPermission
     */
    description?: string;
    /**
     * eg. "user.viewer"
     * @type {string}
     * @memberof UserSvcPermission
     */
    id?: string;
    /**
     * eg. "User Viewer"
     * @type {string}
     * @memberof UserSvcPermission
     */
    name?: string;
    /**
     * Service who owns the permission
     *
     * Uncertain if this aligns with the system's use of slugs.
     * Issue encountered: I renamed Docker Svc to Container Svc in two steps (by mistake).
     * The name/slug had already changed to "container-svc," but data was still being saved
     * in the "dockerSvcCredentials" table.
     * After renaming the tables as well, I hit a "cannot update unowned permission" error
     * because ownership relies on this field rather than the user slug. YMMV.
     * @type {string}
     * @memberof UserSvcPermission
     */
    ownerId?: string;
    /**
     *
     * @type {string}
     * @memberof UserSvcPermission
     */
    updatedAt?: string;
}
/**
 * Check if a given object implements the UserSvcPermission interface.
 */
export declare function instanceOfUserSvcPermission(value: object): value is UserSvcPermission;
export declare function UserSvcPermissionFromJSON(json: any): UserSvcPermission;
export declare function UserSvcPermissionFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcPermission;
export declare function UserSvcPermissionToJSON(json: any): UserSvcPermission;
export declare function UserSvcPermissionToJSONTyped(value?: UserSvcPermission | null, ignoreDiscriminator?: boolean): any;
