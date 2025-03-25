/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
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
export declare const PromptSvcPromptType: {
    readonly PromptTypeImageTextToText: "Image-Text-to-Text";
    readonly PromptTypeVisualQuestionAnswering: "Visual Question Answering";
    readonly PromptTypeDocumentQuestionAnswering: "Document Question Answering";
    readonly PromptTypeTextToImage: "Text-to-Image";
    readonly PromptTypeImageToImage: "Image-to-Image";
    readonly PromptTypeImageToVideo: "Image-to-Video";
    readonly PromptTypeUnconditionalImageGeneration: "Unconditional Image Generation";
    readonly PromptTypeTextToVideo: "Text-to-Video";
    readonly PromptTypeZeroShotImageClassification: "Zero-Shot Image Classification";
    readonly PromptTypeZeroShotObjectDetection: "Zero-Shot Object Detection";
    readonly PromptTypeTextTo3D: "Text-to-3D";
    readonly PromptTypeImageTo3D: "Image-to-3D";
    readonly PromptTypeImageFeatureExtraction: "Image Feature Extraction";
    readonly PromptTypeKeypointDetection: "Keypoint Detection";
    readonly PromptTypeTextToText: "Text-to-Text";
    readonly PromptTypeQuestionAnswering: "Question Answering";
    readonly PromptTypeTranslation: "Translation";
    readonly PromptTypeSummarization: "Summarization";
    readonly PromptTypeTextGeneration: "Text Generation";
    readonly PromptTypeFillMask: "Fill-Mask";
    readonly PromptTypeTextToSpeech: "Text-to-Speech";
    readonly PromptTypeTextToAudio: "Text-to-Audio";
    readonly PromptTypeAutomaticSpeechRecognition: "Automatic Speech Recognition";
    readonly PromptTypeAudioToAudio: "Audio-to-Audio";
    readonly PromptTypeAudioClassification: "Audio Classification";
    readonly PromptTypeReinforcementLearning: "Reinforcement Learning";
    readonly PromptTypeRobotics: "Robotics";
    readonly PromptTypeGraphMachineLearning: "Graph Machine Learning";
};
export type PromptSvcPromptType = typeof PromptSvcPromptType[keyof typeof PromptSvcPromptType];
export declare function instanceOfPromptSvcPromptType(value: any): boolean;
export declare function PromptSvcPromptTypeFromJSON(json: any): PromptSvcPromptType;
export declare function PromptSvcPromptTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcPromptType;
export declare function PromptSvcPromptTypeToJSON(value?: PromptSvcPromptType | null): any;
export declare function PromptSvcPromptTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): PromptSvcPromptType;
