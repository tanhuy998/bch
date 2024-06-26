import AdvanceValidationFormDelegator from "./advancedValidationFormdelegator";

export default class CollectableFormDelegator extends AdvanceValidationFormDelegator {

    // interceptSubmission() {
       
    //     const isValid = super.validateModel();

    //     if (!isValid) {

    //         super.onValidationFailed();
    //     }
        
    //     return isValid;
    // }

    notPassValidation() {

        const notValid = !super.validateModel();

        if (notValid) {

            super.onValidationFailed();
        }

        return notValid;
    }
}