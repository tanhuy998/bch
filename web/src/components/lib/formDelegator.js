import { useNavigate } from "react-router-dom";

export default class FormDelegator {

    /**@type {Function} */
    #navigator;


    get dataModel() {

        return undefined;
    }

    reset() {
        
    }

    setNavigator(hook) {

        if (typeof hook !== 'function') {

            throw new Error('FormDelegator navigator must be hook function');
        }

        this.#navigator = hook;
    }

    navigate(path) {

        this.#navigator.call(undefined, path, {replace: true});   
    }

    /**
     * 
     * @returns {string?}
     */
    shouldNavigate() {

        return undefined;
    }

    onSuccess() {


    }

    /**
     * if dataModel evaluated, formData will be the dataModel.
     * Otherwise, formData will be the DOM form object.
     * 
     * @param {any} formData
     * 
     * @returns {boolean}
     */
    interceptSubmission() {


    }

    /**
     * if dataModel evaluated, formData will be the dataModel.
     * Otherwise, formData will be the DOM form object.
     * 
     * @param {any} formData
     * 
     * @returns {boolean}
     */
    validateModel(formData) {

        return true;
    }

    onValidationFailed() {


    }

    validateEveryInput(formData) {

        return true;
    }
}