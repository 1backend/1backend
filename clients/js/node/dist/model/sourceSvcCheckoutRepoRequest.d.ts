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
export declare class SourceSvcCheckoutRepoRequest {
    /**
    * Password or token for HTTPS auth
    */
    'password'?: string;
    /**
    * SSH private key (optional for SSH connection)
    */
    'ssh_key'?: string;
    /**
    * Password for SSH private key if encrypted (optional)
    */
    'ssh_key_pwd'?: string;
    /**
    * Token for HTTPS auth (optional for SSH)
    */
    'token'?: string;
    /**
    * Full repository URL (e.g., https://github.com/user/repo)
    */
    'url'?: string;
    /**
    * Username for HTTPS or SSH user (optional for SSH)
    */
    'username'?: string;
    /**
    * Branch, tag, or commit SHA
    */
    'version'?: string;
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
