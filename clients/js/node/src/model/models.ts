import localVarRequest from 'request';

export * from './chatSvcEventMessageAdded';
export * from './chatSvcEventThreadAdded';
export * from './chatSvcEventThreadUpdate';
export * from './chatSvcListMessagesRequest';
export * from './chatSvcListMessagesResponse';
export * from './chatSvcListThreadsRequest';
export * from './chatSvcListThreadsResponse';
export * from './chatSvcMessage';
export * from './chatSvcSaveMessageRequest';
export * from './chatSvcSaveThreadRequest';
export * from './chatSvcSaveThreadResponse';
export * from './chatSvcThread';
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
export * from './imageSvcErrorResponse';
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
export * from './proxySvcCert';
export * from './proxySvcErrorResponse';
export * from './proxySvcListCertsRequest';
export * from './proxySvcListCertsResponse';
export * from './proxySvcListRoutesRequest';
export * from './proxySvcListRoutesResponse';
export * from './proxySvcRoute';
export * from './proxySvcRouteInput';
export * from './proxySvcSaveRoutesRequest';
export * from './proxySvcSaveRoutesResponse';
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
export * from './userSvcAuthToken';
export * from './userSvcChangePasswordRequest';
export * from './userSvcContact';
export * from './userSvcContactInput';
export * from './userSvcCreateUserRequest';
export * from './userSvcEnroll';
export * from './userSvcEnrollInput';
export * from './userSvcErrorResponse';
export * from './userSvcExchangeTokenRequest';
export * from './userSvcExchangeTokenResponse';
export * from './userSvcGetPublicKeyResponse';
export * from './userSvcHasPermissionResponse';
export * from './userSvcListEnrollsRequest';
export * from './userSvcListEnrollsResponse';
export * from './userSvcListOrganizationsRequest';
export * from './userSvcListOrganizationsResponse';
export * from './userSvcListPermissionsResponse';
export * from './userSvcListPermitsRequest';
export * from './userSvcListPermitsResponse';
export * from './userSvcListUsersOrderBy';
export * from './userSvcListUsersRequest';
export * from './userSvcListUsersResponse';
export * from './userSvcLoginRequest';
export * from './userSvcLoginResponse';
export * from './userSvcOrderDirection';
export * from './userSvcOrganization';
export * from './userSvcPermit';
export * from './userSvcPermitInput';
export * from './userSvcReadSelfRequest';
export * from './userSvcReadSelfResponse';
export * from './userSvcRefreshTokenResponse';
export * from './userSvcRegisterRequest';
export * from './userSvcRegisterResponse';
export * from './userSvcResetPasswordRequest';
export * from './userSvcRevokeTokensRequest';
export * from './userSvcSaveEnrollsRequest';
export * from './userSvcSaveEnrollsResponse';
export * from './userSvcSaveOrganizationRequest';
export * from './userSvcSaveOrganizationResponse';
export * from './userSvcSavePermitsRequest';
export * from './userSvcSaveSelfRequest';
export * from './userSvcSaveUserRequest';
export * from './userSvcUser';
export * from './userSvcUserInput';
export * from './userSvcUserRecord';

import * as fs from 'fs';

export interface RequestDetailedFile {
    value: Buffer;
    options?: {
        filename?: string;
        contentType?: string;
    }
}

export type RequestFile = string | Buffer | fs.ReadStream | RequestDetailedFile;


