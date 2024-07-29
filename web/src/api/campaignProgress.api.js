import AuthEndpoint from "../backend/autEndpoint";
import CampaignSignedCandidatesEndpoint from "./campaignSignedCandidates.api";
import CampaignCandidatesProgressEndpoint from "./campaignSignedCandidates.api";
import CampaignUnsignedCandidatesEndpoint from "./campaignUnSingedCandidates.api";

export default class CampaignProgressEndpoint extends AuthEndpoint {

    #campaignSignedCandidateEndpoint = new CampaignSignedCandidatesEndpoint();
    #campaignUnSignedCandidatesEndpoint = new CampaignUnsignedCandidatesEndpoint();

    get signedCandidates() {

        return this.#campaignSignedCandidateEndpoint;
    }

    get unSignedCandidates() {

        return this.#campaignUnSignedCandidatesEndpoint;
    }

    constructor() {

        super({uri: '/campaigns'})
    }

    // /**
    //  * 
    //  * @param {string} campaignUUID 
    //  */
    // fetch(campaignUUID) {
    //     console.log(campaignUUID)
    //     return super.fetch(
    //         undefined, undefined, `/${campaignUUID}/progress`
    //     );
    // }

    // async fetch(query = {}, extraURI) {

    //     return super.fetch(
    //         undefined, query, extraURI
    //     );
    // }

    async fetchReport(campaignUUID) {

        return super.fetch(
            undefined, undefined, `/${campaignUUID}/progress`
        );
    }
}