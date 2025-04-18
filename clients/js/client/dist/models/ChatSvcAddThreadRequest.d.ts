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
import type { ChatSvcThread } from './ChatSvcThread';
/**
 *
 * @export
 * @interface ChatSvcAddThreadRequest
 */
export interface ChatSvcAddThreadRequest {
    /**
     *
     * @type {ChatSvcThread}
     * @memberof ChatSvcAddThreadRequest
     */
    thread?: ChatSvcThread;
}
/**
 * Check if a given object implements the ChatSvcAddThreadRequest interface.
 */
export declare function instanceOfChatSvcAddThreadRequest(value: object): value is ChatSvcAddThreadRequest;
export declare function ChatSvcAddThreadRequestFromJSON(json: any): ChatSvcAddThreadRequest;
export declare function ChatSvcAddThreadRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcAddThreadRequest;
export declare function ChatSvcAddThreadRequestToJSON(json: any): ChatSvcAddThreadRequest;
export declare function ChatSvcAddThreadRequestToJSONTyped(value?: ChatSvcAddThreadRequest | null, ignoreDiscriminator?: boolean): any;