import { ChatSvcEventMessageAdded } from './chatSvcEventMessageAdded';
import { ChatSvcEventThreadAdded } from './chatSvcEventThreadAdded';
import { ChatSvcEventThreadUpdate } from './chatSvcEventThreadUpdate';
import { ChatSvcListMessagesRequest } from './chatSvcListMessagesRequest';
import { ChatSvcListMessagesResponse } from './chatSvcListMessagesResponse';
import { ChatSvcListThreadsRequest } from './chatSvcListThreadsRequest';
import { ChatSvcListThreadsResponse } from './chatSvcListThreadsResponse';
import { ChatSvcMessage } from './chatSvcMessage';
import { ChatSvcSaveMessageRequest } from './chatSvcSaveMessageRequest';
import { ChatSvcSaveThreadRequest } from './chatSvcSaveThreadRequest';
import { ChatSvcSaveThreadResponse } from './chatSvcSaveThreadResponse';
import { ChatSvcThread } from './chatSvcThread';
import { ConfigSvcConfig } from './configSvcConfig';
import { ConfigSvcGetConfigResponse } from './configSvcGetConfigResponse';
import { ConfigSvcSaveConfigRequest } from './configSvcSaveConfigRequest';
import { ContainerSvcAsset } from './containerSvcAsset';
import { ContainerSvcBuildImageRequest } from './containerSvcBuildImageRequest';
import { ContainerSvcCapabilities } from './containerSvcCapabilities';
import { ContainerSvcContainer } from './containerSvcContainer';
import { ContainerSvcContainerIsRunningResponse } from './containerSvcContainerIsRunningResponse';
import { ContainerSvcDaemonInfoResponse } from './containerSvcDaemonInfoResponse';
import { ContainerSvcEnvVar } from './containerSvcEnvVar';
import { ContainerSvcErrorResponse } from './containerSvcErrorResponse';
import { ContainerSvcGetContainerSummaryResponse } from './containerSvcGetContainerSummaryResponse';
import { ContainerSvcGetHostResponse } from './containerSvcGetHostResponse';
import { ContainerSvcImagePullableResponse } from './containerSvcImagePullableResponse';
import { ContainerSvcKeep } from './containerSvcKeep';
import { ContainerSvcLabel } from './containerSvcLabel';
import { ContainerSvcListContainersRequest } from './containerSvcListContainersRequest';
import { ContainerSvcListContainersResponse } from './containerSvcListContainersResponse';
import { ContainerSvcListLogsRequest } from './containerSvcListLogsRequest';
import { ContainerSvcListLogsResponse } from './containerSvcListLogsResponse';
import { ContainerSvcLog } from './containerSvcLog';
import { ContainerSvcNetwork } from './containerSvcNetwork';
import { ContainerSvcPortMapping } from './containerSvcPortMapping';
import { ContainerSvcResources } from './containerSvcResources';
import { ContainerSvcRunContainerRequest } from './containerSvcRunContainerRequest';
import { ContainerSvcRunContainerResponse } from './containerSvcRunContainerResponse';
import { ContainerSvcStopContainerRequest } from './containerSvcStopContainerRequest';
import { ContainerSvcVolume } from './containerSvcVolume';
import { DataSvcCreateObjectFields } from './dataSvcCreateObjectFields';
import { DataSvcCreateObjectRequest } from './dataSvcCreateObjectRequest';
import { DataSvcCreateObjectResponse } from './dataSvcCreateObjectResponse';
import { DataSvcDeleteObjectRequest } from './dataSvcDeleteObjectRequest';
import { DataSvcErrorResponse } from './dataSvcErrorResponse';
import { DataSvcObject } from './dataSvcObject';
import { DataSvcQueryRequest } from './dataSvcQueryRequest';
import { DataSvcQueryResponse } from './dataSvcQueryResponse';
import { DataSvcUpdateObjectsRequest } from './dataSvcUpdateObjectsRequest';
import { DataSvcUpsertObjectRequest } from './dataSvcUpsertObjectRequest';
import { DataSvcUpsertObjectResponse } from './dataSvcUpsertObjectResponse';
import { DatastoreFilter } from './datastoreFilter';
import { DatastoreOp } from './datastoreOp';
import { DatastoreOrderBy } from './datastoreOrderBy';
import { DatastoreQuery } from './datastoreQuery';
import { DatastoreSortingType } from './datastoreSortingType';
import { DeploySvcAutoScalingConfig } from './deploySvcAutoScalingConfig';
import { DeploySvcDeleteDeploymentRequest } from './deploySvcDeleteDeploymentRequest';
import { DeploySvcDeployment } from './deploySvcDeployment';
import { DeploySvcDeploymentStatus } from './deploySvcDeploymentStatus';
import { DeploySvcDeploymentStrategy } from './deploySvcDeploymentStrategy';
import { DeploySvcErrorResponse } from './deploySvcErrorResponse';
import { DeploySvcListDeploymentsResponse } from './deploySvcListDeploymentsResponse';
import { DeploySvcResourceLimits } from './deploySvcResourceLimits';
import { DeploySvcSaveDeploymentRequest } from './deploySvcSaveDeploymentRequest';
import { DeploySvcStrategyType } from './deploySvcStrategyType';
import { DeploySvcTargetRegion } from './deploySvcTargetRegion';
import { EmailSvcAttachment } from './emailSvcAttachment';
import { EmailSvcErrorResponse } from './emailSvcErrorResponse';
import { EmailSvcSendEmailRequest } from './emailSvcSendEmailRequest';
import { EmailSvcSendEmailResponse } from './emailSvcSendEmailResponse';
import { FileSvcDownload } from './fileSvcDownload';
import { FileSvcDownloadFileRequest } from './fileSvcDownloadFileRequest';
import { FileSvcDownloadsResponse } from './fileSvcDownloadsResponse';
import { FileSvcErrorResponse } from './fileSvcErrorResponse';
import { FileSvcGetDownloadResponse } from './fileSvcGetDownloadResponse';
import { FileSvcListUploadsRequest } from './fileSvcListUploadsRequest';
import { FileSvcListUploadsResponse } from './fileSvcListUploadsResponse';
import { FileSvcUpload } from './fileSvcUpload';
import { FileSvcUploadFileResponse } from './fileSvcUploadFileResponse';
import { FirehoseSvcErrorResponse } from './firehoseSvcErrorResponse';
import { FirehoseSvcEvent } from './firehoseSvcEvent';
import { FirehoseSvcEventPublishRequest } from './firehoseSvcEventPublishRequest';
import { ImageSvcErrorResponse } from './imageSvcErrorResponse';
import { ModelSvcArchitectures } from './modelSvcArchitectures';
import { ModelSvcAsset } from './modelSvcAsset';
import { ModelSvcContainer } from './modelSvcContainer';
import { ModelSvcCudaParameters } from './modelSvcCudaParameters';
import { ModelSvcDefaultParameters } from './modelSvcDefaultParameters';
import { ModelSvcEnvVar } from './modelSvcEnvVar';
import { ModelSvcErrorResponse } from './modelSvcErrorResponse';
import { ModelSvcGetModelResponse } from './modelSvcGetModelResponse';
import { ModelSvcKeep } from './modelSvcKeep';
import { ModelSvcListModelsResponse } from './modelSvcListModelsResponse';
import { ModelSvcListPlatformsResponse } from './modelSvcListPlatformsResponse';
import { ModelSvcModel } from './modelSvcModel';
import { ModelSvcModelStatus } from './modelSvcModelStatus';
import { ModelSvcPlatform } from './modelSvcPlatform';
import { ModelSvcStatusResponse } from './modelSvcStatusResponse';
import { PolicySvcBlocklistParameters } from './policySvcBlocklistParameters';
import { PolicySvcCheckRequest } from './policySvcCheckRequest';
import { PolicySvcCheckResponse } from './policySvcCheckResponse';
import { PolicySvcEntity } from './policySvcEntity';
import { PolicySvcErrorResponse } from './policySvcErrorResponse';
import { PolicySvcInstance } from './policySvcInstance';
import { PolicySvcParameters } from './policySvcParameters';
import { PolicySvcRateLimitParameters } from './policySvcRateLimitParameters';
import { PolicySvcScope } from './policySvcScope';
import { PolicySvcTemplateId } from './policySvcTemplateId';
import { PolicySvcUpsertInstanceRequest } from './policySvcUpsertInstanceRequest';
import { PromptSvcEngineParameters } from './promptSvcEngineParameters';
import { PromptSvcErrorResponse } from './promptSvcErrorResponse';
import { PromptSvcListPromptsRequest } from './promptSvcListPromptsRequest';
import { PromptSvcListPromptsResponse } from './promptSvcListPromptsResponse';
import { PromptSvcLlamaCppParameters } from './promptSvcLlamaCppParameters';
import { PromptSvcParameters } from './promptSvcParameters';
import { PromptSvcPrompt } from './promptSvcPrompt';
import { PromptSvcPromptRequest } from './promptSvcPromptRequest';
import { PromptSvcPromptResponse } from './promptSvcPromptResponse';
import { PromptSvcPromptStatus } from './promptSvcPromptStatus';
import { PromptSvcPromptType } from './promptSvcPromptType';
import { PromptSvcRemovePromptRequest } from './promptSvcRemovePromptRequest';
import { PromptSvcStableDiffusionParameters } from './promptSvcStableDiffusionParameters';
import { PromptSvcStreamChunk } from './promptSvcStreamChunk';
import { PromptSvcStreamChunkType } from './promptSvcStreamChunkType';
import { PromptSvcTextToImageParameters } from './promptSvcTextToImageParameters';
import { PromptSvcTextToTextParameters } from './promptSvcTextToTextParameters';
import { PromptSvcTypesResponse } from './promptSvcTypesResponse';
import { ProxySvcCert } from './proxySvcCert';
import { ProxySvcErrorResponse } from './proxySvcErrorResponse';
import { ProxySvcListCertsRequest } from './proxySvcListCertsRequest';
import { ProxySvcListCertsResponse } from './proxySvcListCertsResponse';
import { ProxySvcListRoutesRequest } from './proxySvcListRoutesRequest';
import { ProxySvcListRoutesResponse } from './proxySvcListRoutesResponse';
import { ProxySvcRoute } from './proxySvcRoute';
import { ProxySvcRouteInput } from './proxySvcRouteInput';
import { ProxySvcSaveRoutesRequest } from './proxySvcSaveRoutesRequest';
import { ProxySvcSaveRoutesResponse } from './proxySvcSaveRoutesResponse';
import { RegistrySvcAPISpec } from './registrySvcAPISpec';
import { RegistrySvcClient } from './registrySvcClient';
import { RegistrySvcDefinition } from './registrySvcDefinition';
import { RegistrySvcEnvVar } from './registrySvcEnvVar';
import { RegistrySvcErrorResponse } from './registrySvcErrorResponse';
import { RegistrySvcGPU } from './registrySvcGPU';
import { RegistrySvcImageSpec } from './registrySvcImageSpec';
import { RegistrySvcInstance } from './registrySvcInstance';
import { RegistrySvcInstanceStatus } from './registrySvcInstanceStatus';
import { RegistrySvcLanguage } from './registrySvcLanguage';
import { RegistrySvcListDefinitionsResponse } from './registrySvcListDefinitionsResponse';
import { RegistrySvcListInstancesResponse } from './registrySvcListInstancesResponse';
import { RegistrySvcListNodesRequest } from './registrySvcListNodesRequest';
import { RegistrySvcListNodesResponse } from './registrySvcListNodesResponse';
import { RegistrySvcNode } from './registrySvcNode';
import { RegistrySvcNodeSelfResponse } from './registrySvcNodeSelfResponse';
import { RegistrySvcPortMapping } from './registrySvcPortMapping';
import { RegistrySvcProcess } from './registrySvcProcess';
import { RegistrySvcRegisterInstanceRequest } from './registrySvcRegisterInstanceRequest';
import { RegistrySvcRepositorySpec } from './registrySvcRepositorySpec';
import { RegistrySvcResourceUsage } from './registrySvcResourceUsage';
import { RegistrySvcSaveDefinitionRequest } from './registrySvcSaveDefinitionRequest';
import { RegistrySvcUsage } from './registrySvcUsage';
import { SecretSvcChecksumAlgorithm } from './secretSvcChecksumAlgorithm';
import { SecretSvcDecryptValueRequest } from './secretSvcDecryptValueRequest';
import { SecretSvcDecryptValueResponse } from './secretSvcDecryptValueResponse';
import { SecretSvcEncryptValueRequest } from './secretSvcEncryptValueRequest';
import { SecretSvcEncryptValueResponse } from './secretSvcEncryptValueResponse';
import { SecretSvcIsSecureResponse } from './secretSvcIsSecureResponse';
import { SecretSvcListSecretsRequest } from './secretSvcListSecretsRequest';
import { SecretSvcListSecretsResponse } from './secretSvcListSecretsResponse';
import { SecretSvcRemoveSecretsRequest } from './secretSvcRemoveSecretsRequest';
import { SecretSvcSaveSecretsRequest } from './secretSvcSaveSecretsRequest';
import { SecretSvcSecret } from './secretSvcSecret';
import { SourceSvcCheckoutRepoRequest } from './sourceSvcCheckoutRepoRequest';
import { SourceSvcCheckoutRepoResponse } from './sourceSvcCheckoutRepoResponse';
import { SourceSvcErrorResponse } from './sourceSvcErrorResponse';
import { StableDiffusionTxt2ImgRequest } from './stableDiffusionTxt2ImgRequest';
import { UserSvcAuthToken } from './userSvcAuthToken';
import { UserSvcChangePasswordRequest } from './userSvcChangePasswordRequest';
import { UserSvcContact } from './userSvcContact';
import { UserSvcContactInput } from './userSvcContactInput';
import { UserSvcCreateUserRequest } from './userSvcCreateUserRequest';
import { UserSvcEnroll } from './userSvcEnroll';
import { UserSvcEnrollInput } from './userSvcEnrollInput';
import { UserSvcErrorResponse } from './userSvcErrorResponse';
import { UserSvcExchangeTokenRequest } from './userSvcExchangeTokenRequest';
import { UserSvcExchangeTokenResponse } from './userSvcExchangeTokenResponse';
import { UserSvcGetPublicKeyResponse } from './userSvcGetPublicKeyResponse';
import { UserSvcHasPermissionResponse } from './userSvcHasPermissionResponse';
import { UserSvcListEnrollsRequest } from './userSvcListEnrollsRequest';
import { UserSvcListEnrollsResponse } from './userSvcListEnrollsResponse';
import { UserSvcListOrganizationsRequest } from './userSvcListOrganizationsRequest';
import { UserSvcListOrganizationsResponse } from './userSvcListOrganizationsResponse';
import { UserSvcListPermissionsResponse } from './userSvcListPermissionsResponse';
import { UserSvcListPermitsRequest } from './userSvcListPermitsRequest';
import { UserSvcListPermitsResponse } from './userSvcListPermitsResponse';
import { UserSvcListUsersOrderBy } from './userSvcListUsersOrderBy';
import { UserSvcListUsersRequest } from './userSvcListUsersRequest';
import { UserSvcListUsersResponse } from './userSvcListUsersResponse';
import { UserSvcLoginRequest } from './userSvcLoginRequest';
import { UserSvcLoginResponse } from './userSvcLoginResponse';
import { UserSvcOrderDirection } from './userSvcOrderDirection';
import { UserSvcOrganization } from './userSvcOrganization';
import { UserSvcPermit } from './userSvcPermit';
import { UserSvcPermitInput } from './userSvcPermitInput';
import { UserSvcReadSelfRequest } from './userSvcReadSelfRequest';
import { UserSvcReadSelfResponse } from './userSvcReadSelfResponse';
import { UserSvcRefreshTokenResponse } from './userSvcRefreshTokenResponse';
import { UserSvcRegisterRequest } from './userSvcRegisterRequest';
import { UserSvcRegisterResponse } from './userSvcRegisterResponse';
import { UserSvcResetPasswordRequest } from './userSvcResetPasswordRequest';
import { UserSvcRevokeTokensRequest } from './userSvcRevokeTokensRequest';
import { UserSvcSaveEnrollsRequest } from './userSvcSaveEnrollsRequest';
import { UserSvcSaveEnrollsResponse } from './userSvcSaveEnrollsResponse';
import { UserSvcSaveOrganizationRequest } from './userSvcSaveOrganizationRequest';
import { UserSvcSaveOrganizationResponse } from './userSvcSaveOrganizationResponse';
import { UserSvcSavePermitsRequest } from './userSvcSavePermitsRequest';
import { UserSvcSaveSelfRequest } from './userSvcSaveSelfRequest';
import { UserSvcSaveUserRequest } from './userSvcSaveUserRequest';
import { UserSvcUser } from './userSvcUser';
import { UserSvcUserInput } from './userSvcUserInput';
import { UserSvcUserRecord } from './userSvcUserRecord';

