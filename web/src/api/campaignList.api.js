import HttpEndpoint from "../backend/endpoint";
import AuthEndpoint from "../backend/autEndpoint";

const DEFAULT_PAGE_LIMIT = 3;


export default class CampaignListEndpoint extends AuthEndpoint {

    constructor({scheme, host, uri} = {}) {

        uri = "campaigns"

        super({scheme, host, uri});
    }

    async fetch(query = {}) {

        
        const res = super.fetch(
            undefined,
            {
                p_pivot: query.p_pivot || undefined,
                p_limit: query.p_limit || DEFAULT_PAGE_LIMIT,
                p_prev: query.p_prev || false,
            },
        )

        return res;
    }
}