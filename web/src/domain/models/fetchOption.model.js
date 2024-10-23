export {
    fetch_options_t as fetch_options
}

function fetch_options_t(body) {

    if (typeof body === 'object') {

        body = JSON.stringify(body);
    }

    /**
     * @type {Object}
     */
    this.body = body

    /**
     * @type {string}
     */
    this.method = "GET"
}