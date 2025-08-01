/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 */
export declare const DeploySvcDeploymentStatus: {
    readonly DeploymentStatusOK: "OK";
    readonly DeploymentStatusError: "Error";
    readonly DeploymentStatusPending: "Pending";
    readonly DeploymentStatusFailed: "Failed";
    readonly DeploymentStatusDeploying: "Deploying";
};
export type DeploySvcDeploymentStatus = typeof DeploySvcDeploymentStatus[keyof typeof DeploySvcDeploymentStatus];
export declare function instanceOfDeploySvcDeploymentStatus(value: any): boolean;
export declare function DeploySvcDeploymentStatusFromJSON(json: any): DeploySvcDeploymentStatus;
export declare function DeploySvcDeploymentStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcDeploymentStatus;
export declare function DeploySvcDeploymentStatusToJSON(value?: DeploySvcDeploymentStatus | null): any;
export declare function DeploySvcDeploymentStatusToJSONTyped(value: any, ignoreDiscriminator: boolean): DeploySvcDeploymentStatus;
