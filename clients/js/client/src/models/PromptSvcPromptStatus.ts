/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
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
export const PromptSvcPromptStatus = {
    PromptStatusScheduled: 'scheduled',
    PromptStatusRunning: 'running',
    PromptStatusCompleted: 'completed',
    PromptStatusErrored: 'errored',
    PromptStatusAbandoned: 'abandoned',
    PromptStatusCanceled: 'canceled'
} as const;
export type PromptSvcPromptStatus = typeof PromptSvcPromptStatus[keyof typeof PromptSvcPromptStatus];


export function instanceOfPromptSvcPromptStatus(value: any): boolean {
    for (const key in PromptSvcPromptStatus) {
        if (Object.prototype.hasOwnProperty.call(PromptSvcPromptStatus, key)) {
            if (PromptSvcPromptStatus[key as keyof typeof PromptSvcPromptStatus] === value) {
                return true;
            }
        }
    }
    return false;
}

export function PromptSvcPromptStatusFromJSON(json: any): PromptSvcPromptStatus {
    return PromptSvcPromptStatusFromJSONTyped(json, false);
}

export function PromptSvcPromptStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcPromptStatus {
    return json as PromptSvcPromptStatus;
}

export function PromptSvcPromptStatusToJSON(value?: PromptSvcPromptStatus | null): any {
    return value as any;
}

export function PromptSvcPromptStatusToJSONTyped(value: any, ignoreDiscriminator: boolean): PromptSvcPromptStatus {
    return value as PromptSvcPromptStatus;
}

