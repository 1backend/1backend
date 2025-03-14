/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface FileSvcUpload
 */
export interface FileSvcUpload {
    /**
     *
     * @type {string}
     * @memberof FileSvcUpload
     */
    createdAt?: string;
    /**
     * Logical file ID spanning all replicas
     * @type {string}
     * @memberof FileSvcUpload
     */
    fileId?: string;
    /**
     * Filename is the original name of the file
     * @type {string}
     * @memberof FileSvcUpload
     */
    fileName?: string;
    /**
     * FilePath is the full node local path of the file
     * @type {string}
     * @memberof FileSvcUpload
     */
    filePath?: string;
    /**
     *
     * @type {number}
     * @memberof FileSvcUpload
     */
    fileSize: number;
    /**
     * Unique ID for this replica
     * @type {string}
     * @memberof FileSvcUpload
     */
    id?: string;
    /**
     * ID of the node storing this replica
     * @type {string}
     * @memberof FileSvcUpload
     */
    nodeId?: string;
    /**
     *
     * @type {string}
     * @memberof FileSvcUpload
     */
    updatedAt?: string;
    /**
     *
     * @type {string}
     * @memberof FileSvcUpload
     */
    userId?: string;
}
/**
 * Check if a given object implements the FileSvcUpload interface.
 */
export declare function instanceOfFileSvcUpload(value: object): value is FileSvcUpload;
export declare function FileSvcUploadFromJSON(json: any): FileSvcUpload;
export declare function FileSvcUploadFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcUpload;
export declare function FileSvcUploadToJSON(json: any): FileSvcUpload;
export declare function FileSvcUploadToJSONTyped(value?: FileSvcUpload | null, ignoreDiscriminator?: boolean): any;
