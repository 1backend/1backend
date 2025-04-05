/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { FirehoseSvcEvent } from './firehoseSvcEvent';

export class FirehoseSvcEventPublishRequest {
    'event'?: FirehoseSvcEvent;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "event",
            "baseName": "event",
            "type": "FirehoseSvcEvent"
        }    ];

    static getAttributeTypeMap() {
        return FirehoseSvcEventPublishRequest.attributeTypeMap;
    }
}

