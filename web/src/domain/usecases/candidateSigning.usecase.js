import { candidate_signing_info_t } from "../models/candidate.model";
import CandidateEducationFormDelegator from "../../pages/candidateSinging/delegators/candidateEducationFormDelegator";
import CandidateJobFormDelegator from "../../pages/candidateSinging/delegators/candidateJobFormDelegator";
import CandidateIdentityFormDelegator from "../../pages/candidateSinging/delegators/candidateIdentityFormDelegator"
import EndpointFormDelegator from "../valueObject/endpointFormDelegator";
import HttpEndpoint from "../../backend/endpoint";
import CandidateSigningEndpoint from "../../api/candidateSigning.api";

export default class CandidateSigningUseCase extends CandidateSigningEndpoint {

    #dataModel;

    #educationFormDelegator = new CandidateEducationFormDelegator();
    #jobFormDelegator = new CandidateJobFormDelegator();
    #civilIdentityFormDelegator = new CandidateIdentityFormDelegator()

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

        const dataModel = this.#dataModel = this.#jobFormDelegator.dataModel;

        dataModel.civilIdentity;
        dataModel.education = this.#educationFormDelegator.dataModel;
    }

    submit(campainUUID, candidateUUID) {

        try {

            super.commit(campainUUID, candidateUUID, this.#dataModel);
        }
        catch (e) {

            alert(e?.message || e);
        }
    }
}