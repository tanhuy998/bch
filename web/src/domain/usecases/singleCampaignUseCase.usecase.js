import CampaignCandidateListEndpoint from "../../api/campaignCandidateList.api";
import SingleCampaignEndPoint from "../../api/singleCampaign.api";
import SingleCampaignRespnsePresenter from "../presenter/response/singleCampaignResponse.presenter";

export default class SingleCampaignUseCase extends SingleCampaignEndPoint {

    #CampaignCandidateListEndpoint = new CampaignCandidateListEndpoint();

    get campaignCandidateListEndpoint() {

        return this.#CampaignCandidateListEndpoint;
    }

    async fetch(uuid) {

        const res = await super.fetch(uuid);

        return new SingleCampaignRespnsePresenter(res?.data);
    }
}