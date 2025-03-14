/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
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
export const DeploySvcStrategyType = {
    StrategyRollingUpdate: 'RollingUpdate',
    StrategyRecreate: 'Recreate'
} as const;
export type DeploySvcStrategyType = typeof DeploySvcStrategyType[keyof typeof DeploySvcStrategyType];


export function instanceOfDeploySvcStrategyType(value: any): boolean {
    for (const key in DeploySvcStrategyType) {
        if (Object.prototype.hasOwnProperty.call(DeploySvcStrategyType, key)) {
            if (DeploySvcStrategyType[key as keyof typeof DeploySvcStrategyType] === value) {
                return true;
            }
        }
    }
    return false;
}

export function DeploySvcStrategyTypeFromJSON(json: any): DeploySvcStrategyType {
    return DeploySvcStrategyTypeFromJSONTyped(json, false);
}

export function DeploySvcStrategyTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcStrategyType {
    return json as DeploySvcStrategyType;
}

export function DeploySvcStrategyTypeToJSON(value?: DeploySvcStrategyType | null): any {
    return value as any;
}

export function DeploySvcStrategyTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): DeploySvcStrategyType {
    return value as DeploySvcStrategyType;
}

