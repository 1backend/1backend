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
 * @interface RegistrySvcRegisterInstanceRequest
 */
export interface RegistrySvcRegisterInstanceRequest {
    /**
     * The ID of the deployment that this instance is an instance of.
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    deploymentId?: string;
    /**
     * Host of the instance address. Required if URL is not provided
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    host?: string;
    /**
     *
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    id?: string;
    /**
     * IP of the instance address. Optional: to register by IP instead of host
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    ip?: string;
    /**
     * Path of the instance address. Optional (e.g., "/api")
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    path?: string;
    /**
     * Port of the instance address. Required if URL is not provided
     * @type {number}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    port?: number;
    /**
     * Scheme of the instance address. Required if URL is not provided.
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    scheme?: string;
    /**
     * Full address URL of the instance.
     * @type {string}
     * @memberof RegistrySvcRegisterInstanceRequest
     */
    url: string;
}
/**
 * Check if a given object implements the RegistrySvcRegisterInstanceRequest interface.
 */
export declare function instanceOfRegistrySvcRegisterInstanceRequest(value: object): value is RegistrySvcRegisterInstanceRequest;
export declare function RegistrySvcRegisterInstanceRequestFromJSON(json: any): RegistrySvcRegisterInstanceRequest;
export declare function RegistrySvcRegisterInstanceRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcRegisterInstanceRequest;
export declare function RegistrySvcRegisterInstanceRequestToJSON(json: any): RegistrySvcRegisterInstanceRequest;
export declare function RegistrySvcRegisterInstanceRequestToJSONTyped(value?: RegistrySvcRegisterInstanceRequest | null, ignoreDiscriminator?: boolean): any;
