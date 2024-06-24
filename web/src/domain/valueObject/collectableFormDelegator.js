import AdvanceValidationFormDelegator from "./advancedValidationFormdelegator";

export default class CollectableFormDelegator extends AdvanceValidationFormDelegator {

    interceptSubmission() {

        return super.validateModel();
    }
}