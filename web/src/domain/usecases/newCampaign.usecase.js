import ErrorResponse from "../../backend/error/errorResponse";
import FormDelegator from "../../components/lib/formDelegator";
import { campaign_model_t } from "../models/campaign.model";
import NewCampaignResponsePresenter from "../../api/presenter/response/newCampaignResponsePresenter";
import CampaignCRUDEndpoint from "../../api/campaignCRUD.api";
import Schema from "validate";
import ErrorTraceFormDelegator from "../valueObject/errorTraceFormDelegator";
import {dayAfterNow} from "../../lib/validator";
import EndpointFormDelegator from "../valueObject/endpointFormDelegator";

export default class NewCampaignUseCase extends EndpointFormDelegator {

    #endpoint = new CampaignCRUDEndpoint();
    #dataModel = new campaign_model_t();
    #validator = new Schema({
        title: {
            type: String,
            required: true,
            length: {min: 5},
        },
        expire: {
            type: Date,
            required: true,
            use: {dayAfterNow},
            message: {
                type: 'invalid date format of exipre.',
                required: 'expire is required.'
            }
        }
    })

    #campaignExpireDateValidateFunc = function (val) {

        return true;
    }

    get validator() {

        return this.#validator;
    }

    get endpoint() {

        return this.#endpoint;
    }

    /**
 * @override
 */
    get dataModel() {

        return this.#dataModel;
    }


    get campaignExpireDateValidateFunc() {

        return this.#campaignExpireDateValidateFunc;
    }

    /**
     * 
     * @param {NewCampaignResponsePresenter} presenter 
     * @returns {string}
     */
    shouldNavigate(presenter) {

        return `/admin/campaign/${presenter.createdUUID}`;
    }

    // /**
    //  * @override
    //  */
    // validateModel() {
        
    //     const issueTime = new Date();
    //     let expireTime = this.#dataModel.expire;
    //     expireTime = !(expireTime instanceof Date) ? new Date(expireTime) : expireTime;

    //     this.#dataModel.issueTime = issueTime;
        
    //     console.log('validate form', this.#dataModel, JSON.stringify(this.#dataModel))

    //     const errors = this.#validator.validate(this.#dataModel);

    //     if (errors.length > 0) {

    //         this.setValidationFailedFootPrint(errors);
    //         return false;
    //     }
        
    //     if (
    //        !(expireTime.getDate() - issueTime.getDate() < 1
    //         || expireTime.getMonth() - issueTime.getMonth() <= 0
    //         || expireTime.getFullYear() - issueTime.getFullYear() <= 0)
    //     ) {
    //         FormDelegator
    //         this.setValidationFailedFootPrint(new Error('expire date must be at least the day after present'));
    //         return false;
    //     }

    //     return true;
    // }

    /**
     * 
     * @override
     */
    validateEveryInput() {

        return true;
    }
}