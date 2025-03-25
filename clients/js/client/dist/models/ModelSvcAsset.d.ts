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
/**
 *
 * @export
 * @interface ModelSvcAsset
 */
export interface ModelSvcAsset {
    /**
     *
     * @type {string}
     * @memberof ModelSvcAsset
     */
    envVarKey: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcAsset
     */
    url: string;
}
/**
 * Check if a given object implements the ModelSvcAsset interface.
 */
export declare function instanceOfModelSvcAsset(value: object): value is ModelSvcAsset;
export declare function ModelSvcAssetFromJSON(json: any): ModelSvcAsset;
export declare function ModelSvcAssetFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcAsset;
export declare function ModelSvcAssetToJSON(json: any): ModelSvcAsset;
export declare function ModelSvcAssetToJSONTyped(value?: ModelSvcAsset | null, ignoreDiscriminator?: boolean): any;
