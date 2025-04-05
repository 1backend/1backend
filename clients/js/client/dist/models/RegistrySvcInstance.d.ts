/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { RegistrySvcInstanceStatus } from './RegistrySvcInstanceStatus';
/**
 *
 * @export
 * @interface RegistrySvcInstance
 */
export interface RegistrySvcInstance {
    /**
     * The ID of the deployment that this instance is an instance of.
     * Only instances deployed by 1Backend have a DeploymentId.
     * Services can be deployed through other means (Docker Compose, K8s, anything),
     * in that case they self-register and will not have a DeploymentId.
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    deploymentId?: string;
    /**
     * Details
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    details?: string;
    /**
     * Host of the instance address. Required if URL is not provided
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    host?: string;
    /**
     * Required: ID of the instance
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    id: string;
    /**
     * IP of the instance address. Optional: to register by IP instead of host
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    ip?: string;
    /**
     * Last time the instance gave a sign of life
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    lastHeartbeat?: string;
    /**
     * NodeURL is the URL of the 1Backend server the instance is running on.
     * To have a NodeURL the instance must either:
     * - Be deployed by 1Backend
     * - Declare the 1Backend server URL when registering its instance
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    nodeUrl?: string;
    /**
     * Path of the instance address. Optional (e.g., "/api")
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    path?: string;
    /**
     * Port of the instance address. Required if URL is not provided
     * @type {number}
     * @memberof RegistrySvcInstance
     */
    port?: number;
    /**
     * Scheme of the instance address. Required if URL is not provided.
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    scheme?: string;
    /**
     * Slug of the account that owns this instance
     * Services that want to be proxied by their slug are advised to self register
     * their instance at startup.
     * Keep in mind, instances might be deployed by 1Backend yet they still won't be 1Backend services
     * and they won't have slugs. Think NGINX, MySQL, etc.
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    slug?: string;
    /**
     * Status
     * @type {RegistrySvcInstanceStatus}
     * @memberof RegistrySvcInstance
     */
    status: RegistrySvcInstanceStatus;
    /**
     * Tags are used to filter instances
     * @type {Array<string>}
     * @memberof RegistrySvcInstance
     */
    tags?: Array<string>;
    /**
     * Full address URL of the instance.
     * @type {string}
     * @memberof RegistrySvcInstance
     */
    url: string;
}
/**
 * Check if a given object implements the RegistrySvcInstance interface.
 */
export declare function instanceOfRegistrySvcInstance(value: object): value is RegistrySvcInstance;
export declare function RegistrySvcInstanceFromJSON(json: any): RegistrySvcInstance;
export declare function RegistrySvcInstanceFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcInstance;
export declare function RegistrySvcInstanceToJSON(json: any): RegistrySvcInstance;
export declare function RegistrySvcInstanceToJSONTyped(value?: RegistrySvcInstance | null, ignoreDiscriminator?: boolean): any;
