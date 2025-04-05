/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ContainerSvcVolume
 */
export interface ContainerSvcVolume {
    /**
     * Destination is the path inside the container.
     * @type {string}
     * @memberof ContainerSvcVolume
     */
    destination?: string;
    /**
     * ReadOnly indicates whether the mount is read-only.
     * @type {boolean}
     * @memberof ContainerSvcVolume
     */
    readOnly?: boolean;
    /**
     * Source is the host path or volume name.
     * @type {string}
     * @memberof ContainerSvcVolume
     */
    source?: string;
}
/**
 * Check if a given object implements the ContainerSvcVolume interface.
 */
export declare function instanceOfContainerSvcVolume(value: object): value is ContainerSvcVolume;
export declare function ContainerSvcVolumeFromJSON(json: any): ContainerSvcVolume;
export declare function ContainerSvcVolumeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcVolume;
export declare function ContainerSvcVolumeToJSON(json: any): ContainerSvcVolume;
export declare function ContainerSvcVolumeToJSONTyped(value?: ContainerSvcVolume | null, ignoreDiscriminator?: boolean): any;
