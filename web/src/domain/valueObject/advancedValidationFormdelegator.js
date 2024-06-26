import Schema from "validate";
import ErrorTraceFormDelegator from "./errorTraceFormDelegator";

export default class AdvanceValidationFormDelegator extends ErrorTraceFormDelegator {

    #validationFailedFootprint;
    #validator;

    // get validationFailedFootPrint() {

    //     return this.#validationFailedFootprint;
    // }

    /**
     * @type {Schema}
     */
    get validator() {


    }

    get validationFailedFootPrint() {

        return this.#validationFailedFootprint;
    }

    /**
     * 
     * @param {Object} rules 
     */
    setValidatorRules(rules) {

        if (typeof rules !== 'object') {

            throw new Error('validator rules must be object that represent the package validator\'s rules');
        }

        this.#validator = new Schema(rules);
    }

    validateModel() {

        const errors = this.validator.validate(this.dataModel);

        if (
            !errors ||
             Array.isArray(errors) && errors.length === 0
        ) {

            return true;
        }

        this.setValidationFailedFootPrint(errors);
    }


    setValidationFailedFootPrint(any) {

        this.#validationFailedFootprint = any;
    }

    /**
     * @override
     */
    onValidationFailed() {

        const errors = this.#validationFailedFootprint;
        console.log('validate errors', errors)
        if (!errors) {

            alert('form validation failed')
            return;
        }

        if (Array.isArray(errors)) {

            const msg = errors.map(
                err => err.message
            ).join("\n");

            alert(msg);
            return
        }

        alert(errors?.message || errors);
    }
}