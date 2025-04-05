/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
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
 * @interface ChatSvcAddThreadResponse
 */
export interface ChatSvcAddThreadResponse {
    /**
     *
     * @type {ChatSvcThread}
     * @memberof ChatSvcAddThreadResponse
     */
    thread?: ChatSvcThread;
}
/**
 * Check if a given object implements the ChatSvcAddThreadResponse interface.
 */
export declare function instanceOfChatSvcAddThreadResponse(value: object): value is ChatSvcAddThreadResponse;
export declare function ChatSvcAddThreadResponseFromJSON(json: any): ChatSvcAddThreadResponse;
export declare function ChatSvcAddThreadResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcAddThreadResponse;
export declare function ChatSvcAddThreadResponseToJSON(json: any): ChatSvcAddThreadResponse;
export declare function ChatSvcAddThreadResponseToJSONTyped(value?: ChatSvcAddThreadResponse | null, ignoreDiscriminator?: boolean): any;
