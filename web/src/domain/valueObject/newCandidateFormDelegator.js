import Schema from "validate";
import FormDelegator from "../../components/lib/formDelegator"
import { candidate_model_t } from "../models/candidate.model";
import { validateIDNumber, validatePeopleName, ageAboveSixteenAndYoungerThanTwentySeven } from "../../lib/validator";
import ErrorTraceFormDelegator from "./errorTraceFormDelegator";
import CandidateCRUDEndpoint from "../../api/candidateCRUD.api";
import EndpointFormDelegator from "./endpointFormDelegator";

export default class NewCandidateFormDelegator extends EndpointFormDelegator {
    
    #refreshEmitter;

    /**@type {string} */
    #campaignUUID;
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
            use: { ageAboveSixteenAndYoungerThanTwentySeven },
            message: {
                use: 'candidate age must be in range from 17 to 26'
            }
        },
        address: {
            type: String,
            required: true,
        },
    });

    get validator() {

        return this.#validator;
    }

    get dataModel() {

        return this.#dataModel;
    }

    get endpoint() {

        return this.#endpoint;
    }

    get campaignUUID() {

        return this.#campaignUUID;
    }

    shouldNavigate() {

        return 0;
    }

    /**
     * 
     * @param {function} emitter 
     */
    setRefreshEmitter(emitter) {

        this.#refreshEmitter = emitter;
    }

    /**
     * 
     * @param {string} uuid 
     */
    setCampaignUUID(uuid) {

        if (!uuid || typeof uuid !== 'string') {

            throw new Error("campaing uuid must be string");
        }

        this.#campaignUUID = uuid;
    }

    reset() {
        console.log('reset')

        this.#dataModel = new candidate_model_t();

        if (typeof this.#refreshEmitter === 'function') {

            this.#refreshEmitter(true);
        }
        //super.navigate(0)
    }

    async interceptSubmission() {

        try {

            const res = await this.#endpoint.create(
                this.#dataModel, this.#campaignUUID
            )

            this.reset();
        }
        catch(e) {

            super._handleError(e);
        }

    }

    validateModel() {

        const dateOfBirth = this.#dataModel.dateOfBirth;

        if (!(dateOfBirth instanceof Date)) {

            this.#dataModel.dateOfBirth = new Date(dateOfBirth);
        }

        return super.validateModel();
    }

    validateEveryInput() {

        return true;
    }
}