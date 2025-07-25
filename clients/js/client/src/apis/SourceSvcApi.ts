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
  SourceSvcCheckoutRepoRequest,
  SourceSvcCheckoutRepoResponse,
  SourceSvcErrorResponse,
} from '../models/index';
import {
    SourceSvcCheckoutRepoRequestFromJSON,
    SourceSvcCheckoutRepoRequestToJSON,
    SourceSvcCheckoutRepoResponseFromJSON,
    SourceSvcCheckoutRepoResponseToJSON,
    SourceSvcErrorResponseFromJSON,
    SourceSvcErrorResponseToJSON,
} from '../models/index';

export interface CheckoutRepoRequest {
    body: SourceSvcCheckoutRepoRequest;
}

/**
 * 
 */
export class SourceSvcApi extends runtime.BaseAPI {

    /**
     * Checkout a git repository over https or ssh at a specific version into a temporary directory. Performs a shallow clone with minimal history for faster checkout.
     * Checkout a git repository
     */
    async checkoutRepoRaw(requestParameters: CheckoutRepoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SourceSvcCheckoutRepoResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling checkoutRepo().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }


        let urlPath = `/source-svc/repo/checkout`;

        const response = await this.request({
            path: urlPath,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SourceSvcCheckoutRepoRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SourceSvcCheckoutRepoResponseFromJSON(jsonValue));
    }

    /**
     * Checkout a git repository over https or ssh at a specific version into a temporary directory. Performs a shallow clone with minimal history for faster checkout.
     * Checkout a git repository
     */
    async checkoutRepo(requestParameters: CheckoutRepoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SourceSvcCheckoutRepoResponse> {
        const response = await this.checkoutRepoRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
