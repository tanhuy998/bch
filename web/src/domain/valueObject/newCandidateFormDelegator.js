import Schema from "validate";
import FormDelegator from "../../components/lib/formDelegator"
import { candidate_model_t } from "../models/candidate.model";
import { validateIDNumber, validatePeopleName, aboveSixteenYearsOld } from "../../lib/validator";
import ErrorTraceFormDelegator from "./errorTraceFormDelegator";
import CandidateCRUDEndpoint from "../../api/candidateCRUD.api";

export default class NewCandidateFormDelegator extends ErrorTraceFormDelegator {

    #endpoint  = new CandidateCRUDEndpoint();
    #dataModel = new candidate_model_t();
    #validator = new Schema({
        name: {
            type: String,
            required: true,
            use: {validatePeopleName}
        },
        idNumber: {
            type: String,
            required: true,
            use: {validateIDNumber}
        },
        dateOfBirth: {
            type: Date,
            required: true,
            use : {aboveSixteenYearsOld}
        },
        address: {
            type: String,
            required: true,
        },
    });

    get dataModel() {

        return this.#dataModel;
    }

    get endpoint() {

        return this.#endpoint;
    }

    /**
     * 
     * @param {any} formData
     * 
     * @returns {boolean}
     */
    interceptSubmission() {


    }

    /**
     * 
     * @param {any} formData
     * 
     * @returns {boolean}
     */
    validateModel() {

        const errors = !this.#validator.validate(this.#dataModel);

        if (errors) {

            this.setValidationFailedFootPrint(errors);
            return false;
        }

        return true;
    }

    validateEveryInput() {

        return true;
    }
}