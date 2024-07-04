import HttpEndpoint from "../backend/endpoint";
import { candidate_signing_info_t } from "../domain/models/candidate.model";

export class CandidateSigningInfoNotFoundError extends Error { }

export default class CandidateSigningEndpoint extends HttpEndpoint {

    constructor() {

        super({uri: '/signing'});
    }

    /**
     * 
     * @param {candidate_signing_info_t} signingInfo 
     */
    async commit(campaignUUID, candidateUUID, signingInfo) {
        
        return super.fetch(
            {
                method: "PATCH",
                body: JSON.stringify({
                    data: signingInfo,
                })
            },
            undefined,
            `/campaign/${campaignUUID}/candidate/${candidateUUID}`,
        )
    }

    async handShake(campaignUUID, candidateUUID) {

        const res = await super.fetchRaw(
            {
                method: "HEAD",
            },
            undefined,
            `/pending/campaign/${campaignUUID}/candidate/${candidateUUID}`
        );
        
        if (res.status !== 204) {

            throw new CandidateSigningInfoNotFoundError();
        }
    }
}

