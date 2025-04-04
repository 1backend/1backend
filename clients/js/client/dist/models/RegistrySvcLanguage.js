/* tslint:disable */
/* eslint-disable */
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
/**
 *
 * @export
 */
export const RegistrySvcLanguage = {
    JavaScript: 'JavaScript',
    Python: 'Python',
    Java: 'Java',
    CSharp: 'C#',
    CPlusPlus: 'C++',
    Ruby: 'Ruby',
    Go: 'Go',
    Swift: 'Swift',
    PHP: 'PHP',
    TypeScript: 'TypeScript',
    Kotlin: 'Kotlin',
    Scala: 'Scala',
    Perl: 'Perl',
    Rust: 'Rust',
    Haskell: 'Haskell',
    Clojure: 'Clojure',
    Elixir: 'Elixir',
    ObjectiveC: 'Objective-C',
    FSharp: 'F#'
};
export function instanceOfRegistrySvcLanguage(value) {
    for (const key in RegistrySvcLanguage) {
        if (Object.prototype.hasOwnProperty.call(RegistrySvcLanguage, key)) {
            if (RegistrySvcLanguage[key] === value) {
                return true;
            }
        }
    }
    return false;
}
export function RegistrySvcLanguageFromJSON(json) {
    return RegistrySvcLanguageFromJSONTyped(json, false);
}
export function RegistrySvcLanguageFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
export function RegistrySvcLanguageToJSON(value) {
    return value;
}
export function RegistrySvcLanguageToJSONTyped(value, ignoreDiscriminator) {
    return value;
}
