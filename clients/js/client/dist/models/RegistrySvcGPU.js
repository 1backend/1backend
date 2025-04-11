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
import { RegistrySvcProcessFromJSON, RegistrySvcProcessToJSON, } from './RegistrySvcProcess';
/**
 * Check if a given object implements the RegistrySvcGPU interface.
 */
export function instanceOfRegistrySvcGPU(value) {
    return true;
}
export function RegistrySvcGPUFromJSON(json) {
    return RegistrySvcGPUFromJSONTyped(json, false);
}
export function RegistrySvcGPUFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'busId': json['busId'] == null ? undefined : json['busId'],
        'computeMode': json['computeMode'] == null ? undefined : json['computeMode'],
        'cudaVersion': json['cudaVersion'] == null ? undefined : json['cudaVersion'],
        'gpuUtilization': json['gpuUtilization'] == null ? undefined : json['gpuUtilization'],
        'id': json['id'] == null ? undefined : json['id'],
        'intraNodeId': json['intraNodeId'] == null ? undefined : json['intraNodeId'],
        'memoryTotal': json['memoryTotal'] == null ? undefined : json['memoryTotal'],
        'memoryUsage': json['memoryUsage'] == null ? undefined : json['memoryUsage'],
        'name': json['name'] == null ? undefined : json['name'],
        'performanceState': json['performanceState'] == null ? undefined : json['performanceState'],
        'powerCap': json['powerCap'] == null ? undefined : json['powerCap'],
        'powerUsage': json['powerUsage'] == null ? undefined : json['powerUsage'],
        'processDetails': json['processDetails'] == null ? undefined : (json['processDetails'].map(RegistrySvcProcessFromJSON)),
        'temperature': json['temperature'] == null ? undefined : json['temperature'],
    };
}
export function RegistrySvcGPUToJSON(json) {
    return RegistrySvcGPUToJSONTyped(json, false);
}
export function RegistrySvcGPUToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'busId': value['busId'],
        'computeMode': value['computeMode'],
        'cudaVersion': value['cudaVersion'],
        'gpuUtilization': value['gpuUtilization'],
        'id': value['id'],
        'intraNodeId': value['intraNodeId'],
        'memoryTotal': value['memoryTotal'],
        'memoryUsage': value['memoryUsage'],
        'name': value['name'],
        'performanceState': value['performanceState'],
        'powerCap': value['powerCap'],
        'powerUsage': value['powerUsage'],
        'processDetails': value['processDetails'] == null ? undefined : (value['processDetails'].map(RegistrySvcProcessToJSON)),
        'temperature': value['temperature'],
    };
}
