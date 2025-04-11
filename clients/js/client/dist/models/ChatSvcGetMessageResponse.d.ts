/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
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
 * @interface ChatSvcGetMessageResponse
 */
export interface ChatSvcGetMessageResponse {
    /**
     *
     * @type {boolean}
     * @memberof ChatSvcGetMessageResponse
     */
    _exists?: boolean;
    /**
     *
     * @type {ChatSvcMessage}
     * @memberof ChatSvcGetMessageResponse
     */
    message?: ChatSvcMessage;
}
/**
 * Check if a given object implements the ChatSvcGetMessageResponse interface.
 */
export declare function instanceOfChatSvcGetMessageResponse(value: object): value is ChatSvcGetMessageResponse;
export declare function ChatSvcGetMessageResponseFromJSON(json: any): ChatSvcGetMessageResponse;
export declare function ChatSvcGetMessageResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcGetMessageResponse;
export declare function ChatSvcGetMessageResponseToJSON(json: any): ChatSvcGetMessageResponse;
export declare function ChatSvcGetMessageResponseToJSONTyped(value?: ChatSvcGetMessageResponse | null, ignoreDiscriminator?: boolean): any;
