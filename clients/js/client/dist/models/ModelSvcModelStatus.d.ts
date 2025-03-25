/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ModelSvcModelStatus
 */
export interface ModelSvcModelStatus {
    /**
     *
     * @type {string}
     * @memberof ModelSvcModelStatus
     */
    address: string;
    /**
     *
     * @type {boolean}
     * @memberof ModelSvcModelStatus
     */
    assetsReady: boolean;
    /**
     * Running triggers onModelLaunch on the frontend.
     * 	Running is true when the model is both running and answering
     * 	- fully loaded.
     * @type {boolean}
     * @memberof ModelSvcModelStatus
     */
    running: boolean;
}
/**
 * Check if a given object implements the ModelSvcModelStatus interface.
 */
export declare function instanceOfModelSvcModelStatus(value: object): value is ModelSvcModelStatus;
export declare function ModelSvcModelStatusFromJSON(json: any): ModelSvcModelStatus;
export declare function ModelSvcModelStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcModelStatus;
export declare function ModelSvcModelStatusToJSON(json: any): ModelSvcModelStatus;
export declare function ModelSvcModelStatusToJSONTyped(value?: ModelSvcModelStatus | null, ignoreDiscriminator?: boolean): any;
