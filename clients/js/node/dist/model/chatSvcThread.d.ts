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
export declare class ChatSvcThread {
    'createdAt'?: string;
    'id': string;
    /**
    * Title of the thread.
    */
    'title'?: string;
    /**
    * TopicIds defines which topics the thread belongs to. Topics can roughly be thought of as tags for threads.
    */
    'topicIds'?: Array<string>;
    'updatedAt'?: string;
    /**
    * UserIds the ids of the users who can see this thread.
    */
    'userIds'?: Array<string>;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
