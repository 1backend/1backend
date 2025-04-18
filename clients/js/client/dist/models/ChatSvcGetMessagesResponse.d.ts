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
import type { ChatSvcMessage } from './ChatSvcMessage';
/**
 *
 * @export
 * @interface ChatSvcGetMessagesResponse
 */
export interface ChatSvcGetMessagesResponse {
    /**
     *
     * @type {Array<ChatSvcMessage>}
     * @memberof ChatSvcGetMessagesResponse
     */
    messages?: Array<ChatSvcMessage>;
}
/**
 * Check if a given object implements the ChatSvcGetMessagesResponse interface.
 */
export declare function instanceOfChatSvcGetMessagesResponse(value: object): value is ChatSvcGetMessagesResponse;
export declare function ChatSvcGetMessagesResponseFromJSON(json: any): ChatSvcGetMessagesResponse;
export declare function ChatSvcGetMessagesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcGetMessagesResponse;
export declare function ChatSvcGetMessagesResponseToJSON(json: any): ChatSvcGetMessagesResponse;
export declare function ChatSvcGetMessagesResponseToJSONTyped(value?: ChatSvcGetMessagesResponse | null, ignoreDiscriminator?: boolean): any;
