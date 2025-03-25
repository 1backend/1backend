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
export declare class ChatSvcMessage {
    'createdAt'?: string;
    /**
    * FileIds defines the file attachments the message has.
    */
    'fileIds'?: Array<string>;
    'id': string;
    'meta'?: {
        [key: string]: any;
    };
    /**
    * Text content of the message eg. \"Hi, what\'s up?\"
    */
    'text'?: string;
    /**
    * ThreadId of the message.
    */
    'threadId': string;
    'updatedAt'?: string;
    /**
    * UserId is the id of the user who wrote the message. For AI messages this field is empty.
    */
    'userId'?: string;
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
