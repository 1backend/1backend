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
import type { SecretSvcSecret } from './SecretSvcSecret';
/**
 *
 * @export
 * @interface SecretSvcListSecretsResponse
 */
export interface SecretSvcListSecretsResponse {
    /**
     *
     * @type {Array<SecretSvcSecret>}
     * @memberof SecretSvcListSecretsResponse
     */
    secrets?: Array<SecretSvcSecret>;
}
/**
 * Check if a given object implements the SecretSvcListSecretsResponse interface.
 */
export declare function instanceOfSecretSvcListSecretsResponse(value: object): value is SecretSvcListSecretsResponse;
export declare function SecretSvcListSecretsResponseFromJSON(json: any): SecretSvcListSecretsResponse;
export declare function SecretSvcListSecretsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SecretSvcListSecretsResponse;
export declare function SecretSvcListSecretsResponseToJSON(json: any): SecretSvcListSecretsResponse;
export declare function SecretSvcListSecretsResponseToJSONTyped(value?: SecretSvcListSecretsResponse | null, ignoreDiscriminator?: boolean): any;
