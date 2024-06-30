import Schema from "validate";
import CollectableFormDelegator from "../../../domain/valueObject/collectableFormDelegator";
import { candidate_signing_education_t } from "../../../domain/models/candidate.model";

export default class CandidateEducationFormDelegator extends CollectableFormDelegator {

    #dataModel = new candidate_signing_education_t();
    #validator = new Schema({
        highestGrade: {
            type: Number,
            message: "Trình độ văn hóa không hợp lệ"
        }
    });

    get validator() {

        return this.#validator;
    }

    get dataModel() {

        return this.#dataModel;
    }

    reset() {


    }
}