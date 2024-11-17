import HttpEndpoint from "../../backend/endpoint";
import { fetch_options } from "../../domain/models/fetchOption.model";

/**
 * @typedef {tenant_t}
 * @property {string} name
 * @property {string} uuid
 * @property {string} isTenantAgent
 */

export default class SwitchTenantUseCase extends HttpEndpoint {

    constructor() {

        super({})
    }

    /**
     *  @returns {Promise<Array<tenant_t>>}
     */
    async fetchUserTenants() {

        const options = new fetch_options()
        options.method = "GET"

        return super.fetch(
            options,
            undefined,
            '/auth/gen/nav'
        )
    }

    /**
     * 
     * @param {string} tenantUUID 
     */
    async switchToTenant(tenantUUID) {

        return super.fetch(
            undefined,
            undefined,
            `/auth/signatures/tenant/${tenantUUID}`
        )
    }
}