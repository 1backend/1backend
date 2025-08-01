/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  PolicySvcCheckRequest,
  PolicySvcCheckResponse,
  PolicySvcErrorResponse,
  PolicySvcUpsertInstanceRequest,
} from '../models/index';
import {
    PolicySvcCheckRequestFromJSON,
    PolicySvcCheckRequestToJSON,
    PolicySvcCheckResponseFromJSON,
    PolicySvcCheckResponseToJSON,
    PolicySvcErrorResponseFromJSON,
    PolicySvcErrorResponseToJSON,
    PolicySvcUpsertInstanceRequestFromJSON,
    PolicySvcUpsertInstanceRequestToJSON,
} from '../models/index';

export interface CheckRequest {
    body: PolicySvcCheckRequest;
}

export interface UpsertInstanceRequest {
    instanceId: string;
    body: PolicySvcUpsertInstanceRequest;
}

/**
 * 
 */
export class PolicySvcApi extends runtime.BaseAPI {

    /**
     * Check records a resource access and returns if the access is allowed.
     * Check
     */
    async checkRaw(requestParameters: CheckRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PolicySvcCheckResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling check().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }


        let urlPath = `/policy-svc/check`;

        const response = await this.request({
            path: urlPath,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: PolicySvcCheckRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PolicySvcCheckResponseFromJSON(jsonValue));
    }

    /**
     * Check records a resource access and returns if the access is allowed.
     * Check
     */
    async check(requestParameters: CheckRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PolicySvcCheckResponse> {
        const response = await this.checkRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Allows user to upsert a new policy instance based on a template.
     * Upsert an Instance
     */
    async upsertInstanceRaw(requestParameters: UpsertInstanceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['instanceId'] == null) {
            throw new runtime.RequiredError(
                'instanceId',
                'Required parameter "instanceId" was null or undefined when calling upsertInstance().'
            );
        }

        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling upsertInstance().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }


        let urlPath = `/policy-svc/instance/{instanceId}`;
        urlPath = urlPath.replace(`{${"instanceId"}}`, encodeURIComponent(String(requestParameters['instanceId'])));

        const response = await this.request({
            path: urlPath,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: PolicySvcUpsertInstanceRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Allows user to upsert a new policy instance based on a template.
     * Upsert an Instance
     */
    async upsertInstance(requestParameters: UpsertInstanceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.upsertInstanceRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
