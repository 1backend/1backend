/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export declare class RegistrySvcRepositorySpec {
    /**
    * Context is the path to the image build context
    */
    'buildContext'?: string;
    /**
    * ContainerFile is the path to the file that contains the container build instructions Relative from the build context. By default, it is assumed to be a Dockerfile.
    */
    'containerFile'?: string;
    /**
    * Ports the container will listen on internally
    */
    'internalPorts': Array<number>;
    /**
    * URL is the URL to the repository
    */
    'url': string;
    /**
    * Version of the code to use
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
