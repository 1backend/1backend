/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface DeploySvcTargetRegion
 */
export interface DeploySvcTargetRegion {
    /**
     * Cluster or node where service should be deployed (e.g., "us-west1", "local-docker")
     * @type {string}
     * @memberof DeploySvcTargetRegion
     */
    cluster?: string;
    /**
     * Optional: Specific zone for the deployment
     * @type {string}
     * @memberof DeploySvcTargetRegion
     */
    zone?: string;
}
/**
 * Check if a given object implements the DeploySvcTargetRegion interface.
 */
export declare function instanceOfDeploySvcTargetRegion(value: object): value is DeploySvcTargetRegion;
export declare function DeploySvcTargetRegionFromJSON(json: any): DeploySvcTargetRegion;
export declare function DeploySvcTargetRegionFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcTargetRegion;
export declare function DeploySvcTargetRegionToJSON(json: any): DeploySvcTargetRegion;
export declare function DeploySvcTargetRegionToJSONTyped(value?: DeploySvcTargetRegion | null, ignoreDiscriminator?: boolean): any;
