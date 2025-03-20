/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import http from 'http';
import { FileSvcDownloadFileRequest } from '../model/fileSvcDownloadFileRequest';
import { FileSvcDownloadsResponse } from '../model/fileSvcDownloadsResponse';
import { FileSvcGetDownloadResponse } from '../model/fileSvcGetDownloadResponse';
import { FileSvcListUploadsRequest } from '../model/fileSvcListUploadsRequest';
import { FileSvcListUploadsResponse } from '../model/fileSvcListUploadsResponse';
import { FileSvcUploadFileResponse } from '../model/fileSvcUploadFileResponse';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
import { RequestFile } from './apis';
export declare enum FileSvcApiApiKeys {
    BearerAuth = 0
}
export declare class FileSvcApi {
    protected _basePath: string;
    protected _defaultHeaders: any;
    protected _useQuerystring: boolean;
    protected authentications: {
        default: Authentication;
        BearerAuth: ApiKeyAuth;
    };
    protected interceptors: Interceptor[];
    constructor(basePath?: string);
    set useQuerystring(value: boolean);
    set basePath(basePath: string);
    set defaultHeaders(defaultHeaders: any);
    get defaultHeaders(): any;
    get basePath(): string;
    setDefaultAuthentication(auth: Authentication): void;
    setApiKey(key: FileSvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Start or resume the download for a specified URL.  Requires the `file-svc:download:create` permission.
     * @summary Download a File
     * @param body Download Request
     */
    downloadFile(body: FileSvcDownloadFileRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: {
            [key: string]: any;
        };
    }>;
    /**
     * Get a download by URL.  Requires the `file-svc:download:view` permission.
     * @summary Get a Download
     * @param url url
     */
    getDownload(url: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: FileSvcGetDownloadResponse;
    }>;
    /**
     * List download details.  Requires the `file-svc:download:view` permission.
     * @summary List Downloads
     */
    listDownloads(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: FileSvcDownloadsResponse;
    }>;
    /**
     * Lists uploaded files, returning only metadata about each upload. To retrieve file content, use the `Serve an Uploaded File` endpoint, which serves a single file per request. Note: Retrieving the contents of multiple files in a single request is not supported currently.  Requires the `file-svc:upload:view` permission.
     * @summary List Uploads
     * @param body List Uploads Request
     */
    listUploads(body?: FileSvcListUploadsRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: FileSvcListUploadsResponse;
    }>;
    /**
     * Pause a download that is currently in progress.  Requires the `file-svc:download:edit` permission.
     * @summary Pause a Download
     * @param url Download URL
     */
    pauseDownload(url: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: {
            [key: string]: any;
        };
    }>;
    /**
     * Serves a previously downloaded file based on its URL.
     * @summary Serve a Downloaded file
     * @param url URL of the file. Even after downloading, the file is still referenced by its original internet URL.
     */
    serveDownload(url: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: Buffer;
    }>;
    /**
     * Retrieves and serves a previously uploaded file using its File ID. Note: The `ID` and `FileID` fields of an upload are different. - `FileID` is a unique identifier for the file itself. - `ID` is a unique identifier for a specific replica of the file. Since 1Backend is a distributed system, files can be replicated across multiple nodes. This means each uploaded file may have multiple records with the same `FileID` but different `ID`s.
     * @summary Serve an Uploaded File
     * @param fileId FileID uniquely identifies the file itself (not an ID, which represents a specific replica)
     */
    serveUpload(fileId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: Buffer;
    }>;
    /**
     * Uploads a file to the server. Currently if using the clients only one file can be uploaded at a time due to this bug https://github.com/OpenAPITools/openapi-generator/issues/11341 Once that is fixed we should have an `PUT /file-svc/uploads`/uploadFiles (note the plural) endpoints. In reality the endpoint \"unofficially\" supports multiple files. YMMV.  Requires the `file-svc:upload:create` permission.
     * @summary Upload a File
     * @param file File to upload
     */
    uploadFile(file: RequestFile, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: FileSvcUploadFileResponse;
    }>;
}
