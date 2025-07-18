/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ModelSvcAsset } from './ModelSvcAsset';
/**
 *
 * @export
 * @interface ModelSvcModel
 */
export interface ModelSvcModel {
    /**
     *
     * @type {Array<ModelSvcAsset>}
     * @memberof ModelSvcModel
     */
    assets?: Array<ModelSvcAsset>;
    /**
     *
     * @type {number}
     * @memberof ModelSvcModel
     */
    bits?: number;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    description?: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    extension?: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    flavour?: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    fullName?: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    id: string;
    /**
     *
     * @type {number}
     * @memberof ModelSvcModel
     */
    maxBits?: number;
    /**
     *
     * @type {number}
     * @memberof ModelSvcModel
     */
    maxRam?: number;
    /**
     *
     * @type {Array<string>}
     * @memberof ModelSvcModel
     */
    mirrors?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    name: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    parameters?: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    platformId: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    promptTemplate?: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    quality?: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    quantComment?: string;
    /**
     *
     * @type {number}
     * @memberof ModelSvcModel
     */
    size?: number;
    /**
     *
     * @type {Array<string>}
     * @memberof ModelSvcModel
     */
    tags?: Array<string>;
    /**
     *
     * @type {boolean}
     * @memberof ModelSvcModel
     */
    uncensored?: boolean;
    /**
     *
     * @type {string}
     * @memberof ModelSvcModel
     */
    version?: string;
}
/**
 * Check if a given object implements the ModelSvcModel interface.
 */
export declare function instanceOfModelSvcModel(value: object): value is ModelSvcModel;
export declare function ModelSvcModelFromJSON(json: any): ModelSvcModel;
export declare function ModelSvcModelFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcModel;
export declare function ModelSvcModelToJSON(json: any): ModelSvcModel;
export declare function ModelSvcModelToJSONTyped(value?: ModelSvcModel | null, ignoreDiscriminator?: boolean): any;
