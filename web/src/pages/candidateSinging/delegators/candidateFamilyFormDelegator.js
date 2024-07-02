import Schema from "validate";
import { candidate_signing_family_t } from "../../../domain/models/candidate.model";
import CollectableFormDelegator from "../../../domain/valueObject/collectableFormDelegator";

export default class CandidateParentFormDelegator extends CollectableFormDelegator {

    /**@type {string} */
    #who;
    #dataModel = new candidate_signing_family_t()
    #validator = new Schema({
        name: {
            type: String,
            required: true,
            message: "Tên là bắt buộc"
        },
        dateOfBirth: {
            type: Date,
            required: true,
            message: "Ngày sinh không hợp lệ",
        }
    })

    get validator() {

        return this.#validator;
    }

    get dataModel() {

        return this.#dataModel;
    }

    constructor(who) {

        super();
        this.#who = who
    }
}