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
import type { ModelSvcKeep } from './ModelSvcKeep';
import type { ModelSvcEnvVar } from './ModelSvcEnvVar';
/**
 *
 * @export
 * @interface ModelSvcContainer
 */
export interface ModelSvcContainer {
    /**
     * Environment variables to be passed to the container (e.g., "DEVICES=all").
     * @type {Array<ModelSvcEnvVar>}
     * @memberof ModelSvcContainer
     */
    envars?: Array<ModelSvcEnvVar>;
    /**
     * Template for constructing the container image name.
     * @type {string}
     * @memberof ModelSvcContainer
     */
    imageTemplate?: string;
    /**
     * List of container paths that should persist across restarts.
     * @type {Array<ModelSvcKeep>}
     * @memberof ModelSvcContainer
     */
    keeps?: Array<ModelSvcKeep>;
    /**
     * Internal port exposed by the container.
     * @type {number}
     * @memberof ModelSvcContainer
     */
    port?: number;
}
/**
 * Check if a given object implements the ModelSvcContainer interface.
 */
export declare function instanceOfModelSvcContainer(value: object): value is ModelSvcContainer;
export declare function ModelSvcContainerFromJSON(json: any): ModelSvcContainer;
export declare function ModelSvcContainerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcContainer;
export declare function ModelSvcContainerToJSON(json: any): ModelSvcContainer;
export declare function ModelSvcContainerToJSONTyped(value?: ModelSvcContainer | null, ignoreDiscriminator?: boolean): any;