/* tslint:disable:no-unused-variable */
let primitives = [
                    "string",
                    "boolean",
                    "double",
                    "integer",
                    "long",
                    "float",
                    "number",
                    "any"
                 ];

let enumsMap: {[index: string]: any} = {
        "DatastoreOp": DatastoreOp,
        "DatastoreSortingType": DatastoreSortingType,
        "DeploySvcDeploymentStatus": DeploySvcDeploymentStatus,
        "DeploySvcStrategyType": DeploySvcStrategyType,
        "PolicySvcEntity": PolicySvcEntity,
        "PolicySvcScope": PolicySvcScope,
        "PolicySvcTemplateId": PolicySvcTemplateId,
        "PromptSvcPromptStatus": PromptSvcPromptStatus,
        "PromptSvcPromptType": PromptSvcPromptType,
        "PromptSvcStreamChunkType": PromptSvcStreamChunkType,
        "RegistrySvcInstanceStatus": RegistrySvcInstanceStatus,
        "RegistrySvcLanguage": RegistrySvcLanguage,
        "SecretSvcChecksumAlgorithm": SecretSvcChecksumAlgorithm,
        "UserSvcListUsersOrderBy": UserSvcListUsersOrderBy,
        "UserSvcOrderDirection": UserSvcOrderDirection,
}

