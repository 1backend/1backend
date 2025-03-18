/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { FileSvcUpload } from './FileSvcUpload';
/**
 *
 * @export
 * @interface FileSvcListUploadsResponse
 */
export interface FileSvcListUploadsResponse {
    /**
     *
     * @type {Array<FileSvcUpload>}
     * @memberof FileSvcListUploadsResponse
     */
    uploads?: Array<FileSvcUpload>;
}
/**
 * Check if a given object implements the FileSvcListUploadsResponse interface.
 */
export declare function instanceOfFileSvcListUploadsResponse(value: object): value is FileSvcListUploadsResponse;
export declare function FileSvcListUploadsResponseFromJSON(json: any): FileSvcListUploadsResponse;
export declare function FileSvcListUploadsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcListUploadsResponse;
export declare function FileSvcListUploadsResponseToJSON(json: any): FileSvcListUploadsResponse;
export declare function FileSvcListUploadsResponseToJSONTyped(value?: FileSvcListUploadsResponse | null, ignoreDiscriminator?: boolean): any;
