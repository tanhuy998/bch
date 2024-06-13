import CampaignListEndpoint from "../../api/campaignList.api";
import CampaignListTableRowManipulator from "../valueObject/campaignListTableRowManipulator";

export default class CampaignListUseCase extends CampaignListEndpoint {

    #tableRowManipulator;

    get tableRowManipulator() {

        return this.#tableRowManipulator;
    }

    #webUri = '/admin/campaigns';
    
    constructor() {

        super();

        this.#tableRowManipulator = new CampaignListTableRowManipulator(this.url);
    }

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