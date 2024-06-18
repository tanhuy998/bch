import CRUDEndpoint from "../backend/crudEndpoint";
import { candidate_model_t } from "../domain/models/candidate.model";

export default class CandidateCRUDEndpoint extends CRUDEndpoint {

    constructor({scheme, host, port} = {}) {

        super({
            scheme, host, port,
            uri: '/candidates',
        })
    }   

    /**
     * 
     * @param {candidate_model_t} model 
     * @returns {NewCampaignResponsePresenter}
     */
    async create(candidateModel) {

        const res = await super.create(candidateModel);

        //return new NewCampaignResponsePresenter(res);
    }

    async read() {


    }

    /**
     * 
     * @param {candidate_model_t} candidateModel 
     */
    async update(candidateModel) {

        super.update(candidateModel)
    }

    /**
     * 
     * @param {string} uuid 
     */
    async delete(uuid) {

        super.delete(uuid);
    }
}