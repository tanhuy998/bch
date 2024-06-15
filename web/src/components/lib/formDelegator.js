export default class FormDelegator {

    get dataModel() {

        return undefined;
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