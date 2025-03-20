/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { RegistrySvcAPISpec } from './RegistrySvcAPISpec';
import type { RegistrySvcPortMapping } from './RegistrySvcPortMapping';
import type { RegistrySvcImageSpec } from './RegistrySvcImageSpec';
import type { RegistrySvcRepositorySpec } from './RegistrySvcRepositorySpec';
import type { RegistrySvcClient } from './RegistrySvcClient';
import type { RegistrySvcEnvVar } from './RegistrySvcEnvVar';
/**
 *
 * @export
 * @interface RegistrySvcDefinition
 */
export interface RegistrySvcDefinition {
    /**
     * API Specs such as OpenAPI definitions etc.
     * @type {Array<RegistrySvcAPISpec>}
     * @memberof RegistrySvcDefinition
     */
    apiSpecs?: Array<RegistrySvcAPISpec>;
    /**
     * Programming language clients such as on npm or GitHub.
     * @type {Array<RegistrySvcClient>}
     * @memberof RegistrySvcDefinition
     */
    clients?: Array<RegistrySvcClient>;
    /**
     * Envars is a map of Renvironment variables that a deployment (see Deploy Svc Deployment) of this definition will REQUIRE to run.
     * E.g., {"DB_URL": "mysql://user:password@host:port/db"}
     * These will be injected into the service instances (see Registry Svc Instance) at runtime.
     * The value of a key here is the default value. The actual value can be overridden at deployment time.
     * @type {Array<RegistrySvcEnvVar>}
     * @memberof RegistrySvcDefinition
     */
    envars?: Array<RegistrySvcEnvVar>;
    /**
     *
     * @type {string}
     * @memberof RegistrySvcDefinition
     */
    id: string;
    /**
     * Container specifications for Docker, K8s, etc.
     * Use this to deploy already built images.
     * @type {RegistrySvcImageSpec}
     * @memberof RegistrySvcDefinition
     */
    image?: RegistrySvcImageSpec;
    /**
     * Ports have host ports and internal ports currently but they
     * really only should have internal ports as host ports should be assigned
     * by the system. Host ports might go away in the future.
     * @type {Array<RegistrySvcPortMapping>}
     * @memberof RegistrySvcDefinition
     */
    ports?: Array<RegistrySvcPortMapping>;
    /**
     * Repository based definitions is an alternative to Image definitions.
     * Instead of deploying an already built and pushed image, a source code repository
     * url can be provided. The container will be built from the source.
     * @type {RegistrySvcRepositorySpec}
     * @memberof RegistrySvcDefinition
     */
    repository?: RegistrySvcRepositorySpec;
}
/**
 * Check if a given object implements the RegistrySvcDefinition interface.
 */
export declare function instanceOfRegistrySvcDefinition(value: object): value is RegistrySvcDefinition;
export declare function RegistrySvcDefinitionFromJSON(json: any): RegistrySvcDefinition;
export declare function RegistrySvcDefinitionFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcDefinition;
export declare function RegistrySvcDefinitionToJSON(json: any): RegistrySvcDefinition;
export declare function RegistrySvcDefinitionToJSONTyped(value?: RegistrySvcDefinition | null, ignoreDiscriminator?: boolean): any;
