import NewCandidateFormDelegator from "./newCandidateFormDelegator";

export default class EditSingleCampaignFormDelegator extends NewCandidateFormDelegator {

    /**@type {string} */
    #candidateUUID;

    /**
     * 
     * @param {string} uuid 
     */
    setCandidateUUID(uuid) {

        this.#candidateUUID = uuid;
    }

    shouldNavigate() {

        return -1;
    }

    async interceptSubmission() {

        try {
            console.log('submitted model', this.dataModel)
            const res = await this.endpoint.update(this.#candidateUUID, this.dataModel);

            this.reset();
        }
        catch (e) {
            console.log('error')
            super._handleError(e);
        }

    }
}