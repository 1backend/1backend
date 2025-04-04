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
 * @interface FirehoseSvcEvent
 */
export interface FirehoseSvcEvent {
    /**
     *
     * @type {object}
     * @memberof FirehoseSvcEvent
     */
    data?: object;
    /**
     *
     * @type {string}
     * @memberof FirehoseSvcEvent
     */
    name?: string;
}
/**
 * Check if a given object implements the FirehoseSvcEvent interface.
 */
export declare function instanceOfFirehoseSvcEvent(value: object): value is FirehoseSvcEvent;
export declare function FirehoseSvcEventFromJSON(json: any): FirehoseSvcEvent;
export declare function FirehoseSvcEventFromJSONTyped(json: any, ignoreDiscriminator: boolean): FirehoseSvcEvent;
export declare function FirehoseSvcEventToJSON(json: any): FirehoseSvcEvent;
export declare function FirehoseSvcEventToJSONTyped(value?: FirehoseSvcEvent | null, ignoreDiscriminator?: boolean): any;
