/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
import * as runtime from '../runtime';
import { UserSvcChangePasswordRequestToJSON, UserSvcCreateUserRequestToJSON, UserSvcGetPublicKeyResponseFromJSON, UserSvcHasPermissionRequestToJSON, UserSvcHasPermissionResponseFromJSON, UserSvcListEnrollsRequestToJSON, UserSvcListEnrollsResponseFromJSON, UserSvcListOrganizationsRequestToJSON, UserSvcListOrganizationsResponseFromJSON, UserSvcListPermissionsResponseFromJSON, UserSvcListPermitsRequestToJSON, UserSvcListPermitsResponseFromJSON, UserSvcListUsersRequestToJSON, UserSvcListUsersResponseFromJSON, UserSvcLoginRequestToJSON, UserSvcLoginResponseFromJSON, UserSvcReadSelfResponseFromJSON, UserSvcRegisterRequestToJSON, UserSvcRegisterResponseFromJSON, UserSvcResetPasswordRequestToJSON, UserSvcSaveEnrollsRequestToJSON, UserSvcSaveEnrollsResponseFromJSON, UserSvcSaveOrganizationRequestToJSON, UserSvcSaveOrganizationResponseFromJSON, UserSvcSavePermitsRequestToJSON, UserSvcSaveProfileRequestToJSON, UserSvcSaveSelfRequestToJSON, } from '../models/index';
/**
 *
 */
