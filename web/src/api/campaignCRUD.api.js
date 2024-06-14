import AuthEndpoint from "../backend/autEndpoint";
import HttpEndpoint from "../backend/endpoint";
import { campaign_model_t } from "../domain/models/campaign.model";

export default class CampaignCRUD extends AuthEndpoint {

    constructor({scheme, host, port} = {}) {

        super({
            scheme, host, port,
            uri: '/campaigns'
        });
    }

    /**
     * 
     * @param {campaign_model_t} model 
     */
    async create(campaignModel) {

        return super.fetch(
            {
                method: 'POST',
                body: JSON.stringify(campaignModel),
            }
        )
    }

    async read() {


    }

    async update() {


    }

    async delete() {


    }
}