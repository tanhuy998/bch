import CandidateCRUDEndpoint from "../../api/candidateCRUD.api";
import CRUDEndpoint from "../../backend/crudEndpoint";
import EditSingleCandidateFormDelegator from "../valueObject/editSingleCandidateFormDelegator";

export default class EditSingleCandidateUseCase extends CandidateCRUDEndpoint {

    #formDelegator = new EditSingleCandidateFormDelegator();
    /**
     * @type {string}
     */
    #editedCandidateUUID;

    get formDelegator() {

        return this.#formDelegator;
    }

    get candidateUUID() {

        return this.#editedCandidateUUID
    }

    constructor() {

        super({})
    }

    fetchCandidate() {

        return super.read(this.#editedCandidateUUID);
    }

    /**
     * 
     * @param {string} candidateUUID 
     */
    setCandidateUUID(candidateUUID) {

        //this.formDelegator.setCandidateUUID(candidateUUID);

        this.#editedCandidateUUID = candidateUUID;
    }
}