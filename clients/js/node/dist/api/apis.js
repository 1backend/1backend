export * from './chatSvcApi';
import { ChatSvcApi } from './chatSvcApi';
export * from './configSvcApi';
import { ConfigSvcApi } from './configSvcApi';
export * from './containerSvcApi';
import { ContainerSvcApi } from './containerSvcApi';
export * from './dataSvcApi';
import { DataSvcApi } from './dataSvcApi';
export * from './deploySvcApi';
import { DeploySvcApi } from './deploySvcApi';
export * from './emailSvcApi';
import { EmailSvcApi } from './emailSvcApi';
export * from './fileSvcApi';
import { FileSvcApi } from './fileSvcApi';
export * from './firehoseSvcApi';
import { FirehoseSvcApi } from './firehoseSvcApi';
export * from './modelSvcApi';
import { ModelSvcApi } from './modelSvcApi';
export * from './policySvcApi';
import { PolicySvcApi } from './policySvcApi';
export * from './promptSvcApi';
import { PromptSvcApi } from './promptSvcApi';
export * from './registrySvcApi';
import { RegistrySvcApi } from './registrySvcApi';
export * from './secretSvcApi';
import { SecretSvcApi } from './secretSvcApi';
export * from './sourceSvcApi';
import { SourceSvcApi } from './sourceSvcApi';
export * from './userSvcApi';
import { UserSvcApi } from './userSvcApi';
export class HttpError extends Error {
    constructor(response, body, statusCode) {
        super('HTTP request failed');
        this.response = response;
        this.body = body;
        this.statusCode = statusCode;
        this.name = 'HttpError';
    }
}
export const APIS = [ChatSvcApi, ConfigSvcApi, ContainerSvcApi, DataSvcApi, DeploySvcApi, EmailSvcApi, FileSvcApi, FirehoseSvcApi, ModelSvcApi, PolicySvcApi, PromptSvcApi, RegistrySvcApi, SecretSvcApi, SourceSvcApi, UserSvcApi];
