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
import type { PromptSvcPromptType } from './PromptSvcPromptType';
import type { ModelSvcArchitectures } from './ModelSvcArchitectures';
/**
 *
 * @export
 * @interface ModelSvcPlatform
 */
export interface ModelSvcPlatform {
    /**
     *
     * @type {ModelSvcArchitectures}
     * @memberof ModelSvcPlatform
     */
    architectures?: ModelSvcArchitectures;
    /**
     *
     * @type {string}
     * @memberof ModelSvcPlatform
     */
    id?: string;
    /**
     *
     * @type {string}
     * @memberof ModelSvcPlatform
     */
    name?: string;
    /**
     * List of prompt types that the AI engine supports.
     * @type {Array<PromptSvcPromptType>}
     * @memberof ModelSvcPlatform
     */
    types?: Array<PromptSvcPromptType>;
    /**
     *
     * @type {number}
     * @memberof ModelSvcPlatform
     */
    version?: number;
}
/**
 * Check if a given object implements the ModelSvcPlatform interface.
 */
export declare function instanceOfModelSvcPlatform(value: object): value is ModelSvcPlatform;
export declare function ModelSvcPlatformFromJSON(json: any): ModelSvcPlatform;
export declare function ModelSvcPlatformFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcPlatform;
export declare function ModelSvcPlatformToJSON(json: any): ModelSvcPlatform;
export declare function ModelSvcPlatformToJSONTyped(value?: ModelSvcPlatform | null, ignoreDiscriminator?: boolean): any;