export class UserSvcApi extends runtime.BaseAPI {
    /**
     * Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.
     * Add a User to an Organization
     */
    addUserToOrganizationRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['organizationId'] == null) {
                throw new runtime.RequiredError('organizationId', 'Required parameter "organizationId" was null or undefined when calling addUserToOrganization().');
            }
            if (requestParameters['userId'] == null) {
                throw new runtime.RequiredError('userId', 'Required parameter "userId" was null or undefined when calling addUserToOrganization().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/organization/{organizationId}/user/{userId}`.replace(`{${"organizationId"}}`, encodeURIComponent(String(requestParameters['organizationId']))).replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: requestParameters['body'],
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.
     * Add a User to an Organization
     */
    addUserToOrganization(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.addUserToOrganizationRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Allows an authenticated user to change their own password.
     * Change User Password
     */
    changePasswordRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling changePassword().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/change-password`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcChangePasswordRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Allows an authenticated user to change their own password.
     * Change User Password
     */
    changePassword(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.changePasswordRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Allows an authenticated administrator to create a new user with specified details.
     * Create a New User
     */
    createUserRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling createUser().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/user`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcCreateUserRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Allows an authenticated administrator to create a new user with specified details.
     * Create a New User
     */
    createUser(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.createUserRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Delete a user based on the user ID.
     * Delete a User
     */
    deleteUserRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['userId'] == null) {
                throw new runtime.RequiredError('userId', 'Required parameter "userId" was null or undefined when calling deleteUser().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/user/{userId}`.replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
                method: 'DELETE',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Delete a user based on the user ID.
     * Delete a User
     */
    deleteUser(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.deleteUserRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Get the public key to parse and verify the JWT.
     * Get Public Key
     */
    getPublicKeyRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            const response = yield this.request({
                path: `/user-svc/public-key`,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcGetPublicKeyResponseFromJSON(jsonValue));
        });
    }
    /**
     * Get the public key to parse and verify the JWT.
     * Get Public Key
     */
    getPublicKey(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.getPublicKeyRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Check whether the caller user has a specific permission. Ideally, this endpoint should rarely be used, as the JWT token already includes all user roles. Caching the `List Permissions` and `List Permits` responses allows services to determine user authorization without repeatedly calling this endpoint.
     * Has Permission
     */
    hasPermissionRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['permission'] == null) {
                throw new runtime.RequiredError('permission', 'Required parameter "permission" was null or undefined when calling hasPermission().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/self/has/{permission}`.replace(`{${"permission"}}`, encodeURIComponent(String(requestParameters['permission']))),
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcHasPermissionRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcHasPermissionResponseFromJSON(jsonValue));
        });
    }
    /**
     * Check whether the caller user has a specific permission. Ideally, this endpoint should rarely be used, as the JWT token already includes all user roles. Caching the `List Permissions` and `List Permits` responses allows services to determine user authorization without repeatedly calling this endpoint.
     * Has Permission
     */
    hasPermission(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.hasPermissionRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * List enrolls. Role, user ID or contact ID must be specified. Caller can only list enrolls of roles they own.
     * List Enrolls
     */
    listEnrollsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling listEnrolls().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/enrolls`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcListEnrollsRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcListEnrollsResponseFromJSON(jsonValue));
        });
    }
    /**
     * List enrolls. Role, user ID or contact ID must be specified. Caller can only list enrolls of roles they own.
     * List Enrolls
     */
    listEnrolls(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.listEnrollsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Requires the `user-svc:organization:view` permission, that only admins have by default.
     * List Organizations
     */
    listOrganizationsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling listOrganizations().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/organizations`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcListOrganizationsRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcListOrganizationsResponseFromJSON(jsonValue));
        });
    }
    /**
     * Requires the `user-svc:organization:view` permission, that only admins have by default.
     * List Organizations
     */
    listOrganizations(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.listOrganizationsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * List permissions by roles. Caller can only list permissions for roles they have.
     * List Permissions
     */
    listPermissionsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['roleId'] == null) {
                throw new runtime.RequiredError('roleId', 'Required parameter "roleId" was null or undefined when calling listPermissions().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/permissions`.replace(`{${"roleId"}}`, encodeURIComponent(String(requestParameters['roleId']))),
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcListPermissionsResponseFromJSON(jsonValue));
        });
    }
    /**
     * List permissions by roles. Caller can only list permissions for roles they have.
     * List Permissions
     */
    listPermissions(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.listPermissionsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * List permits. Requires the `user-svc:permit:view` permission, which only admins have by default. &todo Users should be able to list permits referring to them.
     * List Permits
     */
    listPermitsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling listPermits().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/permits`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcListPermitsRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcListPermitsResponseFromJSON(jsonValue));
        });
    }
    /**
     * List permits. Requires the `user-svc:permit:view` permission, which only admins have by default. &todo Users should be able to list permits referring to them.
     * List Permits
     */
    listPermits(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.listPermitsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Fetches a list of users with optional query filters and pagination. Requires the `user-svc:user:view` permission that only admins have by default.
     * List Users
     */
    listUsersRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/users`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcListUsersRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcListUsersResponseFromJSON(jsonValue));
        });
    }
    /**
     * Fetches a list of users with optional query filters and pagination. Requires the `user-svc:user:view` permission that only admins have by default.
     * List Users
     */
    listUsers() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.listUsersRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Authenticates a user and returns a token.
     * Login
     */
    loginRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling login().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            const response = yield this.request({
                path: `/user-svc/login`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcLoginRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcLoginResponseFromJSON(jsonValue));
        });
    }
    /**
     * Authenticates a user and returns a token.
     * Login
     */
    login(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.loginRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieves user information based on the authentication token in the request header. Typically called by single-page applications during the initial page load. While some details (such as roles, slug, user ID, and active organization ID) can be extracted from the JWT, this endpoint returns additional data, including the full user object and associated organizations.
     * Read Self
     */
    readSelfRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/self`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcReadSelfResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieves user information based on the authentication token in the request header. Typically called by single-page applications during the initial page load. While some details (such as roles, slug, user ID, and active organization ID) can be extracted from the JWT, this endpoint returns additional data, including the full user object and associated organizations.
     * Read Self
     */
    readSelf(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.readSelfRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Register a new user with a name, email, and password.
     * Register
     */
    registerRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling register().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            const response = yield this.request({
                path: `/user-svc/register`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcRegisterRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcRegisterResponseFromJSON(jsonValue));
        });
    }
    /**
     * Register a new user with a name, email, and password.
     * Register
     */
    register(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.registerRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.
     * Remove a User from an Organization
     */
    removeUserFromOrganizationRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['organizationId'] == null) {
                throw new runtime.RequiredError('organizationId', 'Required parameter "organizationId" was null or undefined when calling removeUserFromOrganization().');
            }
            if (requestParameters['userId'] == null) {
                throw new runtime.RequiredError('userId', 'Required parameter "userId" was null or undefined when calling removeUserFromOrganization().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/organization/{organizationId}/user/{userId}`.replace(`{${"organizationId"}}`, encodeURIComponent(String(requestParameters['organizationId']))).replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
                method: 'DELETE',
                headers: headerParameters,
                query: queryParameters,
                body: requestParameters['body'],
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.
     * Remove a User from an Organization
     */
    removeUserFromOrganization(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.removeUserFromOrganizationRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Allows an administrator to change a user\'s password.
     * Reset Password
     */
    resetPasswordRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['userId'] == null) {
                throw new runtime.RequiredError('userId', 'Required parameter "userId" was null or undefined when calling resetPassword().');
            }
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling resetPassword().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/{userId}/reset-password`.replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcResetPasswordRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Allows an administrator to change a user\'s password.
     * Reset Password
     */
    resetPassword(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.resetPasswordRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Enroll a list of users by contact or user Id to acquire a role. Works on future or current users.  A user can only enroll an other user to a role if the user owns that role.  A user \"owns\" a role in the following cases: - A static role where the role ID is prefixed with the caller\'s slug. - Any dynamic or static role where the caller is an admin.  Examples: - A user with the slug \"joe-doe\" owns roles like \"joe-doe:any-custom-role\". - A user with any slug who has the role \"my-service:admin\" owns \"my-service:user\". - A user with any slug who has the role \"user-svc:org:{%orgId}:admin\" owns \"user-svc:org:{%orgId}:user\".
     * Save Enrolls
     */
    saveEnrollsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling saveEnrolls().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/enrolls`,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcSaveEnrollsRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcSaveEnrollsResponseFromJSON(jsonValue));
        });
    }
    /**
     * Enroll a list of users by contact or user Id to acquire a role. Works on future or current users.  A user can only enroll an other user to a role if the user owns that role.  A user \"owns\" a role in the following cases: - A static role where the role ID is prefixed with the caller\'s slug. - Any dynamic or static role where the caller is an admin.  Examples: - A user with the slug \"joe-doe\" owns roles like \"joe-doe:any-custom-role\". - A user with any slug who has the role \"my-service:admin\" owns \"my-service:user\". - A user with any slug who has the role \"user-svc:org:{%orgId}:admin\" owns \"user-svc:org:{%orgId}:user\".
     * Save Enrolls
     */
    saveEnrolls(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.saveEnrollsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Allows a logged-in user to save an organization. The user initiating the request will be assigned the role of admin for that organization. The initiating user will receive a dynamic role in the format `user-svc:org:{organizationId}:admin`, where `{organizationId}` is a unique identifier for the saved organization. Dynamic roles are generated based on specific user-resource associations (in this case the resource being the organization), offering more flexible permission management compared to static roles.
     * Save an Organization
     */
    saveOrganizationRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling saveOrganization().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/organization`,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcSaveOrganizationRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => UserSvcSaveOrganizationResponseFromJSON(jsonValue));
        });
    }
    /**
     * Allows a logged-in user to save an organization. The user initiating the request will be assigned the role of admin for that organization. The initiating user will receive a dynamic role in the format `user-svc:org:{organizationId}:admin`, where `{organizationId}` is a unique identifier for the saved organization. Dynamic roles are generated based on specific user-resource associations (in this case the resource being the organization), offering more flexible permission management compared to static roles.
     * Save an Organization
     */
    saveOrganization(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.saveOrganizationRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Save permits. // @Description Permits give access to users with certain slugs and roles to permissions.
     * Save Permits
     */
    savePermitsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling savePermits().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/permits`,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcSavePermitsRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Save permits. // @Description Permits give access to users with certain slugs and roles to permissions.
     * Save Permits
     */
    savePermits(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.savePermitsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Save user\'s own profile information.
     * Save User Profile
     */
    saveSelfRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling saveSelf().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/self`,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcSaveSelfRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Save user\'s own profile information.
     * Save User Profile
     */
    saveSelf(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.saveSelfRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Save user information based on the provided user ID. It is intended for admins, because it uses the `user-svc:user:edit` permission which only admins have. For a user to edit its own profile, see saveSelf.
     * Save User
     */
    saveUserRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['userId'] == null) {
                throw new runtime.RequiredError('userId', 'Required parameter "userId" was null or undefined when calling saveUser().');
            }
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling saveUser().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/user-svc/user/{userId}`.replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: UserSvcSaveProfileRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Save user information based on the provided user ID. It is intended for admins, because it uses the `user-svc:user:edit` permission which only admins have. For a user to edit its own profile, see saveSelf.
     * Save User
     */
    saveUser(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.saveUserRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
}
