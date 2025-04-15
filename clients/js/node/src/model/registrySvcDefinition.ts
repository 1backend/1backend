/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { RegistrySvcAPISpec } from './registrySvcAPISpec';
import { RegistrySvcClient } from './registrySvcClient';
import { RegistrySvcEnvVar } from './registrySvcEnvVar';
import { RegistrySvcImageSpec } from './registrySvcImageSpec';
import { RegistrySvcPortMapping } from './registrySvcPortMapping';
import { RegistrySvcRepositorySpec } from './registrySvcRepositorySpec';

export class RegistrySvcDefinition {
    /**
    * API Specs such as OpenAPI definitions etc.
    */
    'apiSpecs'?: Array<RegistrySvcAPISpec>;
    /**
    * Programming language clients such as on npm or GitHub.
    */
    'clients'?: Array<RegistrySvcClient>;
    /**
    * Envars is a map of Renvironment variables that a deployment (see Deploy Svc Deployment) of this definition will REQUIRE to run. E.g., {\"DB_URL\": \"mysql://user:password@host:port/db\"} These will be injected into the service instances (see Registry Svc Instance) at runtime. The value of a key here is the default value. The actual value can be overridden at deployment time.
    */
    'envars'?: Array<RegistrySvcEnvVar>;
    'id': string;
    /**
    * Container specifications for Docker, K8s, etc. Use this to deploy already built images.
    */
    'image'?: RegistrySvcImageSpec;
    /**
    * Ports have host ports and internal ports currently but they really only should have internal ports as host ports should be assigned by the system. Host ports might go away in the future.
    */
    'ports'?: Array<RegistrySvcPortMapping>;
    /**
    * Repository based definitions is an alternative to Image definitions. Instead of deploying an already built and pushed image, a source code repository url can be provided. The container will be built from the source.
    */
    'repository'?: RegistrySvcRepositorySpec;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "apiSpecs",
            "baseName": "apiSpecs",
            "type": "Array<RegistrySvcAPISpec>"
        },
        {
            "name": "clients",
            "baseName": "clients",
            "type": "Array<RegistrySvcClient>"
        },
        {
            "name": "envars",
            "baseName": "envars",
            "type": "Array<RegistrySvcEnvVar>"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "image",
            "baseName": "image",
            "type": "RegistrySvcImageSpec"
        },
        {
            "name": "ports",
            "baseName": "ports",
            "type": "Array<RegistrySvcPortMapping>"
        },
        {
            "name": "repository",
            "baseName": "repository",
            "type": "RegistrySvcRepositorySpec"
        }    ];

    static getAttributeTypeMap() {
        return RegistrySvcDefinition.attributeTypeMap;
    }
}

