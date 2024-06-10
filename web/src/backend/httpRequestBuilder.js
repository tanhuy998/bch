/**
 * @typedef {import('./endpoint.js')} HttpEndpoint 
 */

export default class HttpRequestBuilder {

    /**@type {HttpEndpoint} */
    #endpoint;
    #obj = {
        method: 'GET',
    };


    /**
     * 
     * @param {HttpEndpoint} endpoint 
     */
    constructor(endpoint) {

        this.#endpoint = endpoint;
    }
    
    method(m) {

        this.#obj.method = m

        return this;    
    }

    body(content) {

        this.#obj.body = content
        
        return this;
    }

    /**
     * 
     * @param {Object} list 
     * @returns 
     */
    headers(list) {

        this.#obj.headers = list

        return this;
    }

    build() {

        return this.#obj
    }

    send() {

        this.#endpoint.fetch(this.#obj)
    }
}
