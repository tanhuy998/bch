import CampaignListEndpoint from "../../api/campaignList.api";

export default class CampaignListUseCase extends CampaignListEndpoint {

    #webUri = '/admin/campaigns';
    
    generateGetSingleCampaignURL(campaignUUID) {

        return this.#webUri + '/' + campaignUUID;
    }

    generateModifySingleCampaignURL(campaignUUID) {

        return this.#webUri + '/';
    }

    generateDeleteSingleCampaignURL(campaignUUID) {

        return this.#webUri;
    }
}