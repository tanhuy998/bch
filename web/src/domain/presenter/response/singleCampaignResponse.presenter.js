import formatLocalDate from "../../../lib/formatLocalDate";

const UNKNOWN = 'Unknown';

export default class SingleCampaignRespnsePresenter {

    #campaignTitle;
    #issueTime;
    #expiredTime;

    get title() {

        return this.#campaignTitle;
    }

    get issueTime() {

        return this.#issueTime;
    }

    get expiredTime() {

        return this.#expiredTime;
    }
 
    constructor(resData) {
        
        this.#init(resData);
    }

    #init(resData) {

        this.#campaignTitle = resData?.title || UNKNOWN;
        this.#issueTime = formatLocalDate(new Date(resData?.issueTime)) || UNKNOWN;
        this.#expiredTime = formatLocalDate(new Date(resData?.expire)) || UNKNOWN;
    }
}