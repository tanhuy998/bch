import CandidateCRUDEndpoint from "../../api/candidateCRUD.api";
import CRUDEndpoint from "../../backend/crudEndpoint";
import EditSingleCandidateFormDelegator from "../valueObject/editSingleCandidateFormDelegator";
import NewCandidateFormDelegator from "../valueObject/newCandidateFormDelegator";
import NewCampaignUseCase from "./newCampaign.usecase";

export default class EditSingleCampaignUseCase extends NewCampaignUseCase {

    //#formDelegator = new EditSingleCandidateFormDelegator();
    // /**
    //  * @type {string}
    //  */
    // #editedCandidateUUID;
    /**@type {string} */
    #campaignUUID;

    // get formDelegator() {

    //     return this.#formDelegator;
    // }

    get candidateUUID() {

        return this.#campaignUUID;
    }

    constructor() {

        super({})
    }

    fetchCampaign() {

        return super.endpoint.read(this.#campaignUUID);
    }

    /**
     * 
     * @param {string} candidateUUID 
     */
    setCampaignUUID(candidateUUID) {

        this.#campaignUUID = candidateUUID;
    }

    shouldNavigate() {

        return -1;
    }

    async interceptSubmission() {

        try {
            console.log('submitted model', this.dataModel)
            const res = await this.endpoint.update(this.#campaignUUID, this.dataModel);

            this.reset();
        }
        catch (e) {
            console.log('error')
            super._handleError(e);
        }

    }
}