let typeMap: {[index: string]: any} = {
    "ChatSvcEventMessageAdded": ChatSvcEventMessageAdded,
    "ChatSvcEventThreadAdded": ChatSvcEventThreadAdded,
    "ChatSvcEventThreadUpdate": ChatSvcEventThreadUpdate,
    "ChatSvcListMessagesRequest": ChatSvcListMessagesRequest,
    "ChatSvcListMessagesResponse": ChatSvcListMessagesResponse,
    "ChatSvcListThreadsRequest": ChatSvcListThreadsRequest,
    "ChatSvcListThreadsResponse": ChatSvcListThreadsResponse,
    "ChatSvcMessage": ChatSvcMessage,
    "ChatSvcSaveMessageRequest": ChatSvcSaveMessageRequest,
    "ChatSvcSaveThreadRequest": ChatSvcSaveThreadRequest,
    "ChatSvcSaveThreadResponse": ChatSvcSaveThreadResponse,
    "ChatSvcThread": ChatSvcThread,
    "ConfigSvcConfig": ConfigSvcConfig,
    "ConfigSvcGetConfigResponse": ConfigSvcGetConfigResponse,
    "ConfigSvcSaveConfigRequest": ConfigSvcSaveConfigRequest,
    "ContainerSvcAsset": ContainerSvcAsset,
    "ContainerSvcBuildImageRequest": ContainerSvcBuildImageRequest,
    "ContainerSvcCapabilities": ContainerSvcCapabilities,
    "ContainerSvcContainer": ContainerSvcContainer,
    "ContainerSvcContainerIsRunningResponse": ContainerSvcContainerIsRunningResponse,
    "ContainerSvcDaemonInfoResponse": ContainerSvcDaemonInfoResponse,
    "ContainerSvcEnvVar": ContainerSvcEnvVar,
    "ContainerSvcErrorResponse": ContainerSvcErrorResponse,
    "ContainerSvcGetContainerSummaryResponse": ContainerSvcGetContainerSummaryResponse,
    "ContainerSvcGetHostResponse": ContainerSvcGetHostResponse,
    "ContainerSvcImagePullableResponse": ContainerSvcImagePullableResponse,
    "ContainerSvcKeep": ContainerSvcKeep,
    "ContainerSvcLabel": ContainerSvcLabel,
    "ContainerSvcListContainersRequest": ContainerSvcListContainersRequest,
    "ContainerSvcListContainersResponse": ContainerSvcListContainersResponse,
    "ContainerSvcListLogsRequest": ContainerSvcListLogsRequest,
    "ContainerSvcListLogsResponse": ContainerSvcListLogsResponse,
    "ContainerSvcLog": ContainerSvcLog,
    "ContainerSvcNetwork": ContainerSvcNetwork,
    "ContainerSvcPortMapping": ContainerSvcPortMapping,
    "ContainerSvcResources": ContainerSvcResources,
    "ContainerSvcRunContainerRequest": ContainerSvcRunContainerRequest,
    "ContainerSvcRunContainerResponse": ContainerSvcRunContainerResponse,
    "ContainerSvcStopContainerRequest": ContainerSvcStopContainerRequest,
    "ContainerSvcVolume": ContainerSvcVolume,
    "DataSvcCreateObjectFields": DataSvcCreateObjectFields,
    "DataSvcCreateObjectRequest": DataSvcCreateObjectRequest,
    "DataSvcCreateObjectResponse": DataSvcCreateObjectResponse,
    "DataSvcDeleteObjectRequest": DataSvcDeleteObjectRequest,
    "DataSvcErrorResponse": DataSvcErrorResponse,
    "DataSvcObject": DataSvcObject,
    "DataSvcQueryRequest": DataSvcQueryRequest,
    "DataSvcQueryResponse": DataSvcQueryResponse,
    "DataSvcUpdateObjectsRequest": DataSvcUpdateObjectsRequest,
    "DataSvcUpsertObjectRequest": DataSvcUpsertObjectRequest,
    "DataSvcUpsertObjectResponse": DataSvcUpsertObjectResponse,
    "DatastoreFilter": DatastoreFilter,
    "DatastoreOrderBy": DatastoreOrderBy,
    "DatastoreQuery": DatastoreQuery,
    "DeploySvcAutoScalingConfig": DeploySvcAutoScalingConfig,
    "DeploySvcDeleteDeploymentRequest": DeploySvcDeleteDeploymentRequest,
    "DeploySvcDeployment": DeploySvcDeployment,
    "DeploySvcDeploymentStrategy": DeploySvcDeploymentStrategy,
    "DeploySvcErrorResponse": DeploySvcErrorResponse,
    "DeploySvcListDeploymentsResponse": DeploySvcListDeploymentsResponse,
    "DeploySvcResourceLimits": DeploySvcResourceLimits,
    "DeploySvcSaveDeploymentRequest": DeploySvcSaveDeploymentRequest,
    "DeploySvcTargetRegion": DeploySvcTargetRegion,
    "EmailSvcAttachment": EmailSvcAttachment,
    "EmailSvcErrorResponse": EmailSvcErrorResponse,
    "EmailSvcSendEmailRequest": EmailSvcSendEmailRequest,
    "EmailSvcSendEmailResponse": EmailSvcSendEmailResponse,
    "FileSvcDownload": FileSvcDownload,
    "FileSvcDownloadFileRequest": FileSvcDownloadFileRequest,
    "FileSvcDownloadsResponse": FileSvcDownloadsResponse,
    "FileSvcErrorResponse": FileSvcErrorResponse,
    "FileSvcGetDownloadResponse": FileSvcGetDownloadResponse,
    "FileSvcListUploadsRequest": FileSvcListUploadsRequest,
    "FileSvcListUploadsResponse": FileSvcListUploadsResponse,
    "FileSvcUpload": FileSvcUpload,
    "FileSvcUploadFileResponse": FileSvcUploadFileResponse,
    "FirehoseSvcErrorResponse": FirehoseSvcErrorResponse,
    "FirehoseSvcEvent": FirehoseSvcEvent,
    "FirehoseSvcEventPublishRequest": FirehoseSvcEventPublishRequest,
    "ImageSvcErrorResponse": ImageSvcErrorResponse,
    "ModelSvcArchitectures": ModelSvcArchitectures,
    "ModelSvcAsset": ModelSvcAsset,
    "ModelSvcContainer": ModelSvcContainer,
    "ModelSvcCudaParameters": ModelSvcCudaParameters,
    "ModelSvcDefaultParameters": ModelSvcDefaultParameters,
    "ModelSvcEnvVar": ModelSvcEnvVar,
    "ModelSvcErrorResponse": ModelSvcErrorResponse,
    "ModelSvcGetModelResponse": ModelSvcGetModelResponse,
    "ModelSvcKeep": ModelSvcKeep,
    "ModelSvcListModelsResponse": ModelSvcListModelsResponse,
    "ModelSvcListPlatformsResponse": ModelSvcListPlatformsResponse,
    "ModelSvcModel": ModelSvcModel,
    "ModelSvcModelStatus": ModelSvcModelStatus,
    "ModelSvcPlatform": ModelSvcPlatform,
    "ModelSvcStatusResponse": ModelSvcStatusResponse,
    "PolicySvcBlocklistParameters": PolicySvcBlocklistParameters,
    "PolicySvcCheckRequest": PolicySvcCheckRequest,
    "PolicySvcCheckResponse": PolicySvcCheckResponse,
    "PolicySvcErrorResponse": PolicySvcErrorResponse,
    "PolicySvcInstance": PolicySvcInstance,
    "PolicySvcParameters": PolicySvcParameters,
    "PolicySvcRateLimitParameters": PolicySvcRateLimitParameters,
    "PolicySvcUpsertInstanceRequest": PolicySvcUpsertInstanceRequest,
    "PromptSvcEngineParameters": PromptSvcEngineParameters,
    "PromptSvcErrorResponse": PromptSvcErrorResponse,
    "PromptSvcListPromptsRequest": PromptSvcListPromptsRequest,
    "PromptSvcListPromptsResponse": PromptSvcListPromptsResponse,
    "PromptSvcLlamaCppParameters": PromptSvcLlamaCppParameters,
    "PromptSvcParameters": PromptSvcParameters,
    "PromptSvcPrompt": PromptSvcPrompt,
    "PromptSvcPromptRequest": PromptSvcPromptRequest,
    "PromptSvcPromptResponse": PromptSvcPromptResponse,
    "PromptSvcRemovePromptRequest": PromptSvcRemovePromptRequest,
    "PromptSvcStableDiffusionParameters": PromptSvcStableDiffusionParameters,
    "PromptSvcStreamChunk": PromptSvcStreamChunk,
    "PromptSvcTextToImageParameters": PromptSvcTextToImageParameters,
    "PromptSvcTextToTextParameters": PromptSvcTextToTextParameters,
    "PromptSvcTypesResponse": PromptSvcTypesResponse,
    "ProxySvcCert": ProxySvcCert,
    "ProxySvcErrorResponse": ProxySvcErrorResponse,
    "ProxySvcListCertsRequest": ProxySvcListCertsRequest,
    "ProxySvcListCertsResponse": ProxySvcListCertsResponse,
    "ProxySvcListRoutesRequest": ProxySvcListRoutesRequest,
    "ProxySvcListRoutesResponse": ProxySvcListRoutesResponse,
    "ProxySvcRoute": ProxySvcRoute,
    "ProxySvcRouteInput": ProxySvcRouteInput,
    "ProxySvcSaveRoutesRequest": ProxySvcSaveRoutesRequest,
    "ProxySvcSaveRoutesResponse": ProxySvcSaveRoutesResponse,
    "RegistrySvcAPISpec": RegistrySvcAPISpec,
    "RegistrySvcClient": RegistrySvcClient,
    "RegistrySvcDefinition": RegistrySvcDefinition,
    "RegistrySvcEnvVar": RegistrySvcEnvVar,
    "RegistrySvcErrorResponse": RegistrySvcErrorResponse,
    "RegistrySvcGPU": RegistrySvcGPU,
    "RegistrySvcImageSpec": RegistrySvcImageSpec,
    "RegistrySvcInstance": RegistrySvcInstance,
    "RegistrySvcListDefinitionsResponse": RegistrySvcListDefinitionsResponse,
    "RegistrySvcListInstancesResponse": RegistrySvcListInstancesResponse,
    "RegistrySvcListNodesRequest": RegistrySvcListNodesRequest,
    "RegistrySvcListNodesResponse": RegistrySvcListNodesResponse,
    "RegistrySvcNode": RegistrySvcNode,
    "RegistrySvcNodeSelfResponse": RegistrySvcNodeSelfResponse,
    "RegistrySvcPortMapping": RegistrySvcPortMapping,
    "RegistrySvcProcess": RegistrySvcProcess,
    "RegistrySvcRegisterInstanceRequest": RegistrySvcRegisterInstanceRequest,
    "RegistrySvcRepositorySpec": RegistrySvcRepositorySpec,
    "RegistrySvcResourceUsage": RegistrySvcResourceUsage,
    "RegistrySvcSaveDefinitionRequest": RegistrySvcSaveDefinitionRequest,
    "RegistrySvcUsage": RegistrySvcUsage,
    "SecretSvcDecryptValueRequest": SecretSvcDecryptValueRequest,
    "SecretSvcDecryptValueResponse": SecretSvcDecryptValueResponse,
    "SecretSvcEncryptValueRequest": SecretSvcEncryptValueRequest,
    "SecretSvcEncryptValueResponse": SecretSvcEncryptValueResponse,
    "SecretSvcIsSecureResponse": SecretSvcIsSecureResponse,
    "SecretSvcListSecretsRequest": SecretSvcListSecretsRequest,
    "SecretSvcListSecretsResponse": SecretSvcListSecretsResponse,
    "SecretSvcRemoveSecretsRequest": SecretSvcRemoveSecretsRequest,
    "SecretSvcSaveSecretsRequest": SecretSvcSaveSecretsRequest,
    "SecretSvcSecret": SecretSvcSecret,
    "SourceSvcCheckoutRepoRequest": SourceSvcCheckoutRepoRequest,
    "SourceSvcCheckoutRepoResponse": SourceSvcCheckoutRepoResponse,
    "SourceSvcErrorResponse": SourceSvcErrorResponse,
    "StableDiffusionTxt2ImgRequest": StableDiffusionTxt2ImgRequest,
    "UserSvcAuthToken": UserSvcAuthToken,
    "UserSvcChangePasswordRequest": UserSvcChangePasswordRequest,
    "UserSvcContact": UserSvcContact,
    "UserSvcContactInput": UserSvcContactInput,
    "UserSvcCreateUserRequest": UserSvcCreateUserRequest,
    "UserSvcEnroll": UserSvcEnroll,
    "UserSvcEnrollInput": UserSvcEnrollInput,
    "UserSvcErrorResponse": UserSvcErrorResponse,
    "UserSvcExchangeTokenRequest": UserSvcExchangeTokenRequest,
    "UserSvcExchangeTokenResponse": UserSvcExchangeTokenResponse,
    "UserSvcGetPublicKeyResponse": UserSvcGetPublicKeyResponse,
    "UserSvcHasPermissionResponse": UserSvcHasPermissionResponse,
    "UserSvcListEnrollsRequest": UserSvcListEnrollsRequest,
    "UserSvcListEnrollsResponse": UserSvcListEnrollsResponse,
    "UserSvcListOrganizationsRequest": UserSvcListOrganizationsRequest,
    "UserSvcListOrganizationsResponse": UserSvcListOrganizationsResponse,
    "UserSvcListPermissionsResponse": UserSvcListPermissionsResponse,
    "UserSvcListPermitsRequest": UserSvcListPermitsRequest,
    "UserSvcListPermitsResponse": UserSvcListPermitsResponse,
    "UserSvcListUsersRequest": UserSvcListUsersRequest,
    "UserSvcListUsersResponse": UserSvcListUsersResponse,
    "UserSvcLoginRequest": UserSvcLoginRequest,
    "UserSvcLoginResponse": UserSvcLoginResponse,
    "UserSvcOrganization": UserSvcOrganization,
    "UserSvcPermit": UserSvcPermit,
    "UserSvcPermitInput": UserSvcPermitInput,
    "UserSvcReadSelfRequest": UserSvcReadSelfRequest,
    "UserSvcReadSelfResponse": UserSvcReadSelfResponse,
    "UserSvcRefreshTokenResponse": UserSvcRefreshTokenResponse,
    "UserSvcRegisterRequest": UserSvcRegisterRequest,
    "UserSvcRegisterResponse": UserSvcRegisterResponse,
    "UserSvcResetPasswordRequest": UserSvcResetPasswordRequest,
    "UserSvcRevokeTokensRequest": UserSvcRevokeTokensRequest,
    "UserSvcSaveEnrollsRequest": UserSvcSaveEnrollsRequest,
    "UserSvcSaveEnrollsResponse": UserSvcSaveEnrollsResponse,
    "UserSvcSaveOrganizationRequest": UserSvcSaveOrganizationRequest,
    "UserSvcSaveOrganizationResponse": UserSvcSaveOrganizationResponse,
    "UserSvcSavePermitsRequest": UserSvcSavePermitsRequest,
    "UserSvcSaveSelfRequest": UserSvcSaveSelfRequest,
    "UserSvcSaveUserRequest": UserSvcSaveUserRequest,
    "UserSvcUser": UserSvcUser,
    "UserSvcUserInput": UserSvcUserInput,
    "UserSvcUserRecord": UserSvcUserRecord,
}

