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
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
import * as runtime from '../runtime';
import { DataSvcCreateObjectRequestToJSON, DataSvcCreateObjectResponseFromJSON, DataSvcDeleteObjectRequestToJSON, DataSvcQueryRequestToJSON, DataSvcQueryResponseFromJSON, DataSvcUpdateObjectsRequestToJSON, DataSvcUpsertObjectRequestToJSON, DataSvcUpsertObjectResponseFromJSON, } from '../models/index';
/**
 *
 */
export class DataSvcApi extends runtime.BaseAPI {
    /**
     * Creates a new object with the provided details. Requires authorization and user authentication.
     * Create a Generic Object
     */
    createObjectRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling createObject().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/data-svc/object`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: DataSvcCreateObjectRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => DataSvcCreateObjectResponseFromJSON(jsonValue));
        });
    }
    /**
     * Creates a new object with the provided details. Requires authorization and user authentication.
     * Create a Generic Object
     */
    createObject(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.createObjectRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Deletes all objects matchin the provided filters.
     * Delete Objects
     */
    deleteObjectsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling deleteObjects().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/data-svc/objects/delete`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: DataSvcDeleteObjectRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Deletes all objects matchin the provided filters.
     * Delete Objects
     */
    deleteObjects(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.deleteObjectsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieves objects from a specified table based on search criteria. Requires authorization and user authentication.   Use helper functions in your respective client library such as condition constructors (`equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
     * Query Objects
     */
    queryObjectsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/data-svc/objects`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: DataSvcQueryRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => DataSvcQueryResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieves objects from a specified table based on search criteria. Requires authorization and user authentication.   Use helper functions in your respective client library such as condition constructors (`equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
     * Query Objects
     */
    queryObjects() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.queryObjectsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Update fields of objects that match the given filters using the provided object. Any fields not included in the incoming object will remain unchanged.
     * Update Objects
     */
    updateObjectsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling updateObjects().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/data-svc/objects/update`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: DataSvcUpdateObjectsRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Update fields of objects that match the given filters using the provided object. Any fields not included in the incoming object will remain unchanged.
     * Update Objects
     */
    updateObjects(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.updateObjectsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Creates a new dynamic object or updates an existing one based on the provided data. Requires authorization and user authentication.
     * Upsert a Generic Object
     */
    upsertObjectRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['objectId'] == null) {
                throw new runtime.RequiredError('objectId', 'Required parameter "objectId" was null or undefined when calling upsertObject().');
            }
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling upsertObject().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/data-svc/object/{objectId}`.replace(`{${"objectId"}}`, encodeURIComponent(String(requestParameters['objectId']))),
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: DataSvcUpsertObjectRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => DataSvcUpsertObjectResponseFromJSON(jsonValue));
        });
    }
    /**
     * Creates a new dynamic object or updates an existing one based on the provided data. Requires authorization and user authentication.
     * Upsert a Generic Object
     */
    upsertObject(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.upsertObjectRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Upserts objects by ids.
     * Upsert Objects
     */
    upsertObjectsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling upsertObjects().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/data-svc/objects/upsert`,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: DataSvcUpsertObjectRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => DataSvcUpsertObjectResponseFromJSON(jsonValue));
        });
    }
    /**
     * Upserts objects by ids.
     * Upsert Objects
     */
    upsertObjects(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.upsertObjectsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
}
