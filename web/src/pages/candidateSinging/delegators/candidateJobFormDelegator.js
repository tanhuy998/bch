import Schema from "validate";
import { candidate_model_t, candidate_signing_info_t } from "../../../domain/models/candidate.model";
import CollectableFormDelegator from "../../../domain/valueObject/collectableFormDelegator";

export default class CandidateJobFormDelegator extends CollectableFormDelegator {

    #dataModel = new candidate_signing_info_t();
    #validator = new Schema({
        
    })

    get validator() {

        return this.#validator;
    }

    get dataModel() {

        return this.#dataModel;
    }

    reset() {

        
    }
}