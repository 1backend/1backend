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
import type { DeploySvcStrategyType } from './DeploySvcStrategyType';
/**
 *
 * @export
 * @interface DeploySvcDeploymentStrategy
 */
export interface DeploySvcDeploymentStrategy {
    /**
     * Max extra replicas during update
     * @type {number}
     * @memberof DeploySvcDeploymentStrategy
     */
    maxSurge?: number;
    /**
     * Max unavailable replicas during update
     * @type {number}
     * @memberof DeploySvcDeploymentStrategy
     */
    maxUnavailable?: number;
    /**
     * Deployment strategy type (RollingUpdate, Recreate, etc.)
     * @type {DeploySvcStrategyType}
     * @memberof DeploySvcDeploymentStrategy
     */
    type?: DeploySvcStrategyType;
}
/**
 * Check if a given object implements the DeploySvcDeploymentStrategy interface.
 */
export declare function instanceOfDeploySvcDeploymentStrategy(value: object): value is DeploySvcDeploymentStrategy;
export declare function DeploySvcDeploymentStrategyFromJSON(json: any): DeploySvcDeploymentStrategy;
export declare function DeploySvcDeploymentStrategyFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcDeploymentStrategy;
export declare function DeploySvcDeploymentStrategyToJSON(json: any): DeploySvcDeploymentStrategy;
export declare function DeploySvcDeploymentStrategyToJSONTyped(value?: DeploySvcDeploymentStrategy | null, ignoreDiscriminator?: boolean): any;
