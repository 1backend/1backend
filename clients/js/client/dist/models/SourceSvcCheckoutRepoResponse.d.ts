/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
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
export declare function instanceOfSourceSvcCheckoutRepoResponse(value: object): value is SourceSvcCheckoutRepoResponse;
export declare function SourceSvcCheckoutRepoResponseFromJSON(json: any): SourceSvcCheckoutRepoResponse;
export declare function SourceSvcCheckoutRepoResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SourceSvcCheckoutRepoResponse;
export declare function SourceSvcCheckoutRepoResponseToJSON(json: any): SourceSvcCheckoutRepoResponse;
export declare function SourceSvcCheckoutRepoResponseToJSONTyped(value?: SourceSvcCheckoutRepoResponse | null, ignoreDiscriminator?: boolean): any;
