/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
export const RegistrySvcInstanceStatus = {
    InstanceStatusUnknown: 'Unknown',
    InstanceStatusHealthy: 'Healthy',
    InstanceStatusDegraded: 'Degraded',
    InstanceStatusUnreachable: 'Unreachable',
    InstanceStatusError: 'Error'
} as const;
export type RegistrySvcInstanceStatus = typeof RegistrySvcInstanceStatus[keyof typeof RegistrySvcInstanceStatus];


export function instanceOfRegistrySvcInstanceStatus(value: any): boolean {
    for (const key in RegistrySvcInstanceStatus) {
        if (Object.prototype.hasOwnProperty.call(RegistrySvcInstanceStatus, key)) {
            if (RegistrySvcInstanceStatus[key as keyof typeof RegistrySvcInstanceStatus] === value) {
                return true;
            }
        }
    }
    return false;
}

export function RegistrySvcInstanceStatusFromJSON(json: any): RegistrySvcInstanceStatus {
    return RegistrySvcInstanceStatusFromJSONTyped(json, false);
}

export function RegistrySvcInstanceStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcInstanceStatus {
    return json as RegistrySvcInstanceStatus;
}

export function RegistrySvcInstanceStatusToJSON(value?: RegistrySvcInstanceStatus | null): any {
    return value as any;
}

export function RegistrySvcInstanceStatusToJSONTyped(value: any, ignoreDiscriminator: boolean): RegistrySvcInstanceStatus {
    return value as RegistrySvcInstanceStatus;
}

