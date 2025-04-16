/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
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
 * @interface FileSvcGetDownloadResponse
 */
export interface FileSvcGetDownloadResponse {
    /**
     *
     * @type {FileSvcDownload}
     * @memberof FileSvcGetDownloadResponse
     */
    download?: FileSvcDownload;
    /**
     *
     * @type {boolean}
     * @memberof FileSvcGetDownloadResponse
     */
    _exists: boolean;
}
/**
 * Check if a given object implements the FileSvcGetDownloadResponse interface.
 */
export declare function instanceOfFileSvcGetDownloadResponse(value: object): value is FileSvcGetDownloadResponse;
export declare function FileSvcGetDownloadResponseFromJSON(json: any): FileSvcGetDownloadResponse;
export declare function FileSvcGetDownloadResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcGetDownloadResponse;
export declare function FileSvcGetDownloadResponseToJSON(json: any): FileSvcGetDownloadResponse;
export declare function FileSvcGetDownloadResponseToJSONTyped(value?: FileSvcGetDownloadResponse | null, ignoreDiscriminator?: boolean): any;
