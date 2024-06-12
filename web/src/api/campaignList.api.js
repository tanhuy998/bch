import HttpEndpoint from "../backend/endpoint";
import AuthEndpoint from "../backend/autEndpoint";
import { DEFAULT_PAGINATION_LIMIT } from "./constant";


export default class CampaignListEndpoint extends AuthEndpoint {

    constructor({scheme, host, port} = {}) {

        super({scheme, host, uri: '/campaigns', port});
    }

    async fetch(query = {}) {

        
        const res = super.fetch(
            undefined,
            {
                p_pivot: query.p_pivot || undefined,
                p_limit: query.p_limit || DEFAULT_PAGINATION_LIMIT,
                p_prev: query.p_prev || false,
            },
        )

        return res;
    }
}