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
import type { FileSvcDownload } from './FileSvcDownload';
/**
 *
 * @export
 * @interface FileSvcDownloadsResponse
 */
export interface FileSvcDownloadsResponse {
    /**
     *
     * @type {Array<FileSvcDownload>}
     * @memberof FileSvcDownloadsResponse
     */
    downloads?: Array<FileSvcDownload>;
}
/**
 * Check if a given object implements the FileSvcDownloadsResponse interface.
 */
export declare function instanceOfFileSvcDownloadsResponse(value: object): value is FileSvcDownloadsResponse;
export declare function FileSvcDownloadsResponseFromJSON(json: any): FileSvcDownloadsResponse;
export declare function FileSvcDownloadsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcDownloadsResponse;
export declare function FileSvcDownloadsResponseToJSON(json: any): FileSvcDownloadsResponse;
export declare function FileSvcDownloadsResponseToJSONTyped(value?: FileSvcDownloadsResponse | null, ignoreDiscriminator?: boolean): any;
