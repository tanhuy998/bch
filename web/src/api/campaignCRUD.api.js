import AuthEndpoint from "../backend/autEndpoint";
import CRUDEndpoint from "../backend/crudEndpoint";
import HttpEndpoint from "../backend/endpoint";
import { campaign_model_t } from "../domain/models/campaign.model";
import NewCampaignResponsePresenter from "./presenter/response/newCampaignResponsePresenter";

export default class CampaignCRUDEndpoint extends CRUDEndpoint {

    constructor({scheme, host, port} = {}) {

        super({
            scheme, host, port,
            uri: '/campaigns'
        });
    }

    /**
     * 
     * @param {campaign_model_t} model 
     * @returns {NewCampaignResponsePresenter}
     */
    async create(campaignModel) {

        const res = await super.create(campaignModel);

        return new NewCampaignResponsePresenter(res);
    }

    async read() {


    }

    /**
     * 
     * @param {campaign_model_t} campaignModel 
     */
    async update(campaignModel) {

        super.update(campaignModel)
    }

    /**
     * 
     * @param {string} uuid 
     */
    async delete(uuid) {

        super.delete(uuid);
    }
}