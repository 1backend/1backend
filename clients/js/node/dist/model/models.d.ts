import localVarRequest from 'request';
export * from './chatSvcAddMessageRequest';
export * from './chatSvcAddThreadRequest';
export * from './chatSvcAddThreadResponse';
export * from './chatSvcEventMessageAdded';
export * from './chatSvcEventThreadAdded';
export * from './chatSvcEventThreadUpdate';
export * from './chatSvcGetMessageResponse';
export * from './chatSvcGetMessagesResponse';
export * from './chatSvcGetThreadResponse';
export * from './chatSvcGetThreadsResponse';
export * from './chatSvcMessage';
export * from './chatSvcThread';
export * from './chatSvcUpdateThreadRequest';
export * from './configSvcConfig';
export * from './configSvcGetConfigResponse';
export * from './configSvcSaveConfigRequest';
export * from './containerSvcAsset';
export * from './containerSvcBuildImageRequest';
export * from './containerSvcCapabilities';
export * from './containerSvcContainer';
export * from './containerSvcContainerIsRunningResponse';
export * from './containerSvcDaemonInfoResponse';
export * from './containerSvcEnvVar';
export * from './containerSvcErrorResponse';
export * from './containerSvcGetContainerSummaryResponse';
export * from './containerSvcGetHostResponse';
export * from './containerSvcImagePullableResponse';
export * from './containerSvcKeep';
export * from './containerSvcLabel';
export * from './containerSvcListContainersRequest';
export * from './containerSvcListContainersResponse';
export * from './containerSvcListLogsRequest';
export * from './containerSvcListLogsResponse';
export * from './containerSvcLog';
export * from './containerSvcNetwork';
export * from './containerSvcPortMapping';
export * from './containerSvcResources';
export * from './containerSvcRunContainerRequest';
export * from './containerSvcRunContainerResponse';
export * from './containerSvcStopContainerRequest';
export * from './containerSvcVolume';
export * from './dataSvcCreateObjectFields';
export * from './dataSvcCreateObjectRequest';
export * from './dataSvcCreateObjectResponse';
export * from './dataSvcDeleteObjectRequest';
export * from './dataSvcErrorResponse';
export * from './dataSvcObject';
export * from './dataSvcQueryRequest';
export * from './dataSvcQueryResponse';
export * from './dataSvcUpdateObjectsRequest';
export * from './dataSvcUpsertObjectRequest';
export * from './dataSvcUpsertObjectResponse';
export * from './datastoreFilter';
export * from './datastoreOp';
export * from './datastoreOrderBy';
export * from './datastoreQuery';
export * from './datastoreSortingType';
export * from './deploySvcAutoScalingConfig';
export * from './deploySvcDeleteDeploymentRequest';
export * from './deploySvcDeployment';
export * from './deploySvcDeploymentStatus';
export * from './deploySvcDeploymentStrategy';
export * from './deploySvcErrorResponse';
export * from './deploySvcListDeploymentsResponse';
export * from './deploySvcResourceLimits';
export * from './deploySvcSaveDeploymentRequest';
export * from './deploySvcStrategyType';
export * from './deploySvcTargetRegion';
export * from './emailSvcAttachment';
export * from './emailSvcErrorResponse';
export * from './emailSvcSendEmailRequest';
export * from './emailSvcSendEmailResponse';
export * from './fileSvcDownload';
export * from './fileSvcDownloadFileRequest';
export * from './fileSvcDownloadsResponse';
export * from './fileSvcErrorResponse';
export * from './fileSvcGetDownloadResponse';
export * from './fileSvcListUploadsRequest';
export * from './fileSvcListUploadsResponse';
export * from './fileSvcUpload';
export * from './fileSvcUploadFileResponse';
export * from './firehoseSvcErrorResponse';
export * from './firehoseSvcEvent';
export * from './firehoseSvcEventPublishRequest';
export * from './modelSvcArchitectures';
export * from './modelSvcAsset';
export * from './modelSvcContainer';
export * from './modelSvcCudaParameters';
export * from './modelSvcDefaultParameters';
export * from './modelSvcEnvVar';
export * from './modelSvcErrorResponse';
export * from './modelSvcGetModelResponse';
export * from './modelSvcKeep';
export * from './modelSvcListModelsResponse';
export * from './modelSvcListPlatformsResponse';
export * from './modelSvcModel';
export * from './modelSvcModelStatus';
export * from './modelSvcPlatform';
export * from './modelSvcStatusResponse';
export * from './policySvcBlocklistParameters';
export * from './policySvcCheckRequest';
export * from './policySvcCheckResponse';
export * from './policySvcEntity';
export * from './policySvcErrorResponse';
export * from './policySvcInstance';
export * from './policySvcParameters';
export * from './policySvcRateLimitParameters';
export * from './policySvcScope';
export * from './policySvcTemplateId';
export * from './policySvcUpsertInstanceRequest';
export * from './promptSvcEngineParameters';
export * from './promptSvcErrorResponse';
export * from './promptSvcListPromptsRequest';
export * from './promptSvcListPromptsResponse';
export * from './promptSvcLlamaCppParameters';
export * from './promptSvcParameters';
export * from './promptSvcPrompt';
export * from './promptSvcPromptRequest';
export * from './promptSvcPromptResponse';
export * from './promptSvcPromptStatus';
export * from './promptSvcPromptType';
export * from './promptSvcRemovePromptRequest';
export * from './promptSvcStableDiffusionParameters';
export * from './promptSvcStreamChunk';
export * from './promptSvcStreamChunkType';
export * from './promptSvcTextToImageParameters';
export * from './promptSvcTextToTextParameters';
export * from './promptSvcTypesResponse';
export * from './registrySvcAPISpec';
export * from './registrySvcClient';
export * from './registrySvcDefinition';
export * from './registrySvcEnvVar';
export * from './registrySvcErrorResponse';
export * from './registrySvcGPU';
export * from './registrySvcImageSpec';
export * from './registrySvcInstance';
export * from './registrySvcInstanceStatus';
export * from './registrySvcLanguage';
export * from './registrySvcListDefinitionsResponse';
export * from './registrySvcListInstancesResponse';
export * from './registrySvcListNodesRequest';
export * from './registrySvcListNodesResponse';
export * from './registrySvcNode';
export * from './registrySvcNodeSelfResponse';
export * from './registrySvcPortMapping';
export * from './registrySvcProcess';
export * from './registrySvcRegisterInstanceRequest';
export * from './registrySvcRepositorySpec';
export * from './registrySvcResourceUsage';
export * from './registrySvcSaveDefinitionRequest';
export * from './registrySvcUsage';
export * from './secretSvcChecksumAlgorithm';
export * from './secretSvcDecryptValueRequest';
export * from './secretSvcDecryptValueResponse';
export * from './secretSvcEncryptValueRequest';
export * from './secretSvcEncryptValueResponse';
export * from './secretSvcIsSecureResponse';
export * from './secretSvcListSecretsRequest';
export * from './secretSvcListSecretsResponse';
export * from './secretSvcRemoveSecretsRequest';
export * from './secretSvcSaveSecretsRequest';
export * from './secretSvcSecret';
export * from './sourceSvcCheckoutRepoRequest';
export * from './sourceSvcCheckoutRepoResponse';
export * from './sourceSvcErrorResponse';
export * from './stableDiffusionTxt2ImgRequest';
export * from './userSvcAddUserToOrganizationRequest';
export * from './userSvcAssignPermissionsRequest';
export * from './userSvcAuthToken';
export * from './userSvcChangePasswordRequest';
export * from './userSvcContact';
export * from './userSvcCreateOrganizationRequest';
export * from './userSvcCreateRoleRequest';
export * from './userSvcCreateRoleResponse';
export * from './userSvcCreateUserRequest';
export * from './userSvcErrorResponse';
export * from './userSvcGetPermissionsResponse';
export * from './userSvcGetPublicKeyResponse';
export * from './userSvcGetRolesResponse';
export * from './userSvcGetUsersRequest';
export * from './userSvcGetUsersResponse';
export * from './userSvcGrant';
export * from './userSvcIsAuthorizedRequest';
export * from './userSvcIsAuthorizedResponse';
export * from './userSvcListGrantsRequest';
export * from './userSvcListGrantsResponse';
export * from './userSvcLoginRequest';
export * from './userSvcLoginResponse';
export * from './userSvcOrganization';
export * from './userSvcPermission';
export * from './userSvcPermissionLink';
export * from './userSvcReadUserByTokenResponse';
export * from './userSvcRegisterRequest';
export * from './userSvcRegisterResponse';
export * from './userSvcResetPasswordRequest';
export * from './userSvcRole';
export * from './userSvcSaveGrantsRequest';
export * from './userSvcSavePermissionsRequest';
export * from './userSvcSavePermissionsResponse';
export * from './userSvcSaveProfileRequest';
export * from './userSvcSetRolePermissionsRequest';
export * from './userSvcUser';
import * as fs from 'fs';
export interface RequestDetailedFile {
    value: Buffer;
    options?: {
        filename?: string;
        contentType?: string;
    };
}
export type RequestFile = string | Buffer | fs.ReadStream | RequestDetailedFile;
export declare class ObjectSerializer {
    static findCorrectType(data: any, expectedType: string): any;
    static serialize(data: any, type: string): any;
    static deserialize(data: any, type: string): any;
}
export interface Authentication {
    /**
    * Apply authentication settings to header and query params.
    */
    applyToRequest(requestOptions: localVarRequest.Options): Promise<void> | void;
}
export declare class HttpBasicAuth implements Authentication {
    username: string;
    password: string;
    applyToRequest(requestOptions: localVarRequest.Options): void;
}
export declare class HttpBearerAuth implements Authentication {
    accessToken: string | (() => string);
    applyToRequest(requestOptions: localVarRequest.Options): void;
}
export declare class ApiKeyAuth implements Authentication {
    private location;
    private paramName;
    apiKey: string;
    constructor(location: string, paramName: string);
    applyToRequest(requestOptions: localVarRequest.Options): void;
}
export declare class OAuth implements Authentication {
    accessToken: string;
    applyToRequest(requestOptions: localVarRequest.Options): void;
}
export declare class VoidAuth implements Authentication {
    username: string;
    password: string;
    applyToRequest(_: localVarRequest.Options): void;
}
export type Interceptor = (requestOptions: localVarRequest.Options) => (Promise<void> | void);
