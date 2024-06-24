import Schema from "validate";
import { civil_identity_t } from "../models/candidate.model";
import CollectableFormDelegator from "./collectableFormDelegator";

export default class CanidateIdentityFormDelegator extends CollectableFormDelegator {

    #dataModel = new civil_identity_t();
    #validator = new Schema({
        
    });

    get validator() {

        return this.#validator;
    }

    get dataModel() {

        return this.#dataModel;
    }

    reset() {

        this.#dataModel = new civil_identity_t();
    }


}