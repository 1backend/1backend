/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import * as runtime from '../runtime';
import type { SourceSvcCheckoutRepoRequest, SourceSvcCheckoutRepoResponse } from '../models/index';
export interface CheckoutRepoRequest {
    body: SourceSvcCheckoutRepoRequest;
}
/**
 *
 */
export declare class SourceSvcApi extends runtime.BaseAPI {
    /**
     * Checkout a git repository over https or ssh at a specific version into a temporary directory. Performs a shallow clone with minimal history for faster checkout.
     * Checkout a git repository
     */
    checkoutRepoRaw(requestParameters: CheckoutRepoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SourceSvcCheckoutRepoResponse>>;
    /**
     * Checkout a git repository over https or ssh at a specific version into a temporary directory. Performs a shallow clone with minimal history for faster checkout.
     * Checkout a git repository
     */
    checkoutRepo(requestParameters: CheckoutRepoRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SourceSvcCheckoutRepoResponse>;
}
