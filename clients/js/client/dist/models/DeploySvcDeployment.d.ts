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
import type { DeploySvcDeploymentStrategy } from './DeploySvcDeploymentStrategy';
import type { DeploySvcAutoScalingConfig } from './DeploySvcAutoScalingConfig';
import type { DeploySvcDeploymentStatus } from './DeploySvcDeploymentStatus';
import type { DeploySvcTargetRegion } from './DeploySvcTargetRegion';
import type { DeploySvcResourceLimits } from './DeploySvcResourceLimits';
/**
 *
 * @export
 * @interface DeploySvcDeployment
 */
export interface DeploySvcDeployment {
    /**
     * Optional: Auto-scaling rules
     * @type {DeploySvcAutoScalingConfig}
     * @memberof DeploySvcDeployment
     */
    autoScaling?: DeploySvcAutoScalingConfig;
    /**
     * DefinitionId is the id of the definition
     * @type {string}
     * @memberof DeploySvcDeployment
     */
    definitionId: string;
    /**
     * Description of what this deployment does
     * @type {string}
     * @memberof DeploySvcDeployment
     */
    description?: string;
    /**
     * Details provides additional information about the deployment's current state,
     * including both success and failure conditions (e.g., "Deployment in progress", "Error pulling image").
     * @type {string}
     * @memberof DeploySvcDeployment
     */
    details?: string;
    /**
     * Envars is a map of environment variables that will be passed down to service instances (see Registry Svc Instance)
     * Also see the Registry Svc Definition for required envars.
     * @type {{ [key: string]: string; }}
     * @memberof DeploySvcDeployment
     */
    envars?: {
        [key: string]: string;
    };
    /**
     * ID of the deployment (e.g., "depl_dbOdi5eLQK")
     * @type {string}
     * @memberof DeploySvcDeployment
     */
    id: string;
    /**
     * Short name for easy reference (e.g., "user-service-v2")
     * @type {string}
     * @memberof DeploySvcDeployment
     */
    name?: string;
    /**
     * Number of container instances to run
     * @type {number}
     * @memberof DeploySvcDeployment
     */
    replicas?: number;
    /**
     * Resource requirements for each replica
     * @type {DeploySvcResourceLimits}
     * @memberof DeploySvcDeployment
     */
    resources?: DeploySvcResourceLimits;
    /**
     * Current status of the deployment (e.g., "OK", "Error", "Pending")
     * @type {DeploySvcDeploymentStatus}
     * @memberof DeploySvcDeployment
     */
    status?: DeploySvcDeploymentStatus;
    /**
     * Deployment strategy (e.g., rolling update)
     * @type {DeploySvcDeploymentStrategy}
     * @memberof DeploySvcDeployment
     */
    strategy?: DeploySvcDeploymentStrategy;
    /**
     * Target deployment regions or clusters
     * @type {Array<DeploySvcTargetRegion>}
     * @memberof DeploySvcDeployment
     */
    targetRegions?: Array<DeploySvcTargetRegion>;
}
/**
 * Check if a given object implements the DeploySvcDeployment interface.
 */
export declare function instanceOfDeploySvcDeployment(value: object): value is DeploySvcDeployment;
export declare function DeploySvcDeploymentFromJSON(json: any): DeploySvcDeployment;
export declare function DeploySvcDeploymentFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcDeployment;
export declare function DeploySvcDeploymentToJSON(json: any): DeploySvcDeployment;
export declare function DeploySvcDeploymentToJSONTyped(value?: DeploySvcDeployment | null, ignoreDiscriminator?: boolean): any;
