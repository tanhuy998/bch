import NewCandidateFormDelegator from "./newCandidateFormDelegator";

export default class EditSingleCandidateFormDelegator extends NewCandidateFormDelegator {

    async interceptSubmission() {

        try {

            const res = await this.endpoint.update(this.dataModel);

            this.reset();
        }
        catch (e) {

            super._handleError(e);
        }

    }
}