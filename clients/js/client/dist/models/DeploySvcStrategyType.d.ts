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
export declare const DeploySvcStrategyType: {
    readonly StrategyRollingUpdate: "RollingUpdate";
    readonly StrategyRecreate: "Recreate";
};
export type DeploySvcStrategyType = typeof DeploySvcStrategyType[keyof typeof DeploySvcStrategyType];
export declare function instanceOfDeploySvcStrategyType(value: any): boolean;
export declare function DeploySvcStrategyTypeFromJSON(json: any): DeploySvcStrategyType;
export declare function DeploySvcStrategyTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcStrategyType;
export declare function DeploySvcStrategyTypeToJSON(value?: DeploySvcStrategyType | null): any;
export declare function DeploySvcStrategyTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): DeploySvcStrategyType;
