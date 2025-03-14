/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { SecretSvcSecret } from './SecretSvcSecret';
/**
 *
 * @export
 * @interface SecretSvcSaveSecretsRequest
 */
export interface SecretSvcSaveSecretsRequest {
    /**
     *
     * @type {Array<SecretSvcSecret>}
     * @memberof SecretSvcSaveSecretsRequest
     */
    secrets?: Array<SecretSvcSecret>;
}
/**
 * Check if a given object implements the SecretSvcSaveSecretsRequest interface.
 */
export declare function instanceOfSecretSvcSaveSecretsRequest(value: object): value is SecretSvcSaveSecretsRequest;
export declare function SecretSvcSaveSecretsRequestFromJSON(json: any): SecretSvcSaveSecretsRequest;
export declare function SecretSvcSaveSecretsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcSaveSecretsRequest;
export declare function SecretSvcSaveSecretsRequestToJSON(json: any): SecretSvcSaveSecretsRequest;
export declare function SecretSvcSaveSecretsRequestToJSONTyped(value?: SecretSvcSaveSecretsRequest | null, ignoreDiscriminator?: boolean): any;
