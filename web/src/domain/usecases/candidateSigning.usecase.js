import { candidate_signing_family_t, candidate_signing_info_t, candidate_signing_politic_t } from "../models/candidate.model";
import CandidateEducationFormDelegator from "../../pages/candidateSinging/delegators/candidateEducationFormDelegator";
import CandidateJobFormDelegator from "../../pages/candidateSinging/delegators/candidateJobFormDelegator";
import CandidateIdentityFormDelegator from "../../pages/candidateSinging/delegators/candidateIdentityFormDelegator"
import EndpointFormDelegator from "../valueObject/endpointFormDelegator";
import HttpEndpoint from "../../backend/endpoint";
import CandidateSigningEndpoint from "../../api/candidateSigning.api";
import CandidateParentFormDelegator from "../../pages/candidateSinging/delegators/candidateFamilyFormDelegator";
import CandidatePoliticFormDelegator from "../../pages/candidateSinging/delegators/candidatePoliticFormDelegator";
import ErrorResponse from "../../backend/error/errorResponse";

export default class CandidateSigningUseCase extends CandidateSigningEndpoint {

    #dataModel;

    #educationFormDelegator = new CandidateEducationFormDelegator();
    #jobFormDelegator = new CandidateJobFormDelegator();
    #civilIdentityFormDelegator = new CandidateIdentityFormDelegator();
    #candidateMotherFormDelegator = new CandidateParentFormDelegator('mẹ');
    #candidateFatherFormDelegator = new CandidateParentFormDelegator('cha');
    #fatherPoliticHistoryFormDelegator = new CandidatePoliticFormDelegator();
    #motherPoliticHistoryFormDelegator = new CandidatePoliticFormDelegator();

    get dataModel() {

        this.#marshall();
        return this.#dataModel;
    }

    get candidateFatherPoliticHistoryFormDelegator() {

        return this.#fatherPoliticHistoryFormDelegator;
    }

    get candidateMotherPoliticHistoryFormDelegator() {

        return this.#motherPoliticHistoryFormDelegator;
    }

    get candidateFatherFormDelegator() {

        return this.#candidateFatherFormDelegator;
    }

    get candidateMotherFormDelegator() {

        return this.#candidateMotherFormDelegator;
    }

    get candidateIdentityFormDelegator() {

        return this.#civilIdentityFormDelegator;
    }

    get candidateJobFormDelegator() {

        return this.#jobFormDelegator;
    }

    get candidateEducationFormDelegator() {

        return this.#educationFormDelegator;
    }

    constructor() {

        super();
    }

    #marshall() {

        const dataModel = this.#dataModel = this.#jobFormDelegator.dataModel;

        dataModel.civilIdentity = this.#civilIdentityFormDelegator.dataModel;
        dataModel.education = this.#educationFormDelegator.dataModel;
        dataModel.politic = this.#fatherPoliticHistoryFormDelegator.dataModel;

        this.#initCandidateFamilyDataModel();
    }

    #initCandidateFamilyDataModel() {

        const candidateSigningDataModel = this.#dataModel;

        const family = candidateSigningDataModel.family = new candidate_signing_family_t();

        family.mother = this.#candidateMotherFormDelegator.dataModel;
        family.mother.politic ||= new candidate_signing_politic_t();
        family.mother.politic.history = this.#motherPoliticHistoryFormDelegator.dataModel;

        family.father = this.#candidateFatherFormDelegator.dataModel;
        family.father.politic ||= new candidate_signing_politic_t();
        family.father.politic.history = this.#fatherPoliticHistoryFormDelegator.dataModel;
    }

    async submit(campainUUID, candidateUUID) {

        try {
            
            this.#marshall();
            console.log('BEFORE ENDPOINT', this.#dataModel)
            console.log(JSON.stringify(this.#dataModel))
            await super.commit(campainUUID, candidateUUID, this.#dataModel);
        }
        catch (e) {

            if (e instanceof ErrorResponse) {

                const msg = await e.responseObject.text();

                alert(msg);
                return;
            }

            alert(e?.message || e);
        }
    }
}