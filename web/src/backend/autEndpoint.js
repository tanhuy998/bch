import { getAccessToken, removeAccessToken, setAccessToken, setUserInfo } from "../lib/authSignature.lib";
import HttpEndpoint from "./endpoint"
import ErrorResponse from "./error/errorResponse";
import UserSessionExpireError from "./error/userSessionExpireError";

export default class AuthEndpoint extends HttpEndpoint {

    static async rotateSignatures() {

        const accessToken = getAccessToken();

        const res = await fetch(
            `${AuthEndpoint.baseURL}/auth/refresh`,
            {
                method: "POST",
                body: JSON.stringify(
                    {
                        data: accessToken,
                    }
                ),
            },
        );

        const status = res.status;

        if (status === 200) {

            const obj = res.json();

            setAccessToken(obj.data.accessToken);
            setUserInfo(obj.data.user);
            return;
        }

        if (status === 401) {

            removeAccessToken();
            throw new UserSessionExpireError();
        }

        throw new ErrorResponse(res);
    }

    async fetch(options = {}, query, extraURI) {    

        const res = this.#_fetch(...arguments);
        const status = res.status;

        if (status === 204) {

            return undefined;
        }

        if (status >= 200 && status < 300) {

            return res.json();
        }

        if (status === 401) {

            AuthEndpoint.rotateSignatures();
            return this.#_fetch(...arguments);
        }

        if (status >= 400) {

            throw new ErrorResponse(res);
        }   
    }

    /**
     * 
     * @param {*} options 
     * @param {*} query 
     * @param {*} extraURI 
     * @returns {Response}
     * 
     * @throws 
     */
    async #_fetch(options = {}, query, extraURI) {

        /** 
          *  in development, cors is setted to all( *), then authorization header will
          *  be ignored and causes cors issue. 
          *  
          *  authorizaration option just been turned on when the back end 's
          *  authorization bussiness completely done, in development
         */

        options.headers ||= {};
        options.headers['Authorization'] = `bearer ${getAccessToken() ?? ''}`;

        return super.fetchRaw(options, query, extraURI);        
    }
}