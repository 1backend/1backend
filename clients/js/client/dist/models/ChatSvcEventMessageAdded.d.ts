/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ChatSvcEventMessageAdded
 */
export interface ChatSvcEventMessageAdded {
    /**
     *
     * @type {string}
     * @memberof ChatSvcEventMessageAdded
     */
    threadId?: string;
}
/**
 * Check if a given object implements the ChatSvcEventMessageAdded interface.
 */
export declare function instanceOfChatSvcEventMessageAdded(value: object): value is ChatSvcEventMessageAdded;
export declare function ChatSvcEventMessageAddedFromJSON(json: any): ChatSvcEventMessageAdded;
export declare function ChatSvcEventMessageAddedFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcEventMessageAdded;
export declare function ChatSvcEventMessageAddedToJSON(json: any): ChatSvcEventMessageAdded;
export declare function ChatSvcEventMessageAddedToJSONTyped(value?: ChatSvcEventMessageAdded | null, ignoreDiscriminator?: boolean): any;
