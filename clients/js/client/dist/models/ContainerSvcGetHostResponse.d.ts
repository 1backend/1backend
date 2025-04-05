/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ContainerSvcGetHostResponse
 */
export interface ContainerSvcGetHostResponse {
    /**
     *
     * @type {string}
     * @memberof ContainerSvcGetHostResponse
     */
    host: string;
}
/**
 * Check if a given object implements the ContainerSvcGetHostResponse interface.
 */
export declare function instanceOfContainerSvcGetHostResponse(value: object): value is ContainerSvcGetHostResponse;
export declare function ContainerSvcGetHostResponseFromJSON(json: any): ContainerSvcGetHostResponse;
export declare function ContainerSvcGetHostResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcGetHostResponse;
export declare function ContainerSvcGetHostResponseToJSON(json: any): ContainerSvcGetHostResponse;
export declare function ContainerSvcGetHostResponseToJSONTyped(value?: ContainerSvcGetHostResponse | null, ignoreDiscriminator?: boolean): any;
