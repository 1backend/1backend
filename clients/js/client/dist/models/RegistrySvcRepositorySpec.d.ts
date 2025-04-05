/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface RegistrySvcRepositorySpec
 */
export interface RegistrySvcRepositorySpec {
    /**
     * Context is the path to the image build context
     * @type {string}
     * @memberof RegistrySvcRepositorySpec
     */
    buildContext?: string;
    /**
     * ContainerFile is the path to the file that contains the container build instructions
     * Relative from the build context. By default, it is assumed to be a Dockerfile.
     * @type {string}
     * @memberof RegistrySvcRepositorySpec
     */
    containerFile?: string;
    /**
     * Ports the container will listen on internally
     * @type {Array<number>}
     * @memberof RegistrySvcRepositorySpec
     */
    internalPorts: Array<number>;
    /**
     * URL is the URL to the repository
     * @type {string}
     * @memberof RegistrySvcRepositorySpec
     */
    url: string;
    /**
     * Version of the code to use
     * @type {string}
     * @memberof RegistrySvcRepositorySpec
     */
    version?: string;
}
/**
 * Check if a given object implements the RegistrySvcRepositorySpec interface.
 */
export declare function instanceOfRegistrySvcRepositorySpec(value: object): value is RegistrySvcRepositorySpec;
export declare function RegistrySvcRepositorySpecFromJSON(json: any): RegistrySvcRepositorySpec;
export declare function RegistrySvcRepositorySpecFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcRepositorySpec;
export declare function RegistrySvcRepositorySpecToJSON(json: any): RegistrySvcRepositorySpec;
export declare function RegistrySvcRepositorySpecToJSONTyped(value?: RegistrySvcRepositorySpec | null, ignoreDiscriminator?: boolean): any;
