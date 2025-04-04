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
import type { ContainerSvcKeep } from './ContainerSvcKeep';
import type { ContainerSvcAsset } from './ContainerSvcAsset';
import type { ContainerSvcNetwork } from './ContainerSvcNetwork';
import type { ContainerSvcCapabilities } from './ContainerSvcCapabilities';
import type { ContainerSvcVolume } from './ContainerSvcVolume';
import type { ContainerSvcLabel } from './ContainerSvcLabel';
import type { ContainerSvcEnvVar } from './ContainerSvcEnvVar';
import type { ContainerSvcResources } from './ContainerSvcResources';
import type { ContainerSvcPortMapping } from './ContainerSvcPortMapping';
/**
 *
 * @export
 * @interface ContainerSvcContainer
 */
export interface ContainerSvcContainer {
    /**
     * Assets maps environment variable names to file URLs.
     * Example: {"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"}
     * These files are downloaded by the File Svc and mounted in the container.
     * The environment variable `MODEL` will point to the local file path in the container.
     * @type {Array<ContainerSvcAsset>}
     * @memberof ContainerSvcContainer
     */
    assets?: Array<ContainerSvcAsset>;
    /**
     * Capabilities define additional runtime features, such as GPU support.
     * @type {ContainerSvcCapabilities}
     * @memberof ContainerSvcContainer
     */
    capabilities?: ContainerSvcCapabilities;
    /**
     * Envs are environment variables set within the container.
     * @type {Array<ContainerSvcEnvVar>}
     * @memberof ContainerSvcContainer
     */
    envs?: Array<ContainerSvcEnvVar>;
    /**
     * Hash is a unique identifier associated with the container.
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    hash?: string;
    /**
     * Id is the unique identifier for the container instance.
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    id?: string;
    /**
     * Image is the Docker image used to create the container.
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    image?: string;
    /**
     * Keeps are paths that persist across container restarts.
     * They function like mounts or volumes, but their external storage location is irrelevant.
     * @type {Array<ContainerSvcKeep>}
     * @memberof ContainerSvcContainer
     */
    keeps?: Array<ContainerSvcKeep>;
    /**
     * Labels are metadata tags assigned to the container.
     * @type {Array<ContainerSvcLabel>}
     * @memberof ContainerSvcContainer
     */
    labels?: Array<ContainerSvcLabel>;
    /**
     * Names are the human-readable aliases assigned to the container.
     * @type {Array<string>}
     * @memberof ContainerSvcContainer
     */
    names?: Array<string>;
    /**
     * Network contains networking-related information for the container.
     * @type {ContainerSvcNetwork}
     * @memberof ContainerSvcContainer
     */
    network?: ContainerSvcNetwork;
    /**
     * Node Id
     * Please see the documentation for the envar OB_NODE_ID
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    nodeId?: string;
    /**
     * Ports maps host ports (keys) to container ports (values).
     * @type {Array<ContainerSvcPortMapping>}
     * @memberof ContainerSvcContainer
     */
    ports?: Array<ContainerSvcPortMapping>;
    /**
     * Resources defines CPU, memory, and disk constraints for the container.
     * @type {ContainerSvcResources}
     * @memberof ContainerSvcContainer
     */
    resources?: ContainerSvcResources;
    /**
     * Runtime specifies the container runtime (e.g., Docker, containerd, etc.).
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    runtime?: string;
    /**
     * Status indicates the current state of the container (e.g., running, stopped).
     * @type {string}
     * @memberof ContainerSvcContainer
     */
    status?: string;
    /**
     * Volumes mounted by the container.
     * @type {Array<ContainerSvcVolume>}
     * @memberof ContainerSvcContainer
     */
    volumes?: Array<ContainerSvcVolume>;
}
/**
 * Check if a given object implements the ContainerSvcContainer interface.
 */
export declare function instanceOfContainerSvcContainer(value: object): value is ContainerSvcContainer;
export declare function ContainerSvcContainerFromJSON(json: any): ContainerSvcContainer;
export declare function ContainerSvcContainerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcContainer;
export declare function ContainerSvcContainerToJSON(json: any): ContainerSvcContainer;
export declare function ContainerSvcContainerToJSONTyped(value?: ContainerSvcContainer | null, ignoreDiscriminator?: boolean): any;
