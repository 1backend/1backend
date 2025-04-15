/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ContainerSvcKeep } from './ContainerSvcKeep';
import type { ContainerSvcAsset } from './ContainerSvcAsset';
import type { ContainerSvcCapabilities } from './ContainerSvcCapabilities';
import type { ContainerSvcLabel } from './ContainerSvcLabel';
import type { ContainerSvcEnvVar } from './ContainerSvcEnvVar';
import type { ContainerSvcPortMapping } from './ContainerSvcPortMapping';
/**
 *
 * @export
 * @interface ContainerSvcRunContainerRequest
 */
export interface ContainerSvcRunContainerRequest {
    /**
     * Assets maps environment variable names to file URLs.
     * Example: {"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"}
     * These files are downloaded by the File Svc and mounted in the container.
     * The environment variable `MODEL` will point to the local file path in the container.
     * @type {Array<ContainerSvcAsset>}
     * @memberof ContainerSvcRunContainerRequest
     */
    assets?: Array<ContainerSvcAsset>;
    /**
     * Capabilities define additional runtime features, such as GPU support.
     * @type {ContainerSvcCapabilities}
     * @memberof ContainerSvcRunContainerRequest
     */
    capabilities?: ContainerSvcCapabilities;
    /**
     * Envs are environment variables set within the container.
     * @type {Array<ContainerSvcEnvVar>}
     * @memberof ContainerSvcRunContainerRequest
     */
    envs?: Array<ContainerSvcEnvVar>;
    /**
     * Hash is a unique identifier for the container
     * @type {string}
     * @memberof ContainerSvcRunContainerRequest
     */
    hash?: string;
    /**
     * Image is the Docker image to use for the container
     * @type {string}
     * @memberof ContainerSvcRunContainerRequest
     */
    image: string;
    /**
     * Keeps are paths that persist across container restarts.
     * They function like mounts or volumes, but their external storage location is irrelevant.
     * @type {Array<ContainerSvcKeep>}
     * @memberof ContainerSvcRunContainerRequest
     */
    keeps?: Array<ContainerSvcKeep>;
    /**
     * Labels are metadata tags assigned to the container.
     * @type {Array<ContainerSvcLabel>}
     * @memberof ContainerSvcRunContainerRequest
     */
    labels?: Array<ContainerSvcLabel>;
    /**
     * Names are the human-readable aliases assigned to the container.
     * @type {Array<string>}
     * @memberof ContainerSvcRunContainerRequest
     */
    names?: Array<string>;
    /**
     * Ports maps host ports (keys) to container ports (values).
     * @type {Array<ContainerSvcPortMapping>}
     * @memberof ContainerSvcRunContainerRequest
     */
    ports?: Array<ContainerSvcPortMapping>;
}
/**
 * Check if a given object implements the ContainerSvcRunContainerRequest interface.
 */
export declare function instanceOfContainerSvcRunContainerRequest(value: object): value is ContainerSvcRunContainerRequest;
export declare function ContainerSvcRunContainerRequestFromJSON(json: any): ContainerSvcRunContainerRequest;
export declare function ContainerSvcRunContainerRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcRunContainerRequest;
export declare function ContainerSvcRunContainerRequestToJSON(json: any): ContainerSvcRunContainerRequest;
export declare function ContainerSvcRunContainerRequestToJSONTyped(value?: ContainerSvcRunContainerRequest | null, ignoreDiscriminator?: boolean): any;
