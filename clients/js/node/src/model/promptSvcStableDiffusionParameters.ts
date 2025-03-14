/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { StableDiffusionTxt2ImgRequest } from './stableDiffusionTxt2ImgRequest';

export class PromptSvcStableDiffusionParameters {
    /**
    * Text to image parameters
    */
    'txt2Img'?: StableDiffusionTxt2ImgRequest;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "txt2Img",
            "baseName": "txt2Img",
            "type": "StableDiffusionTxt2ImgRequest"
        }    ];

    static getAttributeTypeMap() {
        return PromptSvcStableDiffusionParameters.attributeTypeMap;
    }
}

