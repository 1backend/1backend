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
export declare const RegistrySvcLanguage: {
    readonly JavaScript: "JavaScript";
    readonly Python: "Python";
    readonly Java: "Java";
    readonly CSharp: "C#";
    readonly CPlusPlus: "C++";
    readonly Ruby: "Ruby";
    readonly Go: "Go";
    readonly Swift: "Swift";
    readonly PHP: "PHP";
    readonly TypeScript: "TypeScript";
    readonly Kotlin: "Kotlin";
    readonly Scala: "Scala";
    readonly Perl: "Perl";
    readonly Rust: "Rust";
    readonly Haskell: "Haskell";
    readonly Clojure: "Clojure";
    readonly Elixir: "Elixir";
    readonly ObjectiveC: "Objective-C";
    readonly FSharp: "F#";
};
export type RegistrySvcLanguage = typeof RegistrySvcLanguage[keyof typeof RegistrySvcLanguage];
export declare function instanceOfRegistrySvcLanguage(value: any): boolean;
export declare function RegistrySvcLanguageFromJSON(json: any): RegistrySvcLanguage;
export declare function RegistrySvcLanguageFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcLanguage;
export declare function RegistrySvcLanguageToJSON(value?: RegistrySvcLanguage | null): any;
export declare function RegistrySvcLanguageToJSONTyped(value: any, ignoreDiscriminator: boolean): RegistrySvcLanguage;
