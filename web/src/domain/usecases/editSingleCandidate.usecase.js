import CRUDEndpoint from "../../backend/crudEndpoint";
import EditSingleCandidateFormDelegator from "../valueObject/editSingleCandidateFormDelegator";

export default class EditSingleCandidateUseCase extends CRUDEndpoint{

    #formDelegator = new EditSingleCandidateFormDelegator();

    get formDelegator() {

        return this.#formDelegator;
    }

    constructor() {

        super({})
    }
}