import Schema from "validate";
import CRUDEndpoint from "../../backend/crudEndpoint";
import { login_user } from "../../domain/models/auth.model";
import AdvanceValidationFormDelegator from "../../domain/valueObject/advancedValidationFormdelegator";
import EndpointFormDelegator from "../../domain/valueObject/endpointFormDelegator";
import LoginEndpoint, { FORM_DELEGATOR_ENDPOINT_ACTION } from "./endpoint";

export default class LoginUseCase extends EndpointFormDelegator {

    #validator = new Schema({
        username: {
            type: String,
            required: true,
        },
        password: {
            type: String,
            required: true,
        },
    });
    #enpoint = new LoginEndpoint()

    #dataModel = new login_user();

    /**
     * @type {LoginEndpoint}
     */
    get endpoint() {

        return this.#enpoint;
    }

    get dataModel() {

        return this.#dataModel;
    }

    get validator() {

        return this.#validator;
    }

    get endpointAction() {

        return FORM_DELEGATOR_ENDPOINT_ACTION
    }

    shouldNavigate() {

        return "/auth/switch";
    }

    validateEveryInput() {

        return true;
    }
}