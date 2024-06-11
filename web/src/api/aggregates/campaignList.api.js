import HttpEndpoint from "../../backend/endpoint";

const DEFAULT_PAGE_LIMIT = 3;


export default class CampaignList extends HttpEndpoint {

    constructor({scheme, host, uri} = {}) {

        uri = "campaigns"

        super({scheme, host, uri});
    }


    async fetch(query = {}) {

        
        const res =  await super.fetch(
            undefined,
            {
                p_pivot: query.p_pivot || undefined,
                p_limit: query.p_limit || DEFAULT_PAGE_LIMIT,
                p_prev: query.p_prev || false,
            },
        )

        return await res.json();
    }
}