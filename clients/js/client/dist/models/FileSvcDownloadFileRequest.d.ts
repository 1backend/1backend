/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface FileSvcDownloadFileRequest
 */
export interface FileSvcDownloadFileRequest {
    /**
     *
     * @type {string}
     * @memberof FileSvcDownloadFileRequest
     */
    folderPath?: string;
    /**
     *
     * @type {string}
     * @memberof FileSvcDownloadFileRequest
     */
    url: string;
}
/**
 * Check if a given object implements the FileSvcDownloadFileRequest interface.
 */
export declare function instanceOfFileSvcDownloadFileRequest(value: object): value is FileSvcDownloadFileRequest;
export declare function FileSvcDownloadFileRequestFromJSON(json: any): FileSvcDownloadFileRequest;
export declare function FileSvcDownloadFileRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcDownloadFileRequest;
export declare function FileSvcDownloadFileRequestToJSON(json: any): FileSvcDownloadFileRequest;
export declare function FileSvcDownloadFileRequestToJSONTyped(value?: FileSvcDownloadFileRequest | null, ignoreDiscriminator?: boolean): any;
