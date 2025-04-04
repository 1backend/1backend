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
 * @interface ModelSvcKeep
 */
export interface ModelSvcKeep {
    /**
     * Path is the absolute path inside the container for the folder that should persist across restarts.
     * @type {string}
     * @memberof ModelSvcKeep
     */
    path?: string;
    /**
     * ReadOnly indicates whether the keep is read-only.
     * @type {boolean}
     * @memberof ModelSvcKeep
     */
    readOnly?: boolean;
}
/**
 * Check if a given object implements the ModelSvcKeep interface.
 */
export declare function instanceOfModelSvcKeep(value: object): value is ModelSvcKeep;
export declare function ModelSvcKeepFromJSON(json: any): ModelSvcKeep;
export declare function ModelSvcKeepFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcKeep;
export declare function ModelSvcKeepToJSON(json: any): ModelSvcKeep;
export declare function ModelSvcKeepToJSONTyped(value?: ModelSvcKeep | null, ignoreDiscriminator?: boolean): any;
