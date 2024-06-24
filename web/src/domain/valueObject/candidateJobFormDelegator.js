import Schema from "validate";
import { candidate_model_t } from "../models/candidate.model";
import CollectableFormDelegator from "./collectableFormDelegator";

export default class CandidateJobFormDelegator extends CollectableFormDelegator {

    #dataModel = new candidate_model_t();
    #validator = new Schema({

    })

    get validator() {

        return this.#validator;
    }

    get dataModel() {

        return this.#dataModel;
    }

    reset() {

        this.#dataModel = new candidate_model_t();
    }
}