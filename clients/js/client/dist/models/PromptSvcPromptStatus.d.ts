/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
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
export declare const PromptSvcPromptStatus: {
    readonly PromptStatusScheduled: "scheduled";
    readonly PromptStatusRunning: "running";
    readonly PromptStatusCompleted: "completed";
    readonly PromptStatusErrored: "errored";
    readonly PromptStatusAbandoned: "abandoned";
    readonly PromptStatusCanceled: "canceled";
};
export type PromptSvcPromptStatus = typeof PromptSvcPromptStatus[keyof typeof PromptSvcPromptStatus];
export declare function instanceOfPromptSvcPromptStatus(value: any): boolean;
export declare function PromptSvcPromptStatusFromJSON(json: any): PromptSvcPromptStatus;
export declare function PromptSvcPromptStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcPromptStatus;
export declare function PromptSvcPromptStatusToJSON(value?: PromptSvcPromptStatus | null): any;
export declare function PromptSvcPromptStatusToJSONTyped(value: any, ignoreDiscriminator: boolean): PromptSvcPromptStatus;
