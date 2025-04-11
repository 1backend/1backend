/* tslint:disable */
/* eslint-disable */
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


/**
 * 
 * @export
 */
export const PromptSvcPromptType = {
    PromptTypeImageTextToText: 'Image-Text-to-Text',
    PromptTypeVisualQuestionAnswering: 'Visual Question Answering',
    PromptTypeDocumentQuestionAnswering: 'Document Question Answering',
    PromptTypeTextToImage: 'Text-to-Image',
    PromptTypeImageToImage: 'Image-to-Image',
    PromptTypeImageToVideo: 'Image-to-Video',
    PromptTypeUnconditionalImageGeneration: 'Unconditional Image Generation',
    PromptTypeTextToVideo: 'Text-to-Video',
    PromptTypeZeroShotImageClassification: 'Zero-Shot Image Classification',
    PromptTypeZeroShotObjectDetection: 'Zero-Shot Object Detection',
    PromptTypeTextTo3D: 'Text-to-3D',
    PromptTypeImageTo3D: 'Image-to-3D',
    PromptTypeImageFeatureExtraction: 'Image Feature Extraction',
    PromptTypeKeypointDetection: 'Keypoint Detection',
    PromptTypeTextToText: 'Text-to-Text',
    PromptTypeQuestionAnswering: 'Question Answering',
    PromptTypeTranslation: 'Translation',
    PromptTypeSummarization: 'Summarization',
    PromptTypeTextGeneration: 'Text Generation',
    PromptTypeFillMask: 'Fill-Mask',
    PromptTypeTextToSpeech: 'Text-to-Speech',
    PromptTypeTextToAudio: 'Text-to-Audio',
    PromptTypeAutomaticSpeechRecognition: 'Automatic Speech Recognition',
    PromptTypeAudioToAudio: 'Audio-to-Audio',
    PromptTypeAudioClassification: 'Audio Classification',
    PromptTypeReinforcementLearning: 'Reinforcement Learning',
    PromptTypeRobotics: 'Robotics',
    PromptTypeGraphMachineLearning: 'Graph Machine Learning'
} as const;
export type PromptSvcPromptType = typeof PromptSvcPromptType[keyof typeof PromptSvcPromptType];


export function instanceOfPromptSvcPromptType(value: any): boolean {
    for (const key in PromptSvcPromptType) {
        if (Object.prototype.hasOwnProperty.call(PromptSvcPromptType, key)) {
            if (PromptSvcPromptType[key as keyof typeof PromptSvcPromptType] === value) {
                return true;
            }
        }
    }
    return false;
}

export function PromptSvcPromptTypeFromJSON(json: any): PromptSvcPromptType {
    return PromptSvcPromptTypeFromJSONTyped(json, false);
}

export function PromptSvcPromptTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcPromptType {
    return json as PromptSvcPromptType;
}

export function PromptSvcPromptTypeToJSON(value?: PromptSvcPromptType | null): any {
    return value as any;
}

export function PromptSvcPromptTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): PromptSvcPromptType {
    return value as PromptSvcPromptType;
}

