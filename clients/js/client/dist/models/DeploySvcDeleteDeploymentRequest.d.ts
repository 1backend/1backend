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
 * @interface DeploySvcDeleteDeploymentRequest
 */
export interface DeploySvcDeleteDeploymentRequest {
    /**
     *
     * @type {string}
     * @memberof DeploySvcDeleteDeploymentRequest
     */
    deploymentId: string;
}
/**
 * Check if a given object implements the DeploySvcDeleteDeploymentRequest interface.
 */
export declare function instanceOfDeploySvcDeleteDeploymentRequest(value: object): value is DeploySvcDeleteDeploymentRequest;
export declare function DeploySvcDeleteDeploymentRequestFromJSON(json: any): DeploySvcDeleteDeploymentRequest;
export declare function DeploySvcDeleteDeploymentRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcDeleteDeploymentRequest;
export declare function DeploySvcDeleteDeploymentRequestToJSON(json: any): DeploySvcDeleteDeploymentRequest;
export declare function DeploySvcDeleteDeploymentRequestToJSONTyped(value?: DeploySvcDeleteDeploymentRequest | null, ignoreDiscriminator?: boolean): any;
