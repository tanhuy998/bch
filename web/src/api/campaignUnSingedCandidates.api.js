import AuthEndpoint from "../backend/autEndpoint";
import preprocessPaginationQuery from "./lib/preprocessPaginationQuery.lib";

export default class CampaignUnsignedCandidatesEndpoint extends AuthEndpoint {

    constructor() {

        super({ uri: '/candidates/unsigned/campaign' });
    }

    async getSigned(campaignUUID, query = {}) {

        return super.fetch(
            undefined, query, `/${campaignUUID}`
        )
    }

    async getUnSigned() {

    }

    /**
     * 
     * method for <PaginationTable /> to fetch data
     * 
     * @param {object} query 
     * @param {string} extraURI 
     * @returns 
     */
    async fetch(query = {}, extraURI) {

        return super.fetch(
            undefined, preprocessPaginationQuery(query), extraURI
        );
    }
}