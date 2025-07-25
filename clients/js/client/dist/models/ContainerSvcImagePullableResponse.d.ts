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
 * @interface ContainerSvcImagePullableResponse
 */
export interface ContainerSvcImagePullableResponse {
    /**
     *
     * @type {boolean}
     * @memberof ContainerSvcImagePullableResponse
     */
    pullable: boolean;
}
/**
 * Check if a given object implements the ContainerSvcImagePullableResponse interface.
 */
export declare function instanceOfContainerSvcImagePullableResponse(value: object): value is ContainerSvcImagePullableResponse;
export declare function ContainerSvcImagePullableResponseFromJSON(json: any): ContainerSvcImagePullableResponse;
export declare function ContainerSvcImagePullableResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcImagePullableResponse;
export declare function ContainerSvcImagePullableResponseToJSON(json: any): ContainerSvcImagePullableResponse;
export declare function ContainerSvcImagePullableResponseToJSONTyped(value?: ContainerSvcImagePullableResponse | null, ignoreDiscriminator?: boolean): any;
