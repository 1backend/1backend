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
import type { RegistrySvcDefinition } from './RegistrySvcDefinition';
/**
 *
 * @export
 * @interface RegistrySvcSaveDefinitionRequest
 */
export interface RegistrySvcSaveDefinitionRequest {
    /**
     *
     * @type {RegistrySvcDefinition}
     * @memberof RegistrySvcSaveDefinitionRequest
     */
    definition?: RegistrySvcDefinition;
}
/**
 * Check if a given object implements the RegistrySvcSaveDefinitionRequest interface.
 */
export declare function instanceOfRegistrySvcSaveDefinitionRequest(value: object): value is RegistrySvcSaveDefinitionRequest;
export declare function RegistrySvcSaveDefinitionRequestFromJSON(json: any): RegistrySvcSaveDefinitionRequest;
export declare function RegistrySvcSaveDefinitionRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcSaveDefinitionRequest;
export declare function RegistrySvcSaveDefinitionRequestToJSON(json: any): RegistrySvcSaveDefinitionRequest;
export declare function RegistrySvcSaveDefinitionRequestToJSONTyped(value?: RegistrySvcSaveDefinitionRequest | null, ignoreDiscriminator?: boolean): any;