// Check if a string starts with another string without using es6 features
function startsWith(str: string, match: string): boolean {
    return str.substring(0, match.length) === match;
}

// Check if a string ends with another string without using es6 features
function endsWith(str: string, match: string): boolean {
    return str.length >= match.length && str.substring(str.length - match.length) === match;
}

const nullableSuffix = " | null";
const optionalSuffix = " | undefined";
const arrayPrefix = "Array<";
const arraySuffix = ">";
const mapPrefix = "{ [key: string]: ";
const mapSuffix = "; }";

export class ObjectSerializer {
    public static findCorrectType(data: any, expectedType: string) {
        if (data == undefined) {
            return expectedType;
        } else if (primitives.indexOf(expectedType.toLowerCase()) !== -1) {
            return expectedType;
        } else if (expectedType === "Date") {
            return expectedType;
        } else {
            if (enumsMap[expectedType]) {
                return expectedType;
            }

            if (!typeMap[expectedType]) {
                return expectedType; // w/e we don't know the type
            }

            // Check the discriminator
            let discriminatorProperty = typeMap[expectedType].discriminator;
            if (discriminatorProperty == null) {
                return expectedType; // the type does not have a discriminator. use it.
            } else {
                if (data[discriminatorProperty]) {
                    var discriminatorType = data[discriminatorProperty];
                    if(typeMap[discriminatorType]){
                        return discriminatorType; // use the type given in the discriminator
                    } else {
                        return expectedType; // discriminator did not map to a type
                    }
                } else {
                    return expectedType; // discriminator was not present (or an empty string)
                }
            }
        }
    }

