import ErrorResponse from "../../backend/error/errorResponse";
import FormDelegator from "../../components/lib/formDelegator";
import { campaign_model_t } from "../models/campaign.model";
import NewCampaignResponsePresenter from "../../api/presenter/response/newCampaignResponsePresenter";
import CampaignCRUDEndpoint from "../../api/campaignCRUD.api";
import Schema from "validate";

export default class NewCampaignUseCase extends FormDelegator {

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
            message: {
                type: 'invalid date format of exipre.',
                required: 'expire is required.'
            }
        }
    })

    #campaignExpireDateValidateFunc = function(val) {

        return true;
    }

    get campaignExpireDateValidateFunc() {

        return this.#campaignExpireDateValidateFunc;
    }

    /**
     * @override
     */
    get dataModel() {
        
        return this.#dataModel;
    }

    /**
     * @override
     */
    async interceptSubmission() {
        console.log('submit', this.#dataModel)
        try {

            const presenter = await this.#endpoint.create(this.#dataModel);

            this.navigate(`/admin/campaign/${presenter.createdUUID}`);
        }
        catch (e) {
            console.log('endpoint throw error')
            this.#handleError(e);
        }
    }

    #handleError(err) {

        if (err instanceof ErrorResponse) {

            return this.#handleEndpointErrorResponse(err);
        }

        alert(err.message);
    }

    /**
     * 
     * @param {ErrorResponse} e 
     */
    async #handleEndpointErrorResponse(e) {

        const resObj = e.responseObject;
        const resVal = await resObj.json();

        alert(resVal.message)
    }

    /**
     * @override
     */
    validateModel() {
        
        const issueTime = new Date();
        let expireTime = this.#dataModel.expire;
        expireTime = !(expireTime instanceof Date) ? new Date(expireTime) : expireTime;

        this.#dataModel.issueTime = issueTime;
        
        console.log('validate form', this.#dataModel, JSON.stringify(this.#dataModel))

        const errors = this.#validator.validate(this.#dataModel);

        if (errors.length > 0) {

            this.setValidationFailedFootPrint(errors);
            return false;
        }
        
        if (
           !(expireTime.getDate() - issueTime.getDate() < 1
            || expireTime.getMonth() - issueTime.getMonth() <= 0
            || expireTime.getFullYear() - issueTime.getFullYear() <= 0)
        ) {
            this.setValidationFailedFootPrint(new Error('expire date must be at least the day after present'));
            return false;
        }

        return true;
    }

    /**
     * @override
     */
    onValidationFailed() {

        const errors = this.validationFailedFootPrint;

        if (!errors) {

            alert('form validation failed')
        }

        if (Array.isArray(errors)) {

            const msg = errors.map(
                err => err.message
            ).join("\n");

            alert(msg);
        }

        alert(errors?.message || errors);
    }

    /**
     * 
     * @override
     */
    validateEveryInput() {

        return true;
    }
}