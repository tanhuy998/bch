import HttpEndpoint from "../backend/endpoint";
import AuthEndpoint from "../backend/autEndpoint";

export default class CampaignCandidateEnpoint extends AuthEndpoint {

    #campaignUUID

    constructor(campaignUUID) {

        super({uri: "candidates"});

        this.#campaignUUID = campaignUUID;
    }
}