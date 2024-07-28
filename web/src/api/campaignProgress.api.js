import AuthEndpoint from "../backend/autEndpoint";
import CampaignCandidatesProgressEndpoint from "./campaignCandidateProgress.api";

export default class CampaignProgressEndpoint extends AuthEndpoint {

    #campaignCandidatesProgressEndpoint = new CampaignCandidatesProgressEndpoint();

    get candidate() {

        return this.#campaignCandidatesProgressEndpoint;
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
}