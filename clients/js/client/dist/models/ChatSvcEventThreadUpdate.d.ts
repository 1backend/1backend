/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ChatSvcEventThreadUpdate
 */
export interface ChatSvcEventThreadUpdate {
    /**
     *
     * @type {string}
     * @memberof ChatSvcEventThreadUpdate
     */
    threadId?: string;
}
/**
 * Check if a given object implements the ChatSvcEventThreadUpdate interface.
 */
export declare function instanceOfChatSvcEventThreadUpdate(value: object): value is ChatSvcEventThreadUpdate;
export declare function ChatSvcEventThreadUpdateFromJSON(json: any): ChatSvcEventThreadUpdate;
export declare function ChatSvcEventThreadUpdateFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcEventThreadUpdate;
export declare function ChatSvcEventThreadUpdateToJSON(json: any): ChatSvcEventThreadUpdate;
export declare function ChatSvcEventThreadUpdateToJSONTyped(value?: ChatSvcEventThreadUpdate | null, ignoreDiscriminator?: boolean): any;
