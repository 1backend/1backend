/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import http from 'http';
import { FirehoseSvcEventPublishRequest } from '../model/firehoseSvcEventPublishRequest';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum FirehoseSvcApiApiKeys {
    BearerAuth = 0
}
export declare class FirehoseSvcApi {
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
    setApiKey(key: FirehoseSvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Publishes an event to the firehose service after authorization check
     * @summary Publish an Event
     * @param event Event to publish
     */
    publishEvent(event: FirehoseSvcEventPublishRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body?: any;
    }>;
    /**
     * Establish a subscription to the firehose events and accept a real time stream of them.
     * @summary Subscribe to the Event Stream
     */
    subscribeToEvents(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: string;
    }>;
}
