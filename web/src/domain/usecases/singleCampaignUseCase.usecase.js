import CampaignCandidateListEndpoint from "../../api/campaignCandidateList.api";
import CampaignProgressEndpoint from "../../api/campaignProgress.api";
import SingleCampaignRespnsePresenter from "../../api/presenter/response/singleCampaignResponse.presenter";
import SingleCampaignEndPoint from "../../api/singleCampaign.api";
import CandidateListTableRowManipulator from "../valueObject/candidateLisitTableRowManipulation";
import NewCandidateFormDelegator from "../valueObject/newCandidateFormDelegator";

export default class SingleCampaignUseCase extends SingleCampaignEndPoint {

    #CampaignCandidateListEndpoint = new CampaignCandidateListEndpoint();
    #newCandidateFormDelegator = new NewCandidateFormDelegator();
    #candidateListTableRowManipulator = new CandidateListTableRowManipulator(this.#newCandidateFormDelegator.endpoint.url);
    #campaignProgressEndpoint = new CampaignProgressEndpoint();

    get campaignProgressEndpoint() {

        return this.#campaignProgressEndpoint;
    }

    get newCandidateFormDelegator() {

        return this.#newCandidateFormDelegator;
    }

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