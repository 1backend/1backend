/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export enum RegistrySvcInstanceStatus {
    InstanceStatusUnknown = <any> 'Unknown',
    InstanceStatusHealthy = <any> 'Healthy',
    InstanceStatusDegraded = <any> 'Degraded',
    InstanceStatusUnreachable = <any> 'Unreachable',
    InstanceStatusError = <any> 'Error'
}
