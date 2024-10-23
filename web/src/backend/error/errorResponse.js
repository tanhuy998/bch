export default class ErrorResponse extends Error {

    /**
     * @type {Response}
     */
    #res

    /**
     * @type {Response}
     */
    get responseObject() {

        return this.#res
    }

    /**
     * 
     * @param {Response} resObj 
     */
    constructor(resObj) {

        super('fetch responses a error status');

        this.#res = resObj;
    }
}