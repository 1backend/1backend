/* tslint:disable */
/* eslint-disable */
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


/**
 * 
 * @export
 */
export const DeploySvcDeploymentStatus = {
    DeploymentStatusOK: 'OK',
    DeploymentStatusError: 'Error',
    DeploymentStatusPending: 'Pending',
    DeploymentStatusFailed: 'Failed',
    DeploymentStatusDeploying: 'Deploying'
} as const;
export type DeploySvcDeploymentStatus = typeof DeploySvcDeploymentStatus[keyof typeof DeploySvcDeploymentStatus];


export function instanceOfDeploySvcDeploymentStatus(value: any): boolean {
    for (const key in DeploySvcDeploymentStatus) {
        if (Object.prototype.hasOwnProperty.call(DeploySvcDeploymentStatus, key)) {
            if (DeploySvcDeploymentStatus[key as keyof typeof DeploySvcDeploymentStatus] === value) {
                return true;
            }
        }
    }
    return false;
}

export function DeploySvcDeploymentStatusFromJSON(json: any): DeploySvcDeploymentStatus {
    return DeploySvcDeploymentStatusFromJSONTyped(json, false);
}

export function DeploySvcDeploymentStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcDeploymentStatus {
    return json as DeploySvcDeploymentStatus;
}

export function DeploySvcDeploymentStatusToJSON(value?: DeploySvcDeploymentStatus | null): any {
    return value as any;
}

export function DeploySvcDeploymentStatusToJSONTyped(value: any, ignoreDiscriminator: boolean): DeploySvcDeploymentStatus {
    return value as DeploySvcDeploymentStatus;
}

