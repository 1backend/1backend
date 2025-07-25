/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  FirehoseSvcErrorResponse,
  FirehoseSvcEventPublishRequest,
} from '../models/index';
import {
    FirehoseSvcErrorResponseFromJSON,
    FirehoseSvcErrorResponseToJSON,
    FirehoseSvcEventPublishRequestFromJSON,
    FirehoseSvcEventPublishRequestToJSON,
} from '../models/index';

export interface PublishEventRequest {
    event: FirehoseSvcEventPublishRequest;
}

/**
 * 
 */
export class FirehoseSvcApi extends runtime.BaseAPI {

    /**
     * Publishes an event to the firehose service after authorization check
     * Publish an Event
     */
    async publishEventRaw(requestParameters: PublishEventRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['event'] == null) {
            throw new runtime.RequiredError(
                'event',
                'Required parameter "event" was null or undefined when calling publishEvent().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }


        let urlPath = `/firehose-svc/event`;

        const response = await this.request({
            path: urlPath,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: FirehoseSvcEventPublishRequestToJSON(requestParameters['event']),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Publishes an event to the firehose service after authorization check
     * Publish an Event
     */
    async publishEvent(requestParameters: PublishEventRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.publishEventRaw(requestParameters, initOverrides);
    }

    /**
     * Establish a subscription to the firehose events and accept a real time stream of them.
     * Subscribe to the Event Stream
     */
    async subscribeToEventsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<string>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }


        let urlPath = `/firehose-svc/events/subscribe`;

        const response = await this.request({
            path: urlPath,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        if (this.isJsonMime(response.headers.get('content-type'))) {
            return new runtime.JSONApiResponse<string>(response);
        } else {
            return new runtime.TextApiResponse(response) as any;
        }
    }

    /**
     * Establish a subscription to the firehose events and accept a real time stream of them.
     * Subscribe to the Event Stream
     */
    async subscribeToEvents(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<string> {
        const response = await this.subscribeToEventsRaw(initOverrides);
        return await response.value();
    }

}
