import HttpEndpoint from "../backend/endpoint";
import AuthEndpoint from "../backend/autEndpoint";
import { DEFAULT_PAGINATION_LIMIT } from "./constant";
import preprocessPaginationQuery from "./lib/preprocessPaginationQuery.lib";


export default class CampaignListEndpoint extends AuthEndpoint {

    constructor({scheme, host, port} = {}) {

        super({scheme, host, uri: '/campaigns', port});
    }

    async fetch(query = {}) {

        
        const res = super.fetch(
            undefined,
            preprocessPaginationQuery(query),
        )

        return res;
    }
}