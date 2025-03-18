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
 * @interface ContainerSvcKeep
 */
export interface ContainerSvcKeep {
    /**
     * Path is the absolute path inside the container for the folder that should persist across restarts.
     * @type {string}
     * @memberof ContainerSvcKeep
     */
    path: string;
    /**
     * ReadOnly indicates whether the keep is read-only.
     * @type {boolean}
     * @memberof ContainerSvcKeep
     */
    readOnly?: boolean;
}
/**
 * Check if a given object implements the ContainerSvcKeep interface.
 */
export declare function instanceOfContainerSvcKeep(value: object): value is ContainerSvcKeep;
export declare function ContainerSvcKeepFromJSON(json: any): ContainerSvcKeep;
export declare function ContainerSvcKeepFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcKeep;
export declare function ContainerSvcKeepToJSON(json: any): ContainerSvcKeep;
export declare function ContainerSvcKeepToJSONTyped(value?: ContainerSvcKeep | null, ignoreDiscriminator?: boolean): any;
