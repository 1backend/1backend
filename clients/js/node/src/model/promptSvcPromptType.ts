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

export enum PromptSvcPromptType {
    PromptTypeImageTextToText = <any> 'Image-Text-to-Text',
    PromptTypeVisualQuestionAnswering = <any> 'Visual Question Answering',
    PromptTypeDocumentQuestionAnswering = <any> 'Document Question Answering',
    PromptTypeTextToImage = <any> 'Text-to-Image',
    PromptTypeImageToImage = <any> 'Image-to-Image',
    PromptTypeImageToVideo = <any> 'Image-to-Video',
    PromptTypeUnconditionalImageGeneration = <any> 'Unconditional Image Generation',
    PromptTypeTextToVideo = <any> 'Text-to-Video',
    PromptTypeZeroShotImageClassification = <any> 'Zero-Shot Image Classification',
    PromptTypeZeroShotObjectDetection = <any> 'Zero-Shot Object Detection',
    PromptTypeTextTo3D = <any> 'Text-to-3D',
    PromptTypeImageTo3D = <any> 'Image-to-3D',
    PromptTypeImageFeatureExtraction = <any> 'Image Feature Extraction',
    PromptTypeKeypointDetection = <any> 'Keypoint Detection',
    PromptTypeTextToText = <any> 'Text-to-Text',
    PromptTypeQuestionAnswering = <any> 'Question Answering',
    PromptTypeTranslation = <any> 'Translation',
    PromptTypeSummarization = <any> 'Summarization',
    PromptTypeTextGeneration = <any> 'Text Generation',
    PromptTypeFillMask = <any> 'Fill-Mask',
    PromptTypeTextToSpeech = <any> 'Text-to-Speech',
    PromptTypeTextToAudio = <any> 'Text-to-Audio',
    PromptTypeAutomaticSpeechRecognition = <any> 'Automatic Speech Recognition',
    PromptTypeAudioToAudio = <any> 'Audio-to-Audio',
    PromptTypeAudioClassification = <any> 'Audio Classification',
    PromptTypeReinforcementLearning = <any> 'Reinforcement Learning',
    PromptTypeRobotics = <any> 'Robotics',
    PromptTypeGraphMachineLearning = <any> 'Graph Machine Learning'
}
