import CampaignListEndpoint from "../../api/aggregates/campaignList.api";

export default class CampaignListUseCase extends CampaignListEndpoint {

    
    generateGetSingleCampaignURL(campaignUUID) {

        return this.url + '/' + campaignUUID;
    }

    generateModifySingleCampaignURL(campaignUUID) {

        return this.url + '/';
    }

    generateDeleteSingleCampaignURL(campaignUUID) {

        return this.ur;
    }
}