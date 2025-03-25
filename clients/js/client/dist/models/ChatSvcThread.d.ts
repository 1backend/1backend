/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface ChatSvcThread
 */
export interface ChatSvcThread {
    /**
     *
     * @type {string}
     * @memberof ChatSvcThread
     */
    createdAt?: string;
    /**
     *
     * @type {string}
     * @memberof ChatSvcThread
     */
    id: string;
    /**
     * Title of the thread.
     * @type {string}
     * @memberof ChatSvcThread
     */
    title?: string;
    /**
     * TopicIds defines which topics the thread belongs to.
     * Topics can roughly be thought of as tags for threads.
     * @type {Array<string>}
     * @memberof ChatSvcThread
     */
    topicIds?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof ChatSvcThread
     */
    updatedAt?: string;
    /**
     * UserIds the ids of the users who can see this thread.
     * @type {Array<string>}
     * @memberof ChatSvcThread
     */
    userIds?: Array<string>;
}
/**
 * Check if a given object implements the ChatSvcThread interface.
 */
export declare function instanceOfChatSvcThread(value: object): value is ChatSvcThread;
export declare function ChatSvcThreadFromJSON(json: any): ChatSvcThread;
export declare function ChatSvcThreadFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcThread;
export declare function ChatSvcThreadToJSON(json: any): ChatSvcThread;
export declare function ChatSvcThreadToJSONTyped(value?: ChatSvcThread | null, ignoreDiscriminator?: boolean): any;
