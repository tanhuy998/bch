import Schema from "validate";
import CollectableFormDelegator from "./collectableFormDelegator";
import { candidate_signing_education_t } from "../models/candidate.model";

export default class CandidateEducationFormDelegator extends CollectableFormDelegator {

    #dataModel = new candidate_signing_education_t();
    #validator = new Schema({

    });

    get validator() {

        return this.#validator;
    }

    get dataModel() {

        return this.#dataModel;
    }

    reset() {

        this.#dataModel = new candidate_signing_education_t();
    }
}