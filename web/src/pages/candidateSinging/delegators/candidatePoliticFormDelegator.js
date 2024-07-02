import Schema from "validate";
import { candidate_signing_politic_t } from "../../../domain/models/candidate.model";
import CollectableFormDelegator from "../../../domain/valueObject/collectableFormDelegator";

export default class CandidatePoliticFormDelegator extends CollectableFormDelegator {

    #dataModel = new candidate_signing_politic_t();
    #validator = new Schema({

    })

    get dataModel() {

        return this.#dataModel;
    }

    get validator() {

        return this.#validator;
    }
}