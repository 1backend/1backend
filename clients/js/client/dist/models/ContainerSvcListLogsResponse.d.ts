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
import type { ContainerSvcLog } from './ContainerSvcLog';
/**
 *
 * @export
 * @interface ContainerSvcListLogsResponse
 */
export interface ContainerSvcListLogsResponse {
    /**
     *
     * @type {Array<ContainerSvcLog>}
     * @memberof ContainerSvcListLogsResponse
     */
    logs?: Array<ContainerSvcLog>;
}
/**
 * Check if a given object implements the ContainerSvcListLogsResponse interface.
 */
export declare function instanceOfContainerSvcListLogsResponse(value: object): value is ContainerSvcListLogsResponse;
export declare function ContainerSvcListLogsResponseFromJSON(json: any): ContainerSvcListLogsResponse;
export declare function ContainerSvcListLogsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcListLogsResponse;
export declare function ContainerSvcListLogsResponseToJSON(json: any): ContainerSvcListLogsResponse;
export declare function ContainerSvcListLogsResponseToJSONTyped(value?: ContainerSvcListLogsResponse | null, ignoreDiscriminator?: boolean): any;