    public static serialize(data: any, type: string): any {
        if (data == undefined) {
            return data;
        } else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        } else if (endsWith(type, nullableSuffix)) {
            let subType: string = type.slice(0, -nullableSuffix.length); // Type | null => Type
            return ObjectSerializer.serialize(data, subType);
        } else if (endsWith(type, optionalSuffix)) {
            let subType: string = type.slice(0, -optionalSuffix.length); // Type | undefined => Type
            return ObjectSerializer.serialize(data, subType);
        } else if (startsWith(type, arrayPrefix)) {
            let subType: string = type.slice(arrayPrefix.length, -arraySuffix.length); // Array<Type> => Type
            let transformedData: any[] = [];
            for (let index = 0; index < data.length; index++) {
                let datum = data[index];
                transformedData.push(ObjectSerializer.serialize(datum, subType));
            }
            return transformedData;
        } else if (startsWith(type, mapPrefix)) {
            let subType: string = type.slice(mapPrefix.length, -mapSuffix.length); // { [key: string]: Type; } => Type
            let transformedData: { [key: string]: any } = {};
            for (let key in data) {
                transformedData[key] = ObjectSerializer.serialize(
                    data[key],
                    subType,
                );
            }
            return transformedData;
        } else if (type === "Date") {
            return data.toISOString();
        } else {
            if (enumsMap[type]) {
                return data;
            }
            if (!typeMap[type]) { // in case we dont know the type
                return data;
            }

            // Get the actual type of this object
            type = this.findCorrectType(data, type);

            // get the map for the correct type.
            let attributeTypes = typeMap[type].getAttributeTypeMap();
            let instance: {[index: string]: any} = {};
            for (let index = 0; index < attributeTypes.length; index++) {
                let attributeType = attributeTypes[index];
                instance[attributeType.baseName] = ObjectSerializer.serialize(data[attributeType.name], attributeType.type);
            }
            return instance;
        }
    }

    public static deserialize(data: any, type: string): any {
        // polymorphism may change the actual type.
        type = ObjectSerializer.findCorrectType(data, type);
        if (data == undefined) {
            return data;
        } else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        } else if (endsWith(type, nullableSuffix)) {
            let subType: string = type.slice(0, -nullableSuffix.length); // Type | null => Type
            return ObjectSerializer.deserialize(data, subType);
        } else if (endsWith(type, optionalSuffix)) {
            let subType: string = type.slice(0, -optionalSuffix.length); // Type | undefined => Type
            return ObjectSerializer.deserialize(data, subType);
        } else if (startsWith(type, arrayPrefix)) {
            let subType: string = type.slice(arrayPrefix.length, -arraySuffix.length); // Array<Type> => Type
            let transformedData: any[] = [];
            for (let index = 0; index < data.length; index++) {
                let datum = data[index];
                transformedData.push(ObjectSerializer.deserialize(datum, subType));
            }
            return transformedData;
        } else if (startsWith(type, mapPrefix)) {
            let subType: string = type.slice(mapPrefix.length, -mapSuffix.length); // { [key: string]: Type; } => Type
            let transformedData: { [key: string]: any } = {};
            for (let key in data) {
                transformedData[key] = ObjectSerializer.deserialize(
                    data[key],
                    subType,
                );
            }
            return transformedData;
        } else if (type === "Date") {
            return new Date(data);
        } else {
            if (enumsMap[type]) {// is Enum
                return data;
            }

            if (!typeMap[type]) { // dont know the type
                return data;
            }
            let instance = new typeMap[type]();
            let attributeTypes = typeMap[type].getAttributeTypeMap();
            for (let index = 0; index < attributeTypes.length; index++) {
                let attributeType = attributeTypes[index];
                instance[attributeType.name] = ObjectSerializer.deserialize(data[attributeType.baseName], attributeType.type);
            }
            return instance;
        }
    }
}

