/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export declare class ContainerSvcLog {
    /**
    * ContainerId is the raw underlying container ID. Eg. Docker container id. Node local.
    */
    'containerId'?: string;
    'content'?: string;
    'createdAt'?: string;
    'id'?: string;
    /**
    * Node Id Please see the documentation for the envar OPENORCH_NODE_ID
    */
    'nodeId'?: string;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
