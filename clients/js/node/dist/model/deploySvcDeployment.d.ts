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
import { DeploySvcAutoScalingConfig } from './deploySvcAutoScalingConfig';
import { DeploySvcDeploymentStatus } from './deploySvcDeploymentStatus';
import { DeploySvcDeploymentStrategy } from './deploySvcDeploymentStrategy';
import { DeploySvcResourceLimits } from './deploySvcResourceLimits';
import { DeploySvcTargetRegion } from './deploySvcTargetRegion';
export declare class DeploySvcDeployment {
    /**
    * Optional: Auto-scaling rules
    */
    'autoScaling'?: DeploySvcAutoScalingConfig;
    /**
    * DefinitionId is the id of the definition
    */
    'definitionId': string;
    /**
    * Description of what this deployment does
    */
    'description'?: string;
    /**
    * Details provides additional information about the deployment\'s current state, including both success and failure conditions (e.g., \"Deployment in progress\", \"Error pulling image\").
    */
    'details'?: string;
    /**
    * Envars is a map of environment variables that will be passed down to service instances (see Registry Svc Instance) Also see the Registry Svc Definition for required envars.
    */
    'envars'?: {
        [key: string]: string;
    };
    /**
    * ID of the deployment (e.g., \"depl_dbOdi5eLQK\")
    */
    'id': string;
    /**
    * Short name for easy reference (e.g., \"user-service-v2\")
    */
    'name'?: string;
    /**
    * Number of container instances to run
    */
    'replicas'?: number;
    /**
    * Resource requirements for each replica
    */
    'resources'?: DeploySvcResourceLimits;
    /**
    * Current status of the deployment (e.g., \"OK\", \"Error\", \"Pending\")
    */
    'status'?: DeploySvcDeploymentStatus;
    /**
    * Deployment strategy (e.g., rolling update)
    */
    'strategy'?: DeploySvcDeploymentStrategy;
    /**
    * Target deployment regions or clusters
    */
    'targetRegions'?: Array<DeploySvcTargetRegion>;
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
export declare namespace DeploySvcDeployment {
}
