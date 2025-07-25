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
/**
 *
 * @export
 * @interface ContainerSvcAsset
 */
export interface ContainerSvcAsset {
    /**
     *
     * @type {string}
     * @memberof ContainerSvcAsset
     */
    envVarKey: string;
    /**
     *
     * @type {string}
     * @memberof ContainerSvcAsset
     */
    url: string;
}
/**
 * Check if a given object implements the ContainerSvcAsset interface.
 */
export declare function instanceOfContainerSvcAsset(value: object): value is ContainerSvcAsset;
export declare function ContainerSvcAssetFromJSON(json: any): ContainerSvcAsset;
export declare function ContainerSvcAssetFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcAsset;
export declare function ContainerSvcAssetToJSON(json: any): ContainerSvcAsset;
export declare function ContainerSvcAssetToJSONTyped(value?: ContainerSvcAsset | null, ignoreDiscriminator?: boolean): any;
