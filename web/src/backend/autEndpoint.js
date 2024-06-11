import HttpEndpoint from "./endpoint"

export default class AuthEndpoint extends HttpEndpoint {

    #authStore;

    #accessToken;
    #refreshToken;

    setAuthStore(store) {

        this.#authStore = store;
    }

    async fetch(options = {}, query) {

        /** 
         *  in development, cors is setted to all( *), then authorization header will
         *  be ignored and causes cors issue. 
         *  
         *  authorizaration option just been turned on when the back end 's
         *  authorization bussiness completely done, in development
         */

        options.headers ||= {};
        options.headers['Authorization'] = `bearer ${this.#accessToken || ''}`;

        return super.fetch(options, query)
    }

    #prepareToken() {


    }

    #refresh() {


    }
}