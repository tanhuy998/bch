const REGEX_VALIDATE_NAME = /^([A-Za-z]+ )+/

export function required(val) {

    return Boolean(val) || false;
}


export default class Validator {

    #validatFuncitons;

    constructor(...validateFuncs) {

        this.#validatFuncitons = validateFuncs;
        this.#init()
    }

    #init() {

        for (const func of this.#validatFuncitons || []) {

            if (typeof func !== 'function') {

                throw new Error('invalid function passed to Validator');
            }
        }
    }

    validate(val) {

        for (const func of this.#validatFuncitons || []) {

            if (!Boolean(func(val))) {

                return false;
            }
        }

        return true;
    }
}