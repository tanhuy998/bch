import AuthEndpoint from "../backend/autEndpoint";

export default class CampaignCandidateListEndpoing extends AuthEndpoint {

    constructor({host, scheme, port }) {

        super({host, scheme, port, uri: '/candidates/canpaigns'});
    }

    async fetch(campaignUUID, query = {}) {

        return super.fetch(undefined, query, '/' + campaignUUID);
    }
}