/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcContactInput } from './userSvcContactInput';
export declare class UserSvcRegisterRequest {
    'app'?: string;
    'contact'?: UserSvcContactInput;
    'device'?: string;
    'name'?: string;
    /**
    * Password of the user.
    */
    'password'?: string;
    /**
    * Slug is a URL-friendly unique (inside the 1Backend platform) identifier for the `user`. Required due to its central role in the platform. If your project has no use for a slug, just derive it from the email or similar.
    */
    'slug': string;
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
