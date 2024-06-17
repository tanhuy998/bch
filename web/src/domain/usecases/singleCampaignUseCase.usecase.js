import CampaignCandidateListEndpoint from "../../api/campaignCandidateList.api";
import SingleCampaignRespnsePresenter from "../../api/presenter/response/singleCampaignResponse.presenter";
import SingleCampaignEndPoint from "../../api/singleCampaign.api";
import CandidateListTableRowManipulator from "../valueObject/candidateLisitTableRowManipulation";

export default class SingleCampaignUseCase extends SingleCampaignEndPoint {

    #CampaignCandidateListEndpoint = new CampaignCandidateListEndpoint();

    #candidateListTableRowManipulator;

    get candidateListTableRowManipulator() {

        return this.#candidateListTableRowManipulator;
    }

    constructor() {

        super();

        this.#candidateListTableRowManipulator = new CandidateListTableRowManipulator(this.url);
    }

    get campaignCandidateListEndpoint() {

        return this.#CampaignCandidateListEndpoint;
    }

    async fetch(uuid) {

        const res = await super.fetch(uuid);

        return new SingleCampaignRespnsePresenter(res?.data);
    }
}