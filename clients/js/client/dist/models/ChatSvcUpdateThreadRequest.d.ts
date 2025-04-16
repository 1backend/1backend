/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
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
 * @interface ChatSvcUpdateThreadRequest
 */
export interface ChatSvcUpdateThreadRequest {
    /**
     *
     * @type {ChatSvcThread}
     * @memberof ChatSvcUpdateThreadRequest
     */
    thread?: ChatSvcThread;
}
/**
 * Check if a given object implements the ChatSvcUpdateThreadRequest interface.
 */
export declare function instanceOfChatSvcUpdateThreadRequest(value: object): value is ChatSvcUpdateThreadRequest;
export declare function ChatSvcUpdateThreadRequestFromJSON(json: any): ChatSvcUpdateThreadRequest;
export declare function ChatSvcUpdateThreadRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcUpdateThreadRequest;
export declare function ChatSvcUpdateThreadRequestToJSON(json: any): ChatSvcUpdateThreadRequest;
export declare function ChatSvcUpdateThreadRequestToJSONTyped(value?: ChatSvcUpdateThreadRequest | null, ignoreDiscriminator?: boolean): any;
