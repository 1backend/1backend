/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
import { PolicySvcCheckRequestToJSON, PolicySvcCheckResponseFromJSON, PolicySvcUpsertInstanceRequestToJSON, } from '../models/index';
/**
 *
 */
export class PolicySvcApi extends runtime.BaseAPI {
    /**
     * Check records a resource access and returns if the access is allowed.
     * Check
     */
    checkRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling check().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/policy-svc/check`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: PolicySvcCheckRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => PolicySvcCheckResponseFromJSON(jsonValue));
        });
    }
    /**
     * Check records a resource access and returns if the access is allowed.
     * Check
     */
    check(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.checkRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Allows user to upsert a new policy instance based on a template.
     * Upsert an Instance
     */
    upsertInstanceRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['instanceId'] == null) {
                throw new runtime.RequiredError('instanceId', 'Required parameter "instanceId" was null or undefined when calling upsertInstance().');
            }
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling upsertInstance().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/policy-svc/instance/{instanceId}`.replace(`{${"instanceId"}}`, encodeURIComponent(String(requestParameters['instanceId']))),
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: PolicySvcUpsertInstanceRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Allows user to upsert a new policy instance based on a template.
     * Upsert an Instance
     */
    upsertInstance(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.upsertInstanceRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
}
