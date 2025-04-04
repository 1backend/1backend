/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ContainerSvcBuildImageRequest
 */
export interface ContainerSvcBuildImageRequest {
    /**
     * ContextPath is the local path to the build context
     * @type {string}
     * @memberof ContainerSvcBuildImageRequest
     */
    contextPath: string;
    /**
     * DockerfilePath is the local path to the Dockerfile
     * @type {string}
     * @memberof ContainerSvcBuildImageRequest
     */
    dockerfilePath?: string;
    /**
     * Name is the name of the image to build
     * @type {string}
     * @memberof ContainerSvcBuildImageRequest
     */
    name: string;
}
/**
 * Check if a given object implements the ContainerSvcBuildImageRequest interface.
 */
export declare function instanceOfContainerSvcBuildImageRequest(value: object): value is ContainerSvcBuildImageRequest;
export declare function ContainerSvcBuildImageRequestFromJSON(json: any): ContainerSvcBuildImageRequest;
export declare function ContainerSvcBuildImageRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcBuildImageRequest;
export declare function ContainerSvcBuildImageRequestToJSON(json: any): ContainerSvcBuildImageRequest;
export declare function ContainerSvcBuildImageRequestToJSONTyped(value?: ContainerSvcBuildImageRequest | null, ignoreDiscriminator?: boolean): any;
