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
 * @interface ChatSvcAddMessageRequest
 */
export interface ChatSvcAddMessageRequest {
    /**
     *
     * @type {ChatSvcMessage}
     * @memberof ChatSvcAddMessageRequest
     */
    message?: ChatSvcMessage;
}
/**
 * Check if a given object implements the ChatSvcAddMessageRequest interface.
 */
export declare function instanceOfChatSvcAddMessageRequest(value: object): value is ChatSvcAddMessageRequest;
export declare function ChatSvcAddMessageRequestFromJSON(json: any): ChatSvcAddMessageRequest;
export declare function ChatSvcAddMessageRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcAddMessageRequest;
export declare function ChatSvcAddMessageRequestToJSON(json: any): ChatSvcAddMessageRequest;
export declare function ChatSvcAddMessageRequestToJSONTyped(value?: ChatSvcAddMessageRequest | null, ignoreDiscriminator?: boolean): any;
