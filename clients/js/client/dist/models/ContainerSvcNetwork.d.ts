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
/**
 *
 * @export
 * @interface ContainerSvcNetwork
 */
export interface ContainerSvcNetwork {
    /**
     * IPAddress is the assigned IP address of the container.
     * @type {string}
     * @memberof ContainerSvcNetwork
     */
    ipAddress?: string;
    /**
     * MacAddress is the container's MAC address if applicable.
     * @type {string}
     * @memberof ContainerSvcNetwork
     */
    macAddress?: string;
    /**
     * Mode specifies the container's network mode (e.g., bridge, host, none, custom).
     * @type {string}
     * @memberof ContainerSvcNetwork
     */
    mode?: string;
}
/**
 * Check if a given object implements the ContainerSvcNetwork interface.
 */
export declare function instanceOfContainerSvcNetwork(value: object): value is ContainerSvcNetwork;
export declare function ContainerSvcNetworkFromJSON(json: any): ContainerSvcNetwork;
export declare function ContainerSvcNetworkFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcNetwork;
export declare function ContainerSvcNetworkToJSON(json: any): ContainerSvcNetwork;
export declare function ContainerSvcNetworkToJSONTyped(value?: ContainerSvcNetwork | null, ignoreDiscriminator?: boolean): any;
