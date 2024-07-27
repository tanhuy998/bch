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

    async read(campaignUUID) {

        const res = await super.read(campaignUUID);
        /**@type {campaign_model_t} */
        const ret = Object.assign(new campaign_model_t(), res.data);

        ret.expire = new Date(res.data.expire);
        ret.issueTime = new Date(res.data.issueTime);

        return ret;
    }   

    /**
     * 
     * @param {campaign_model_t} campaignModel 
     */
    async update(campaignUUID, campaignModel) {

        return super.update(campaignUUID, campaignModel)
    }

    /**
     * 
     * @param {string} uuid 
     */
    async delete(uuid) {

        return super.delete(uuid);
    }
}