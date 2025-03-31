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
import * as runtime from '../runtime';
import type { UserSvcAssignPermissionsRequest, UserSvcChangePasswordRequest, UserSvcCreateOrganizationRequest, UserSvcCreateOrganizationResponse, UserSvcCreateRoleRequest, UserSvcCreateRoleResponse, UserSvcCreateUserRequest, UserSvcGetPermissionsResponse, UserSvcGetPublicKeyResponse, UserSvcIsAuthorizedRequest, UserSvcIsAuthorizedResponse, UserSvcListGrantsRequest, UserSvcListGrantsResponse, UserSvcListInvitesRequest, UserSvcListInvitesResponse, UserSvcListRolesResponse, UserSvcListUsersRequest, UserSvcListUsersResponse, UserSvcLoginRequest, UserSvcLoginResponse, UserSvcReadUserByTokenResponse, UserSvcRegisterRequest, UserSvcRegisterResponse, UserSvcResetPasswordRequest, UserSvcSaveGrantsRequest, UserSvcSaveInvitesRequest, UserSvcSaveInvitesResponse, UserSvcSavePermissionsRequest, UserSvcSavePermissionsResponse, UserSvcSaveProfileRequest, UserSvcSetRolePermissionsRequest } from '../models/index';
export interface AddUserToOrganizationRequest {
    organizationId: string;
    userId: string;
    body?: object;
}
export interface AssignPermissionsRequest {
    body: UserSvcAssignPermissionsRequest;
}
export interface AssignRoleRequest {
    userId: string;
    roleId: string;
    body?: object;
}
export interface ChangePasswordRequest {
    body: UserSvcChangePasswordRequest;
}
export interface CreateOrganizationRequest {
    body: UserSvcCreateOrganizationRequest;
}
export interface CreateRoleRequest {
    body: UserSvcCreateRoleRequest;
}
export interface CreateUserRequest {
    body: UserSvcCreateUserRequest;
}
export interface DeleteRoleRequest {
    roleId: string;
}
export interface DeleteUserRequest {
    userId: string;
}
export interface GetPermissionsByRoleRequest {
    roleId: string;
}
export interface IsAuthorizedRequest {
    permissionId: string;
    body?: UserSvcIsAuthorizedRequest;
}
export interface ListGrantsRequest {
    body: UserSvcListGrantsRequest;
}
export interface ListInvitesRequest {
    body: UserSvcListInvitesRequest;
}
export interface ListUsersRequest {
    body?: UserSvcListUsersRequest;
}
export interface LoginRequest {
    body: UserSvcLoginRequest;
}
export interface RegisterRequest {
    body: UserSvcRegisterRequest;
}
export interface RemoveUserFromOrganizationRequest {
    organizationId: string;
    userId: string;
    body?: object;
}
export interface ResetPasswordRequest {
    userId: string;
    body: UserSvcResetPasswordRequest;
}
export interface SaveGrantsRequest {
    body: UserSvcSaveGrantsRequest;
}
export interface SaveInvitesRequest {
    body: UserSvcSaveInvitesRequest;
}
export interface SavePermissionsRequest {
    body: UserSvcSavePermissionsRequest;
}
export interface SaveSelfRequest {
    userId: string;
    body: UserSvcSaveProfileRequest;
}
export interface SaveUserRequest {
    userId: string;
    body: UserSvcSaveProfileRequest;
}
export interface SetRolePermissionRequest {
    roleId: string;
    body: UserSvcSetRolePermissionsRequest;
}
/**
 *
 */
