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
     * @param {string} campaignUUID
     * @returns {NewCampaignResponsePresenter}
     */
    async create(candidateModel, campaignUUID) {

        const res = await super.fetch(
            {
                method: 'POST',
                body: JSON.stringify({
                    data: candidateModel
                })
            },
            undefined,
            `/campaign/${campaignUUID}`
        )

        return res;
        //return new NewCampaignResponsePresenter(res);
    }

    async read(uuid) {

        const res = await super.read(uuid);
        
        return res.data;
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