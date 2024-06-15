import CampaignCRUD from "../../api/campaignCRUD.api";
import { campaign_model_t } from "../models/campaign.model";

export default class NewCampaignUseCase extends CampaignCRUD {

    #dataModel = new campaign_model_t();

    #campaignExpireDateValidateFunc = function(val) {

        return true;
    }

    get campaignExpireDateValidateFunc() {

        return this.#campaignExpireDateValidateFunc;
    }

    get dataModel() {
        
        return this.#dataModel;
    }

}