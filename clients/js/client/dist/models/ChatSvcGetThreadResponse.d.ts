/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
 * @interface ChatSvcGetThreadResponse
 */
export interface ChatSvcGetThreadResponse {
    /**
     *
     * @type {boolean}
     * @memberof ChatSvcGetThreadResponse
     */
    _exists?: boolean;
    /**
     *
     * @type {ChatSvcThread}
     * @memberof ChatSvcGetThreadResponse
     */
    thread?: ChatSvcThread;
}
/**
 * Check if a given object implements the ChatSvcGetThreadResponse interface.
 */
export declare function instanceOfChatSvcGetThreadResponse(value: object): value is ChatSvcGetThreadResponse;
export declare function ChatSvcGetThreadResponseFromJSON(json: any): ChatSvcGetThreadResponse;
export declare function ChatSvcGetThreadResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcGetThreadResponse;
export declare function ChatSvcGetThreadResponseToJSON(json: any): ChatSvcGetThreadResponse;
export declare function ChatSvcGetThreadResponseToJSONTyped(value?: ChatSvcGetThreadResponse | null, ignoreDiscriminator?: boolean): any;
