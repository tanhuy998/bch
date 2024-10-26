import ErrorResponse from "./error/errorResponse";
import HttpRequestBuilder from "./httpRequestBuilder";
import MockEndpoint from "./mockEndpoint";

const DEFAULT_HOST = 'localhost';
const DEFAULT_SCHEME = 'http';
const DEFAULT_URI = '';
const DEFAULT_PORT = 8000;

const REGEX_DOUBLE_SLASH = /\/\//


export default class HttpEndpoint extends MockEndpoint{

    static #mock = false;

    static useMock() {

        this.#mock = true;
    }

    /**
     * @return {string}
     */
    static get baseURL() {
        
        return `${DEFAULT_SCHEME}://${DEFAULT_HOST}${typeof DEFAULT_PORT === 'number' ? ":" + DEFAULT_PORT : ""}`
    }

    /**@type {String} */
    #host;
    #scheme;
    #uri;
    #port;

    #initUrl;

    #defaultRequestOptions = {
        mode: 'cors',
    }

    #defaultRequestHeaders = {
        //"Content-Type": "application/json;charset=UTF-8",
    }

    get url() {

        return this.#initUrl;
    }

    get baseURL() {

        return `${this.#scheme}://${this.#host}${typeof this.#port === 'number' ? ":" + this.#port : ""}`
    }

    constructor({ scheme, host, port, uri } = {}) {

        super();

        this.#port = port || DEFAULT_PORT;
        this.#host = host || DEFAULT_HOST;
        this.#scheme = scheme || DEFAULT_SCHEME;
        this.#uri = uri || DEFAULT_URI;

        this.#initUrl = `${this.#scheme}://${this.#host}${typeof this.#port === 'number' ? ":" + this.#port : ""}${this.#uri}`;
    }

    prepareOptions() {

        return new HttpRequestBuilder(this);
    }

    async fetch(options = {}, query, extraURI) {

        // options.mode = 'cors';

        // extraURI = typeof extraURI === 'string' && extraURI !== '' ? extraURI : '';

        // const queryStr = typeof query === 'object' ? '?' + new URLSearchParams(query) : '';


        // const res = await fetch(
        //     this.#initUrl + extraURI + queryStr,
        //     {
        //         ...this.#defaultRequestOptions,
        //         ...options,
        //         headers: {
        //             ...this.#defaultRequestHeaders,
        //             ...options.headers,
        //         }
        //     }
        // )

        const res = await this.fetchRaw(options, query, extraURI);

        if (!res.ok) {

            throw new ErrorResponse(res);
        }

        if (res.status === 204) {
            
            return;
        }

        return res.json();
    }

    async fetchRaw(options = {}, query, extraURI) {

        options.mode = 'cors';

        extraURI = typeof extraURI === 'string' && extraURI !== '' ? extraURI : '';

        const queryStr = typeof query === 'object' ? '?' + new URLSearchParams(query) : '';


        return fetch(
            this.#initUrl + extraURI + queryStr,
            {
                ...this.#defaultRequestOptions,
                ...options,
                headers: {
                    ...this.#defaultRequestHeaders,
                    ...options.headers,
                },
                credentials: 'include',
            }
        )
    }

    #fetchMock() {


    }
}
