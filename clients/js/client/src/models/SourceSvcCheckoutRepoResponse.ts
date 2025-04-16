/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface SourceSvcCheckoutRepoResponse
 */
export interface SourceSvcCheckoutRepoResponse {
    /**
     * Directory where the repository was checked out
     * @type {string}
     * @memberof SourceSvcCheckoutRepoResponse
     */
    dir?: string;
}

/**
 * Check if a given object implements the SourceSvcCheckoutRepoResponse interface.
 */
export function instanceOfSourceSvcCheckoutRepoResponse(value: object): value is SourceSvcCheckoutRepoResponse {
    return true;
}

export function SourceSvcCheckoutRepoResponseFromJSON(json: any): SourceSvcCheckoutRepoResponse {
    return SourceSvcCheckoutRepoResponseFromJSONTyped(json, false);
}

export function SourceSvcCheckoutRepoResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SourceSvcCheckoutRepoResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'dir': json['dir'] == null ? undefined : json['dir'],
    };
}

export function SourceSvcCheckoutRepoResponseToJSON(json: any): SourceSvcCheckoutRepoResponse {
    return SourceSvcCheckoutRepoResponseToJSONTyped(json, false);
}

export function SourceSvcCheckoutRepoResponseToJSONTyped(value?: SourceSvcCheckoutRepoResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'dir': value['dir'],
    };
}

