/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    EntDisease,
    EntDiseaseFromJSON,
    EntDiseaseToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderToJSON,
    EntMedicalCare,
    EntMedicalCareFromJSON,
    EntMedicalCareToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientToJSON,
} from '../models';

export interface CreateDiseaseRequest {
    disease: EntDisease;
}

export interface CreateEmployeeRequest {
    employee: EntEmployee;
}

export interface CreateGenderRequest {
    gender: EntMedicalCare;
}

export interface CreateMedicalCareRequest {
    medicalCare: EntMedicalCare;
}

export interface CreatePatientRequest {
    patient: EntPatient;
}

export interface GetDiseaseRequest {
    id: number;
}

export interface GetEmployeeRequest {
    id: number;
}

export interface GetGenderRequest {
    id: number;
}

export interface GetMedicalCareRequest {
    id: number;
}

export interface GetPatientRequest {
    id: number;
}

export interface ListDiseaseRequest {
    limit?: number;
    offset?: number;
}

export interface ListEmployeeRequest {
    limit?: number;
    offset?: number;
}

export interface ListGenderRequest {
    limit?: number;
    offset?: number;
}

export interface ListMedicalCareRequest {
    limit?: number;
    offset?: number;
}

export interface ListPatientRequest {
    limit?: number;
    offset?: number;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create Disease
     * Create Disease
     */
    async createDiseaseRaw(requestParameters: CreateDiseaseRequest): Promise<runtime.ApiResponse<EntDisease>> {
        if (requestParameters.disease === null || requestParameters.disease === undefined) {
            throw new runtime.RequiredError('disease','Required parameter requestParameters.disease was null or undefined when calling createDisease.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/diseases`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntDiseaseToJSON(requestParameters.disease),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntDiseaseFromJSON(jsonValue));
    }

    /**
     * Create Disease
     * Create Disease
     */
    async createDisease(requestParameters: CreateDiseaseRequest): Promise<EntDisease> {
        const response = await this.createDiseaseRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create Employee
     * Create Employee
     */
    async createEmployeeRaw(requestParameters: CreateEmployeeRequest): Promise<runtime.ApiResponse<EntEmployee>> {
        if (requestParameters.employee === null || requestParameters.employee === undefined) {
            throw new runtime.RequiredError('employee','Required parameter requestParameters.employee was null or undefined when calling createEmployee.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/employees`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntEmployeeToJSON(requestParameters.employee),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEmployeeFromJSON(jsonValue));
    }

    /**
     * Create Employee
     * Create Employee
     */
    async createEmployee(requestParameters: CreateEmployeeRequest): Promise<EntEmployee> {
        const response = await this.createEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create Gender
     * Create Gender
     */
    async createGenderRaw(requestParameters: CreateGenderRequest): Promise<runtime.ApiResponse<EntGender>> {
        if (requestParameters.gender === null || requestParameters.gender === undefined) {
            throw new runtime.RequiredError('gender','Required parameter requestParameters.gender was null or undefined when calling createGender.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/genders`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntGenderToJSON(requestParameters.gender),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntGenderFromJSON(jsonValue));
    }

    /**
     * Create Gender
     * Create Gender
     */
    async createGender(requestParameters: CreateGenderRequest): Promise<EntGender> {
        const response = await this.createGenderRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create MedicalCare
     * Create MedicalCare
     */
    async createMedicalCareRaw(requestParameters: CreateMedicalCareRequest): Promise<runtime.ApiResponse<EntMedicalCare>> {
        if (requestParameters.medicalCare === null || requestParameters.medicalCare === undefined) {
            throw new runtime.RequiredError('medicalCare','Required parameter requestParameters.medicalCare was null or undefined when calling createMedicalCare.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/medicalcares`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntMedicalCareToJSON(requestParameters.medicalCare),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntMedicalCareFromJSON(jsonValue));
    }

    /**
     * Create MedicalCare
     * Create MedicalCare
     */
    async createMedicalCare(requestParameters: CreateMedicalCareRequest): Promise<EntMedicalCare> {
        const response = await this.createMedicalCareRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create Patient
     * Create Patient
     */
    async createPatientRaw(requestParameters: CreatePatientRequest): Promise<runtime.ApiResponse<EntPatient>> {
        if (requestParameters.patient === null || requestParameters.patient === undefined) {
            throw new runtime.RequiredError('patient','Required parameter requestParameters.patient was null or undefined when calling createPatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/patients`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntPatientToJSON(requestParameters.patient),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPatientFromJSON(jsonValue));
    }

    /**
     * Create Patient
     * Create Patient
     */
    async createPatient(requestParameters: CreatePatientRequest): Promise<EntPatient> {
        const response = await this.createPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * get Disease by ID
     * Get a Disease entity by ID
     */
    async getDiseaseRaw(requestParameters: GetDiseaseRequest): Promise<runtime.ApiResponse<EntDisease>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getDisease.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/diseases/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntDiseaseFromJSON(jsonValue));
    }

    /**
     * get Disease by ID
     * Get a Disease entity by ID
     */
    async getDisease(requestParameters: GetDiseaseRequest): Promise<EntDisease> {
        const response = await this.getDiseaseRaw(requestParameters);
        return await response.value();
    }

    /**
     * get Employee by ID
     * Get a Employee entity by ID
     */
    async getEmployeeRaw(requestParameters: GetEmployeeRequest): Promise<runtime.ApiResponse<EntEmployee>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getEmployee.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/employees/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEmployeeFromJSON(jsonValue));
    }

    /**
     * get Employee by ID
     * Get a Employee entity by ID
     */
    async getEmployee(requestParameters: GetEmployeeRequest): Promise<EntEmployee> {
        const response = await this.getEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * get Gender by ID
     * Get a Gender entity by ID
     */
    async getGenderRaw(requestParameters: GetGenderRequest): Promise<runtime.ApiResponse<EntGender>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getGender.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/genders/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntGenderFromJSON(jsonValue));
    }

    /**
     * get Gender by ID
     * Get a Gender entity by ID
     */
    async getGender(requestParameters: GetGenderRequest): Promise<EntGender> {
        const response = await this.getGenderRaw(requestParameters);
        return await response.value();
    }

    /**
     * get MedicalCare by ID
     * Get a MedicalCare entity by ID
     */
    async getMedicalCareRaw(requestParameters: GetMedicalCareRequest): Promise<runtime.ApiResponse<EntMedicalCare>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getMedicalCare.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/medicalcares/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntMedicalCareFromJSON(jsonValue));
    }

    /**
     * get MedicalCare by ID
     * Get a MedicalCare entity by ID
     */
    async getMedicalCare(requestParameters: GetMedicalCareRequest): Promise<EntMedicalCare> {
        const response = await this.getMedicalCareRaw(requestParameters);
        return await response.value();
    }

    /**
     * get Patient by ID
     * Get a Patient entity by ID
     */
    async getPatientRaw(requestParameters: GetPatientRequest): Promise<runtime.ApiResponse<EntPatient>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getPatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/patients/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPatientFromJSON(jsonValue));
    }

    /**
     * get Patient by ID
     * Get a Patient entity by ID
     */
    async getPatient(requestParameters: GetPatientRequest): Promise<EntPatient> {
        const response = await this.getPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * list Disease entities
     * List Disease entities
     */
    async listDiseaseRaw(requestParameters: ListDiseaseRequest): Promise<runtime.ApiResponse<Array<EntDisease>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/diseases`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntDiseaseFromJSON));
    }

    /**
     * list Disease entities
     * List Disease entities
     */
    async listDisease(requestParameters: ListDiseaseRequest): Promise<Array<EntDisease>> {
        const response = await this.listDiseaseRaw(requestParameters);
        return await response.value();
    }

    /**
     * list Employee entities
     * List Employee entities
     */
    async listEmployeeRaw(requestParameters: ListEmployeeRequest): Promise<runtime.ApiResponse<Array<EntEmployee>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/employees`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntEmployeeFromJSON));
    }

    /**
     * list Employee entities
     * List Employee entities
     */
    async listEmployee(requestParameters: ListEmployeeRequest): Promise<Array<EntEmployee>> {
        const response = await this.listEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * list Gender entities
     * List Gender entities
     */
    async listGenderRaw(requestParameters: ListGenderRequest): Promise<runtime.ApiResponse<Array<EntGender>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/genders`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntGenderFromJSON));
    }

    /**
     * list Gender entities
     * List Gender entities
     */
    async listGender(requestParameters: ListGenderRequest): Promise<Array<EntGender>> {
        const response = await this.listGenderRaw(requestParameters);
        return await response.value();
    }

    /**
     * list MedicalCare entities
     * List MedicalCare entities
     */
    async listMedicalCareRaw(requestParameters: ListMedicalCareRequest): Promise<runtime.ApiResponse<Array<EntMedicalCare>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/medicalcares`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntMedicalCareFromJSON));
    }

    /**
     * list MedicalCare entities
     * List MedicalCare entities
     */
    async listMedicalCare(requestParameters: ListMedicalCareRequest): Promise<Array<EntMedicalCare>> {
        const response = await this.listMedicalCareRaw(requestParameters);
        return await response.value();
    }

    /**
     * list Patient entities
     * List Patient entities
     */
    async listPatientRaw(requestParameters: ListPatientRequest): Promise<runtime.ApiResponse<Array<EntPatient>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/patients`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntPatientFromJSON));
    }

    /**
     * list Patient entities
     * List Patient entities
     */
    async listPatient(requestParameters: ListPatientRequest): Promise<Array<EntPatient>> {
        const response = await this.listPatientRaw(requestParameters);
        return await response.value();
    }

}
