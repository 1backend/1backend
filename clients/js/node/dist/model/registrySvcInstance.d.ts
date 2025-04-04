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
import { RegistrySvcInstanceStatus } from './registrySvcInstanceStatus';
export declare class RegistrySvcInstance {
    /**
    * The ID of the deployment that this instance is an instance of. Only instances deployed by 1Backend have a DeploymentId. Services can be deployed through other means (Docker Compose, K8s, anything), in that case they self-register and will not have a DeploymentId.
    */
    'deploymentId'?: string;
    /**
    * Details
    */
    'details'?: string;
    /**
    * Host of the instance address. Required if URL is not provided
    */
    'host'?: string;
    /**
    * Required: ID of the instance
    */
    'id': string;
    /**
    * IP of the instance address. Optional: to register by IP instead of host
    */
    'ip'?: string;
    /**
    * Last time the instance gave a sign of life
    */
    'lastHeartbeat'?: string;
    /**
    * NodeURL is the URL of the 1Backend server the instance is running on. To have a NodeURL the instance must either: - Be deployed by 1Backend - Declare the 1Backend server URL when registering its instance
    */
    'nodeUrl'?: string;
    /**
    * Path of the instance address. Optional (e.g., \"/api\")
    */
    'path'?: string;
    /**
    * Port of the instance address. Required if URL is not provided
    */
    'port'?: number;
    /**
    * Scheme of the instance address. Required if URL is not provided.
    */
    'scheme'?: string;
    /**
    * Slug of the account that owns this instance Services that want to be proxied by their slug are advised to self register their instance at startup. Keep in mind, instances might be deployed by 1Backend yet they still won\'t be 1Backend services and they won\'t have slugs. Think NGINX, MySQL, etc.
    */
    'slug'?: string;
    /**
    * Status
    */
    'status': RegistrySvcInstanceStatus;
    /**
    * Tags are used to filter instances
    */
    'tags'?: Array<string>;
    /**
    * Full address URL of the instance.
    */
    'url': string;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
export declare namespace RegistrySvcInstance {
}
