/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
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
import { DeploySvcDeleteDeploymentRequestToJSON, DeploySvcListDeploymentsResponseFromJSON, DeploySvcSaveDeploymentRequestToJSON, } from '../models/index';
/**
 *
 */
export class DeploySvcApi extends runtime.BaseAPI {
    /**
     * Delete a deployment.
     * Delete Deployment
     */
    deleteDeploymentRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/deploy-svc/deployment`,
                method: 'DELETE',
                headers: headerParameters,
                query: queryParameters,
                body: DeploySvcDeleteDeploymentRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Delete a deployment.
     * Delete Deployment
     */
    deleteDeployment() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.deleteDeploymentRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieve a list of deployments.
     * List Deployments
     */
    listDeploymentsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/deploy-svc/deployments`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: requestParameters['body'],
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => DeploySvcListDeploymentsResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieve a list of deployments.
     * List Deployments
     */
    listDeployments() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.listDeploymentsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Save a deployment.
     * Save Deployment
     */
    saveDeploymentRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/deploy-svc/deployment`,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: DeploySvcSaveDeploymentRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Save a deployment.
     * Save Deployment
     */
    saveDeployment() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.saveDeploymentRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
}
