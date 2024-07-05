import Schema from "validate";
import { civil_identity_t } from "../../../domain/models/candidate.model";
import CollectableFormDelegator from "../../../domain/valueObject/collectableFormDelegator";
import {validateIDNumber, validatePeopleName, ageAboveSixteenAndYoungerThanTwentySeven, validateFormalName} from "../../../lib/validator";

export default class CanidateIdentityFormDelegator extends CollectableFormDelegator {

    #dataModel = new civil_identity_t();
    #validator = new Schema({
        idNumber: {
            type: String,
            required: true,
            use: {validateIDNumber},
            message: "Số căn cước công dân không hợp lệ",
        },
        name: {
            type: String,
            required: true,
            use: {validatePeopleName},
            message: "Tên không hợp lệ",
        },
        dateOfBirth: {
            type: Date,
            required: true,
            use: {ageAboveSixteenAndYoungerThanTwentySeven},
            message: "Ngày sinh không hợp lệ"
        },
        birthPlace: {
            type: String,
            required: true,
            message: "Nơi đăng ký khai sinh không hợp lệ"
        },
        ethnicity: {
            type: String,
            required: true,
            use: {validateFormalName},
            message: "Dân tộc không hợp lệ",
        },
        religion: {
            type: String,
            required: false,
            use: {validateFormalName},
            message: "Tôn giáo không hợp lệ",
        },
        permanentResident: {
            type: String,
            required: true,
            message: "Địa chỉ thường trú là bắt buộc",
        },
        male: {
            type: Boolean,
            required: true,
        },
        placeOfOrigin: {
            type: String,
            required: true,
            use: {validateFormalName},
            message: "Quê quán không hợp lệ",
        },
        nationality: {
            type: String,
            required: true,
            use: {validateFormalName},
            message: "Quốc tịch không hợp lệ và không được để trống"
        },
    });

    get validator() {

        return this.#validator;
    }

    get dataModel() {

        return this.#dataModel;
    }
    
    notPassValidation() {  

        console.log(this.#dataModel);

        return super.notPassValidation();
    }

    reset() {

    
    }
}