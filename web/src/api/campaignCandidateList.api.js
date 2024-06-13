import AuthEndpoint from "../backend/autEndpoint";
import preprocessPaginationQuery from "./lib/preprocessPaginationQuery.lib";

export default class CampaignCandidateListEndpoint extends AuthEndpoint {

    constructor({host, scheme, port } = {}) {

        super({host, scheme, port, uri: '/candidates/campaign'});
    }

    async fetch(query = {}, campaignUUID) {

        const res = await super.fetch(
            undefined, preprocessPaginationQuery(query), '/' + campaignUUID
        );

        console.log('api', res)

        return res;
    }
}