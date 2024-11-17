import HttpEndpoint from "../../backend/endpoint";
import { fetch_options } from "../../domain/models/fetchOption.model";

export const FORM_DELEGATOR_ENDPOINT_ACTION = 'submit'

export default class LoginEndpoint extends HttpEndpoint {

    constructor() {

        super({
            uri: "/auth/gen"
        })
    }

    /**
     * 
     * @param {string} username 
     * @param {string} password 
     */
    submit(dataModel) {

        const options = new fetch_options({data: dataModel})
        options.method = "POST";

        return super.fetch(
            options, undefined, "/credentials"
        );
    }

    /**
     *  @returns {bool}
     */
    async isLoggedIn() {

        const options = new fetch_options()
        options.method = "HEAD"

        try {

            const res = await super.fetch(options);

            return true;
        }
        catch (e) {

            return false;
        }
    }
}