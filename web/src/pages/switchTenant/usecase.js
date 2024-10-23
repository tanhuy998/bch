import HttpEndpoint from "../../backend/endpoint";
import { fetch_options } from "../../domain/models/fetchOption.model";

export default class SwitchTenantUseCase extends HttpEndpoint {

    constructor() {

        super({
            uri: '/auth/nav'
        })
    }

    /**
     *  @returns {Promise<Array<tenant_t>>}
     */
    async fetchUserTenants() {

        const options = new fetch_options()
        options.method = "GET"

        return super.fetch(
            options
        )
    }
}