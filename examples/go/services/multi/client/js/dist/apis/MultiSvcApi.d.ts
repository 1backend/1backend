/**
 * Multi Svc
 * An example service for bootstrapping.
 *
 * The version of the OpenAPI document: 0.3.0-rc.8
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import * as runtime from '../runtime';
import type { MultiSvcCountPetsResponse } from '../models/index';
export interface CountPetsRequest {
    body?: object;
}
/**
 *
 */
export declare class MultiSvcApi extends runtime.BaseAPI {
    /**
     * Count the pets living in the Multi Svc.
     * Count Pets
     */
    countPetsRaw(requestParameters: CountPetsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<MultiSvcCountPetsResponse>>;
    /**
     * Count the pets living in the Multi Svc.
     * Count Pets
     */
    countPets(requestParameters?: CountPetsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<MultiSvcCountPetsResponse>;
}
