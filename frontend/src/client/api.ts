/* tslint:disable */
/* eslint-disable */
/**
 * Yadro Impulse Test Task
 * The goal of this task is to implement REST API and SPA to cover the described OpenAPI specification. Any Go library or frontend framework is acceptable. You may use https://editor.swagger.io to preview this spec.
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: edu@yadro.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, RawAxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError, operationServerMap } from './base';

/**
 * 
 * @export
 * @interface GetPet200Response
 */
export interface GetPet200Response {
    /**
     * Pet photo in ASCII art format
     * @type {string}
     * @memberof GetPet200Response
     */
    'ascii': string;
    /**
     * Description of pet on the photo
     * @type {string}
     * @memberof GetPet200Response
     */
    'description': string;
}
/**
 * 
 * @export
 * @interface GetPet500Response
 */
export interface GetPet500Response {
    /**
     * error message
     * @type {string}
     * @memberof GetPet500Response
     */
    'message': string;
}
/**
 * 
 * @export
 * @interface UploadPet400Response
 */
export interface UploadPet400Response {
    /**
     * error message
     * @type {string}
     * @memberof UploadPet400Response
     */
    'message'?: string | null;
}

/**
 * PetApi - axios parameter creator
 * @export
 */
export const PetApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Get existing pet
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPet: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/pet`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * Upload new pet
         * @param {GetPet200Response} petProperties Pet properties
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        uploadPet: async (petProperties: GetPet200Response, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'petProperties' is not null or undefined
            assertParamExists('uploadPet', 'petProperties', petProperties)
            const localVarPath = `/pet`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(petProperties, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * PetApi - functional programming interface
 * @export
 */
export const PetApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = PetApiAxiosParamCreator(configuration)
    return {
        /**
         * Get existing pet
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getPet(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetPet200Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getPet(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['PetApi.getPet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * Upload new pet
         * @param {GetPet200Response} petProperties Pet properties
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async uploadPet(petProperties: GetPet200Response, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.uploadPet(petProperties, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['PetApi.uploadPet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * PetApi - factory interface
 * @export
 */
export const PetApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = PetApiFp(configuration)
    return {
        /**
         * Get existing pet
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPet(options?: RawAxiosRequestConfig): AxiosPromise<GetPet200Response> {
            return localVarFp.getPet(options).then((request) => request(axios, basePath));
        },
        /**
         * Upload new pet
         * @param {GetPet200Response} petProperties Pet properties
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        uploadPet(petProperties: GetPet200Response, options?: RawAxiosRequestConfig): AxiosPromise<void> {
            return localVarFp.uploadPet(petProperties, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * PetApi - object-oriented interface
 * @export
 * @class PetApi
 * @extends {BaseAPI}
 */
export class PetApi extends BaseAPI {
    /**
     * Get existing pet
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PetApi
     */
    public getPet(options?: RawAxiosRequestConfig) {
        return PetApiFp(this.configuration).getPet(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Upload new pet
     * @param {GetPet200Response} petProperties Pet properties
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PetApi
     */
    public uploadPet(petProperties: GetPet200Response, options?: RawAxiosRequestConfig) {
        return PetApiFp(this.configuration).uploadPet(petProperties, options).then((request) => request(this.axios, this.basePath));
    }
}



