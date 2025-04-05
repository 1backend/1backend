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
import type { ChatSvcThread } from './ChatSvcThread';
/**
 *
 * @export
 * @interface ChatSvcGetThreadsResponse
 */
export interface ChatSvcGetThreadsResponse {
    /**
     *
     * @type {Array<ChatSvcThread>}
     * @memberof ChatSvcGetThreadsResponse
     */
    threads?: Array<ChatSvcThread>;
}
/**
 * Check if a given object implements the ChatSvcGetThreadsResponse interface.
 */
export declare function instanceOfChatSvcGetThreadsResponse(value: object): value is ChatSvcGetThreadsResponse;
export declare function ChatSvcGetThreadsResponseFromJSON(json: any): ChatSvcGetThreadsResponse;
export declare function ChatSvcGetThreadsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcGetThreadsResponse;
export declare function ChatSvcGetThreadsResponseToJSON(json: any): ChatSvcGetThreadsResponse;
export declare function ChatSvcGetThreadsResponseToJSONTyped(value?: ChatSvcGetThreadsResponse | null, ignoreDiscriminator?: boolean): any;