export declare class UserSvcApi extends runtime.BaseAPI {
    /**
     * Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.
     * Add a User to an Organization
     */
    addUserToOrganizationRaw(requestParameters: AddUserToOrganizationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.
     * Add a User to an Organization
     */
    addUserToOrganization(requestParameters: AddUserToOrganizationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Assign permissions to roles.  Requires the `user-svc:permission:assign` permission.
     * Assign Permissions
     */
    assignPermissionsRaw(requestParameters: AssignPermissionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Assign permissions to roles.  Requires the `user-svc:permission:assign` permission.
     * Assign Permissions
     */
    assignPermissions(requestParameters: AssignPermissionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Assigns a role to a user. The caller can only assign roles they own. A user \"owns\" a role in the following cases: - A static role where the role ID is prefixed with the caller\'s slug. - Any dynamic or static role where the caller is an admin.  Examples: - A user with the slug \"joe-doe\" owns roles like \"joe-doe:any-custom-role\". - A user with any slug who has the role \"my-service:admin\" owns \"my-service:user\". - A user with any slug who has the role \"user-svc:org:{%orgId}:admin\" owns \"user-svc:org:{%orgId}:user\".
     * Assign Role
     */
    assignRoleRaw(requestParameters: AssignRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Assigns a role to a user. The caller can only assign roles they own. A user \"owns\" a role in the following cases: - A static role where the role ID is prefixed with the caller\'s slug. - Any dynamic or static role where the caller is an admin.  Examples: - A user with the slug \"joe-doe\" owns roles like \"joe-doe:any-custom-role\". - A user with any slug who has the role \"my-service:admin\" owns \"my-service:user\". - A user with any slug who has the role \"user-svc:org:{%orgId}:admin\" owns \"user-svc:org:{%orgId}:user\".
     * Assign Role
     */
    assignRole(requestParameters: AssignRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Allows an authenticated user to change their own password.
     * Change User Password
     */
    changePasswordRaw(requestParameters: ChangePasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Allows an authenticated user to change their own password.
     * Change User Password
     */
    changePassword(requestParameters: ChangePasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Allows a logged-in user to create a new organization. The user initiating the request will be assigned the role of admin for that organization. The initiating user will receive a dynamic role in the format `user-svc:org:{organizationId}:admin`, where `{organizationId}` is a unique identifier for the created organization. Dynamic roles are generated based on specific user-resource associations (in this case the resource being the organization), offering more flexible permission management compared to static roles.
     * Create an Organization
     */
    createOrganizationRaw(requestParameters: CreateOrganizationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcCreateOrganizationResponse>>;
    /**
     * Allows a logged-in user to create a new organization. The user initiating the request will be assigned the role of admin for that organization. The initiating user will receive a dynamic role in the format `user-svc:org:{organizationId}:admin`, where `{organizationId}` is a unique identifier for the created organization. Dynamic roles are generated based on specific user-resource associations (in this case the resource being the organization), offering more flexible permission management compared to static roles.
     * Create an Organization
     */
    createOrganization(requestParameters: CreateOrganizationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcCreateOrganizationResponse>;
    /**
     * Create a new role. <b>The role ID must be prefixed by the caller\'s slug.</b> Eg. if the caller\'s slug is `petstore-svc` the role should look like `petstore-svc:admin`. The user account who creates the role will become the owner of that role, and only the owner will be able to edit the role.  Requires the `user-svc:role:create` permission.
     * Create a New Role
     */
    createRoleRaw(requestParameters: CreateRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcCreateRoleResponse>>;
    /**
     * Create a new role. <b>The role ID must be prefixed by the caller\'s slug.</b> Eg. if the caller\'s slug is `petstore-svc` the role should look like `petstore-svc:admin`. The user account who creates the role will become the owner of that role, and only the owner will be able to edit the role.  Requires the `user-svc:role:create` permission.
     * Create a New Role
     */
    createRole(requestParameters: CreateRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcCreateRoleResponse>;
    /**
     * Allows an authenticated administrator to create a new user with specified details.
     * Create a New User
     */
    createUserRaw(requestParameters: CreateUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Allows an authenticated administrator to create a new user with specified details.
     * Create a New User
     */
    createUser(requestParameters: CreateUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Delete a role based on the role ID.
     * Delete a Role
     */
    deleteRoleRaw(requestParameters: DeleteRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Delete a role based on the role ID.
     * Delete a Role
     */
    deleteRole(requestParameters: DeleteRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Delete a user based on the user ID.
     * Delete a User
     */
    deleteUserRaw(requestParameters: DeleteUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Delete a user based on the user ID.
     * Delete a User
     */
    deleteUser(requestParameters: DeleteUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Retrieve permissions associated with a specific role ID.
     * Get Permissions by Role
     */
    getPermissionsByRoleRaw(requestParameters: GetPermissionsByRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcGetPermissionsResponse>>;
    /**
     * Retrieve permissions associated with a specific role ID.
     * Get Permissions by Role
     */
    getPermissionsByRole(requestParameters: GetPermissionsByRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcGetPermissionsResponse>;
    /**
     * Get the public key to parse and verify the JWT.
     * Get Public Key
     */
    getPublicKeyRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcGetPublicKeyResponse>>;
    /**
     * Get the public key to parse and verify the JWT.
     * Get Public Key
     */
    getPublicKey(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcGetPublicKeyResponse>;
    /**
     * Verify whether a user has a specific permission. Ideally, this endpoint should rarely be used, as the JWT token already includes all user roles. Caching the `Get Permissions by Role` responses allows services to determine user authorization without repeatedly calling this endpoint.
     * Is Authorized
     */
    isAuthorizedRaw(requestParameters: IsAuthorizedRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcIsAuthorizedResponse>>;
    /**
     * Verify whether a user has a specific permission. Ideally, this endpoint should rarely be used, as the JWT token already includes all user roles. Caching the `Get Permissions by Role` responses allows services to determine user authorization without repeatedly calling this endpoint.
     * Is Authorized
     */
    isAuthorized(requestParameters: IsAuthorizedRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcIsAuthorizedResponse>;
    /**
     * List grants.  Grants define which slugs are assigned specific permissions, overriding the default configuration.  Requires the `user-svc:grant:view` permission.
     * List Grants
     */
    listGrantsRaw(requestParameters: ListGrantsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcListGrantsResponse>>;
    /**
     * List grants.  Grants define which slugs are assigned specific permissions, overriding the default configuration.  Requires the `user-svc:grant:view` permission.
     * List Grants
     */
    listGrants(requestParameters: ListGrantsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcListGrantsResponse>;
    /**
     * List user invites stored in the database.
     * List Invites
     */
    listInvitesRaw(requestParameters: ListInvitesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcListInvitesResponse>>;
    /**
     * List user invites stored in the database.
     * List Invites
     */
    listInvites(requestParameters: ListInvitesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcListInvitesResponse>;
    /**
     * Retrieve all roles from the user service.
     * List Roles
     */
    listRolesRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcListRolesResponse>>;
    /**
     * Retrieve all roles from the user service.
     * List Roles
     */
    listRoles(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcListRolesResponse>;
    /**
     * Fetches a list of users with optional query filters and pagination.
     * List Users
     */
    listUsersRaw(requestParameters: ListUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcListUsersResponse>>;
    /**
     * Fetches a list of users with optional query filters and pagination.
     * List Users
     */
    listUsers(requestParameters?: ListUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcListUsersResponse>;
    /**
     * Authenticates a user and returns a token.
     * Login
     */
    loginRaw(requestParameters: LoginRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcLoginResponse>>;
    /**
     * Authenticates a user and returns a token.
     * Login
     */
    login(requestParameters: LoginRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcLoginResponse>;
    /**
     * Retrieve user information based on an authentication token.
     * Read User by Token
     */
    readUserByTokenRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcReadUserByTokenResponse>>;
    /**
     * Retrieve user information based on an authentication token.
     * Read User by Token
     */
    readUserByToken(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcReadUserByTokenResponse>;
    /**
     * Register a new user with a name, email, and password.
     * Register
     */
    registerRaw(requestParameters: RegisterRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcRegisterResponse>>;
    /**
     * Register a new user with a name, email, and password.
     * Register
     */
    register(requestParameters: RegisterRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcRegisterResponse>;
    /**
     * Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.
     * Remove a User from an Organization
     */
    removeUserFromOrganizationRaw(requestParameters: RemoveUserFromOrganizationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.
     * Remove a User from an Organization
     */
    removeUserFromOrganization(requestParameters: RemoveUserFromOrganizationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Allows an administrator to change a user\'s password.
     * Reset Password
     */
    resetPasswordRaw(requestParameters: ResetPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Allows an administrator to change a user\'s password.
     * Reset Password
     */
    resetPassword(requestParameters: ResetPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Save grants.  Grants define which slugs are assigned specific permissions, overriding the default configuration.  Requires the `user-svc:grant:create` permission.
     * Save Grants
     */
    saveGrantsRaw(requestParameters: SaveGrantsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Save grants.  Grants define which slugs are assigned specific permissions, overriding the default configuration.  Requires the `user-svc:grant:create` permission.
     * Save Grants
     */
    saveGrants(requestParameters: SaveGrantsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Invite a list of users by contact ID to acquire a role. Works on future or current users. A user can only invite an other user to a role if the user owns that role.  A user \"owns\" a role in the following cases: - A static role where the role ID is prefixed with the caller\'s slug. - Any dynamic or static role where the caller is an admin.  Examples: - A user with the slug \"joe-doe\" owns roles like \"joe-doe:any-custom-role\". - A user with any slug who has the role \"my-service:admin\" owns \"my-service:user\". - A user with any slug who has the role \"user-svc:org:{%orgId}:admin\" owns \"user-svc:org:{%orgId}:user\".
     * Save Invites
     */
    saveInvitesRaw(requestParameters: SaveInvitesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcSaveInvitesResponse>>;
    /**
     * Invite a list of users by contact ID to acquire a role. Works on future or current users. A user can only invite an other user to a role if the user owns that role.  A user \"owns\" a role in the following cases: - A static role where the role ID is prefixed with the caller\'s slug. - Any dynamic or static role where the caller is an admin.  Examples: - A user with the slug \"joe-doe\" owns roles like \"joe-doe:any-custom-role\". - A user with any slug who has the role \"my-service:admin\" owns \"my-service:user\". - A user with any slug who has the role \"user-svc:org:{%orgId}:admin\" owns \"user-svc:org:{%orgId}:user\".
     * Save Invites
     */
    saveInvites(requestParameters: SaveInvitesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcSaveInvitesResponse>;
    /**
     * Creates or updates a list of permissions. <b>The permission ID must be prefixed by the callers slug.</b> Eg. if the owner\'s slug is `petstore-svc` the permission should look like `petstore-svc:pet:edit`.  Requires the `user-svc:permission:create` permission.
     * Save Permissions
     */
    savePermissionsRaw(requestParameters: SavePermissionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UserSvcSavePermissionsResponse>>;
    /**
     * Creates or updates a list of permissions. <b>The permission ID must be prefixed by the callers slug.</b> Eg. if the owner\'s slug is `petstore-svc` the permission should look like `petstore-svc:pet:edit`.  Requires the `user-svc:permission:create` permission.
     * Save Permissions
     */
    savePermissions(requestParameters: SavePermissionsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UserSvcSavePermissionsResponse>;
    /**
     * Save user\'s own profile information.
     * Save User Profile
     */
    saveSelfRaw(requestParameters: SaveSelfRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Save user\'s own profile information.
     * Save User Profile
     */
    saveSelf(requestParameters: SaveSelfRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Save user information based on the provided user ID. It is intended for admins, because it uses the `user-svc:user:edit` permission which only admins have. For a user to edit its own profile, see saveSelf.
     * Save User
     */
    saveUserRaw(requestParameters: SaveUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Save user information based on the provided user ID. It is intended for admins, because it uses the `user-svc:user:edit` permission which only admins have. For a user to edit its own profile, see saveSelf.
     * Save User
     */
    saveUser(requestParameters: SaveUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Set permissions for a specified role. The caller can add permissions it owns to any role. If the caller tries to add a permission it doesn\'t own to a role, `StatusBadRequest` will be returned.
     * Set Role Permissions
     */
    setRolePermissionRaw(requestParameters: SetRolePermissionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Set permissions for a specified role. The caller can add permissions it owns to any role. If the caller tries to add a permission it doesn\'t own to a role, `StatusBadRequest` will be returned.
     * Set Role Permissions
     */
    setRolePermission(requestParameters: SetRolePermissionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
}
