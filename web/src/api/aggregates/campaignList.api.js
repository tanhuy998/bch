import HttpEndpoint from "../../backend/endpoint";

const DEFAULT_PAGE_LIMIT = 3;


export default class CampaignList extends HttpEndpoint {

    constructor({scheme, host, uri} = {}) {

        uri = "campaigns"

        super({scheme, host, uri});
    }


    async fetch(pivot, pageLimit = DEFAULT_PAGE_LIMIT, isPrev) {

        const res =  await super.fetch(
            undefined,
            {
                p_pivot: pivot,
                p_limit: pageLimit,
                p_isPrev: isPrev,
            }
        )

        return await res.json();
    }
}