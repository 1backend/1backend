/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ChatSvcEventThreadAdded
 */
export interface ChatSvcEventThreadAdded {
    /**
     *
     * @type {string}
     * @memberof ChatSvcEventThreadAdded
     */
    threadId?: string;
}
/**
 * Check if a given object implements the ChatSvcEventThreadAdded interface.
 */
export declare function instanceOfChatSvcEventThreadAdded(value: object): value is ChatSvcEventThreadAdded;
export declare function ChatSvcEventThreadAddedFromJSON(json: any): ChatSvcEventThreadAdded;
export declare function ChatSvcEventThreadAddedFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcEventThreadAdded;
export declare function ChatSvcEventThreadAddedToJSON(json: any): ChatSvcEventThreadAdded;
export declare function ChatSvcEventThreadAddedToJSONTyped(value?: ChatSvcEventThreadAdded | null, ignoreDiscriminator?: boolean): any;
