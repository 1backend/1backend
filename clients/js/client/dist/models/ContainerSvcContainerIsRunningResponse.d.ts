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
 * @interface ContainerSvcContainerIsRunningResponse
 */
export interface ContainerSvcContainerIsRunningResponse {
    /**
     *
     * @type {boolean}
     * @memberof ContainerSvcContainerIsRunningResponse
     */
    isRunning: boolean;
}
/**
 * Check if a given object implements the ContainerSvcContainerIsRunningResponse interface.
 */
export declare function instanceOfContainerSvcContainerIsRunningResponse(value: object): value is ContainerSvcContainerIsRunningResponse;
export declare function ContainerSvcContainerIsRunningResponseFromJSON(json: any): ContainerSvcContainerIsRunningResponse;
export declare function ContainerSvcContainerIsRunningResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcContainerIsRunningResponse;
export declare function ContainerSvcContainerIsRunningResponseToJSON(json: any): ContainerSvcContainerIsRunningResponse;
export declare function ContainerSvcContainerIsRunningResponseToJSONTyped(value?: ContainerSvcContainerIsRunningResponse | null, ignoreDiscriminator?: boolean): any;
