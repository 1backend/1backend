/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ChatSvcMessage
 */
export interface ChatSvcMessage {
    /**
     *
     * @type {string}
     * @memberof ChatSvcMessage
     */
    createdAt?: string;
    /**
     * FileIds defines the file attachments the message has.
     * @type {Array<string>}
     * @memberof ChatSvcMessage
     */
    fileIds?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof ChatSvcMessage
     */
    id: string;
    /**
     *
     * @type {{ [key: string]: any; }}
     * @memberof ChatSvcMessage
     */
    meta?: {
        [key: string]: any;
    };
    /**
     * Text content of the message eg. "Hi, what's up?"
     * @type {string}
     * @memberof ChatSvcMessage
     */
    text?: string;
    /**
     * ThreadId of the message.
     * @type {string}
     * @memberof ChatSvcMessage
     */
    threadId: string;
    /**
     *
     * @type {string}
     * @memberof ChatSvcMessage
     */
    updatedAt?: string;
    /**
     * UserId is the id of the user who wrote the message.
     * For AI messages this field is empty.
     * @type {string}
     * @memberof ChatSvcMessage
     */
    userId?: string;
}
/**
 * Check if a given object implements the ChatSvcMessage interface.
 */
export declare function instanceOfChatSvcMessage(value: object): value is ChatSvcMessage;
export declare function ChatSvcMessageFromJSON(json: any): ChatSvcMessage;
export declare function ChatSvcMessageFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcMessage;
export declare function ChatSvcMessageToJSON(json: any): ChatSvcMessage;
export declare function ChatSvcMessageToJSONTyped(value?: ChatSvcMessage | null, ignoreDiscriminator?: boolean): any;
