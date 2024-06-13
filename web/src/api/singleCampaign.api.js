import AuthEndpoint from "../backend/autEndpoint";

export default class SingleCampaignEndPoint extends AuthEndpoint {

    constructor({scheme, host, port} = {}) {

        super({uri: '/campaigns', scheme, host, port})
    }

    async fetch(campaignUUID) {

        return super.fetch(undefined, undefined, '/' + campaignUUID)
    }
}