/* tslint:disable */
/* eslint-disable */
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


import * as runtime from '../runtime';
import type {
  FileSvcDownloadFileRequest,
  FileSvcDownloadsResponse,
  FileSvcErrorResponse,
  FileSvcGetDownloadResponse,
  FileSvcListUploadsRequest,
  FileSvcListUploadsResponse,
  FileSvcUploadFileResponse,
} from '../models/index';
import {
    FileSvcDownloadFileRequestFromJSON,
    FileSvcDownloadFileRequestToJSON,
    FileSvcDownloadsResponseFromJSON,
    FileSvcDownloadsResponseToJSON,
    FileSvcErrorResponseFromJSON,
    FileSvcErrorResponseToJSON,
    FileSvcGetDownloadResponseFromJSON,
    FileSvcGetDownloadResponseToJSON,
    FileSvcListUploadsRequestFromJSON,
    FileSvcListUploadsRequestToJSON,
    FileSvcListUploadsResponseFromJSON,
    FileSvcListUploadsResponseToJSON,
    FileSvcUploadFileResponseFromJSON,
    FileSvcUploadFileResponseToJSON,
} from '../models/index';

export interface DownloadFileRequest {
    body: FileSvcDownloadFileRequest;
}

export interface GetDownloadRequest {
    url: string;
}

export interface ListUploadsRequest {
    body?: FileSvcListUploadsRequest;
}

export interface PauseDownloadRequest {
    url: string;
}

export interface ServeDownloadRequest {
    url: string;
}

export interface ServeUploadRequest {
    fileId: string;
}

export interface UploadFileRequest {
    file: Blob;
}

/**
 * 
 */
export class FileSvcApi extends runtime.BaseAPI {

