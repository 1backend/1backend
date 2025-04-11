/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export enum PromptSvcPromptStatus {
    PromptStatusScheduled = <any> 'scheduled',
    PromptStatusRunning = <any> 'running',
    PromptStatusCompleted = <any> 'completed',
    PromptStatusErrored = <any> 'errored',
    PromptStatusAbandoned = <any> 'abandoned',
    PromptStatusCanceled = <any> 'canceled'
}
