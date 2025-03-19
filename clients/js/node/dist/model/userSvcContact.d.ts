/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export declare class UserSvcContact {
    'createdAt'?: string;
    'deletedAt'?: string;
    /**
    * The unique identifier, which can be a URL.  Example values: \"joe12\" (1backend username), \"twitter.com/thejoe\" (twitter url), \"joe@joesdomain.com\" (email)
    */
    'id'?: string;
    /**
    * If this is the primary contact method
    */
    'isPrimary'?: boolean;
    /**
    * Platform of the contact (e.g., \"email\", \"phone\", \"twitter\")
    */
    'platform'?: string;
    'updatedAt'?: string;
    'userId'?: string;
    /**
    * Value is the platform local unique identifier. Ie. while the `id` of a Twitter contact is `twitter.com/thejoe`, the value will be only `thejoe`. For email and phones the `id` and the `value` will be the same. This field mostly exists for display purposes.  Example values: \"joe12\" (1backend username), \"thejoe\" (twitter username), \"joe@joesdomain.com\" (email)
    */
    'value'?: string;
    /**
    * Whether the contact is verified
    */
    'verified'?: boolean;
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