    /**
     * Start or resume the download for a specified URL.  Requires the `file-svc:download:create` permission.
     * Download a File
     */
    async downloadFileRaw(requestParameters: DownloadFileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{ [key: string]: any; }>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling downloadFile().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/file-svc/download`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: FileSvcDownloadFileRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Start or resume the download for a specified URL.  Requires the `file-svc:download:create` permission.
     * Download a File
     */
    async downloadFile(requestParameters: DownloadFileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{ [key: string]: any; }> {
        const response = await this.downloadFileRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get a download by URL.  Requires the `file-svc:download:view` permission.
     * Get a Download
     */
    async getDownloadRaw(requestParameters: GetDownloadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FileSvcGetDownloadResponse>> {
        if (requestParameters['url'] == null) {
            throw new runtime.RequiredError(
                'url',
                'Required parameter "url" was null or undefined when calling getDownload().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/file-svc/download/{url}`.replace(`{${"url"}}`, encodeURIComponent(String(requestParameters['url']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => FileSvcGetDownloadResponseFromJSON(jsonValue));
    }

    /**
     * Get a download by URL.  Requires the `file-svc:download:view` permission.
     * Get a Download
     */
    async getDownload(requestParameters: GetDownloadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FileSvcGetDownloadResponse> {
        const response = await this.getDownloadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List download details.  Requires the `file-svc:download:view` permission.
     * List Downloads
     */
    async listDownloadsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FileSvcDownloadsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/file-svc/downloads`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => FileSvcDownloadsResponseFromJSON(jsonValue));
    }

    /**
     * List download details.  Requires the `file-svc:download:view` permission.
     * List Downloads
     */
    async listDownloads(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FileSvcDownloadsResponse> {
        const response = await this.listDownloadsRaw(initOverrides);
        return await response.value();
    }

    /**
     * List the uploaded files.  Requires the `file-svc:upload:view` permission.
     * List Uploads
     */
    async listUploadsRaw(requestParameters: ListUploadsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FileSvcListUploadsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/file-svc/uploads`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: FileSvcListUploadsRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => FileSvcListUploadsResponseFromJSON(jsonValue));
    }

    /**
     * List the uploaded files.  Requires the `file-svc:upload:view` permission.
     * List Uploads
     */
    async listUploads(requestParameters: ListUploadsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FileSvcListUploadsResponse> {
        const response = await this.listUploadsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Pause a download that is currently in progress.  Requires the `file-svc:download:edit` permission.
     * Pause a Download
     */
    async pauseDownloadRaw(requestParameters: PauseDownloadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{ [key: string]: any; }>> {
        if (requestParameters['url'] == null) {
            throw new runtime.RequiredError(
                'url',
                'Required parameter "url" was null or undefined when calling pauseDownload().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/file-svc/download/{url}/pause`.replace(`{${"url"}}`, encodeURIComponent(String(requestParameters['url']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Pause a download that is currently in progress.  Requires the `file-svc:download:edit` permission.
     * Pause a Download
     */
    async pauseDownload(requestParameters: PauseDownloadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{ [key: string]: any; }> {
        const response = await this.pauseDownloadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Serves a previously downloaded file based on its URL.
     * Serve a Downloaded file
     */
    async serveDownloadRaw(requestParameters: ServeDownloadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Blob>> {
        if (requestParameters['url'] == null) {
            throw new runtime.RequiredError(
                'url',
                'Required parameter "url" was null or undefined when calling serveDownload().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/file-svc/serve/download/{url}`.replace(`{${"url"}}`, encodeURIComponent(String(requestParameters['url']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.BlobApiResponse(response);
    }

    /**
     * Serves a previously downloaded file based on its URL.
     * Serve a Downloaded file
     */
    async serveDownload(requestParameters: ServeDownloadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Blob> {
        const response = await this.serveDownloadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieves and serves a previously uploaded file using its File ID. Note: The `ID` and `FileID` fields of an upload are different. - `FileID` is a unique identifier for the file itself. - `ID` is a unique identifier for a specific replica of the file. Since OpenOrch is a distributed system, files can be replicated across multiple nodes. This means each uploaded file may have multiple records with the same `FileID` but different `ID`s.
     * Serve an Uploaded File
     */
    async serveUploadRaw(requestParameters: ServeUploadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Blob>> {
        if (requestParameters['fileId'] == null) {
            throw new runtime.RequiredError(
                'fileId',
                'Required parameter "fileId" was null or undefined when calling serveUpload().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/file-svc/serve/upload/{fileId}`.replace(`{${"fileId"}}`, encodeURIComponent(String(requestParameters['fileId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.BlobApiResponse(response);
    }

    /**
     * Retrieves and serves a previously uploaded file using its File ID. Note: The `ID` and `FileID` fields of an upload are different. - `FileID` is a unique identifier for the file itself. - `ID` is a unique identifier for a specific replica of the file. Since OpenOrch is a distributed system, files can be replicated across multiple nodes. This means each uploaded file may have multiple records with the same `FileID` but different `ID`s.
     * Serve an Uploaded File
     */
    async serveUpload(requestParameters: ServeUploadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Blob> {
        const response = await this.serveUploadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Uploads a file to the server. Currently if using the clients only one file can be uploaded at a time due to this bug https://github.com/OpenAPITools/openapi-generator/issues/11341 Once that is fixed we should have an `PUT /file-svc/uploads`/uploadFiles (note the plural) endpoints. In reality the endpoint \"unofficially\" supports multiple files. YMMV.  Requires the `file-svc:upload:create` permission.
     * Upload a File
     */
    async uploadFileRaw(requestParameters: UploadFileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FileSvcUploadFileResponse>> {
        if (requestParameters['file'] == null) {
            throw new runtime.RequiredError(
                'file',
                'Required parameter "file" was null or undefined when calling uploadFile().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const consumes: runtime.Consume[] = [
            { contentType: 'multipart/form-data' },
        ];
        // @ts-ignore: canConsumeForm may be unused
        const canConsumeForm = runtime.canConsumeForm(consumes);

        let formParams: { append(param: string, value: any): any };
        let useForm = false;
        // use FormData to transmit files using content-type "multipart/form-data"
        useForm = canConsumeForm;
        if (useForm) {
            formParams = new FormData();
        } else {
            formParams = new URLSearchParams();
        }

        if (requestParameters['file'] != null) {
            formParams.append('file', requestParameters['file'] as any);
        }

        const response = await this.request({
            path: `/file-svc/upload`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => FileSvcUploadFileResponseFromJSON(jsonValue));
    }

    /**
     * Uploads a file to the server. Currently if using the clients only one file can be uploaded at a time due to this bug https://github.com/OpenAPITools/openapi-generator/issues/11341 Once that is fixed we should have an `PUT /file-svc/uploads`/uploadFiles (note the plural) endpoints. In reality the endpoint \"unofficially\" supports multiple files. YMMV.  Requires the `file-svc:upload:create` permission.
     * Upload a File
     */
    async uploadFile(requestParameters: UploadFileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FileSvcUploadFileResponse> {
        const response = await this.uploadFileRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
