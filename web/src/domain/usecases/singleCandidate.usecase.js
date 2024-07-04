import CandidateCRUDEndpoint from "../../api/candidateCRUD.api";
import CandidateSigningEndpoint from "../../api/candidateSigning.api";

export default class SingleCandidateUseCase extends CandidateCRUDEndpoint {

    #signingEnpoint = new CandidateSigningEndpoint();

    async isOpenSigning(campaignUUID, candidateUUID) {

        try {

            await this.#signingEnpoint.handShake(campaignUUID, candidateUUID);

            return true;
        } 
        catch (e) {

            return false;
        }
    }
}