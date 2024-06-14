export default class ErrorResponse extends Error {

    #res

    get fetchResponse() {

        return this.#res
    }

    constructor(resObj) {

        super('fetch responses a error status');

        this.#res = resObj;
    }
}