export interface Authentication {
    /**
    * Apply authentication settings to header and query params.
    */
    applyToRequest(requestOptions: localVarRequest.Options): Promise<void> | void;
}

export class HttpBasicAuth implements Authentication {
    public username: string = '';
    public password: string = '';

    applyToRequest(requestOptions: localVarRequest.Options): void {
        requestOptions.auth = {
            username: this.username, password: this.password
        }
    }
}

export class HttpBearerAuth implements Authentication {
    public accessToken: string | (() => string) = '';

    applyToRequest(requestOptions: localVarRequest.Options): void {
        if (requestOptions && requestOptions.headers) {
            const accessToken = typeof this.accessToken === 'function'
                            ? this.accessToken()
                            : this.accessToken;
            requestOptions.headers["Authorization"] = "Bearer " + accessToken;
        }
    }
}

export class ApiKeyAuth implements Authentication {
    public apiKey: string = '';

    constructor(private location: string, private paramName: string) {
    }

    applyToRequest(requestOptions: localVarRequest.Options): void {
        if (this.location == "query") {
            (<any>requestOptions.qs)[this.paramName] = this.apiKey;
        } else if (this.location == "header" && requestOptions && requestOptions.headers) {
            requestOptions.headers[this.paramName] = this.apiKey;
        } else if (this.location == 'cookie' && requestOptions && requestOptions.headers) {
            if (requestOptions.headers['Cookie']) {
                requestOptions.headers['Cookie'] += '; ' + this.paramName + '=' + encodeURIComponent(this.apiKey);
            }
            else {
                requestOptions.headers['Cookie'] = this.paramName + '=' + encodeURIComponent(this.apiKey);
            }
        }
    }
}

export class OAuth implements Authentication {
    public accessToken: string = '';

    applyToRequest(requestOptions: localVarRequest.Options): void {
        if (requestOptions && requestOptions.headers) {
            requestOptions.headers["Authorization"] = "Bearer " + this.accessToken;
        }
    }
}

export class VoidAuth implements Authentication {
    public username: string = '';
    public password: string = '';

    applyToRequest(_: localVarRequest.Options): void {
        // Do nothing
    }
}

export type Interceptor = (requestOptions: localVarRequest.Options) => (Promise<void> | void);
