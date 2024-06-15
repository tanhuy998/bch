import CampaignCRUD from "../../api/campaignCRUD.api";
import ErrorResponse from "../../backend/error/errorResponse";
import FormDelegator from "../../components/lib/formDelegator";
import { campaign_model_t } from "../models/campaign.model";

const NOT_VALIDATE_MODEL_FIELDS = new Set(['desciption']);

export default class NewCampaignUseCase extends FormDelegator {

    #endpoint = new CampaignCRUD();
    #dataModel = new campaign_model_t();

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

            await this.#endpoint.create(this.#dataModel);
        }
        catch (e) {

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
        let expireTime;

        for (const field in campaign_model_t) {

            if (NOT_VALIDATE_MODEL_FIELDS.has(field)) {

                continue;
            }

            const value = this.#dataModel[field];

            if ( !value ) {

                return false;
            }

            if (field === 'expire') {
                console.log(value)
                expireTime = new Date(value);
                continue;
            }
        }

        // if (
        //     expireTime.getDate() - issueTime.getDate() < 1
        //     || expireTime.getMonth() - issueTime.getMonth() <= 0
        //     || expireTime.getFullYear() - issueTime.getFullYear() <= 0
        // ) {

        //     return false;
        // }

        return true;
    }

    /**
     * @override
     */
    onValidationFailed() {

        
    }

    /**
     * 
     * @override
     */
    validateEveryInput() {

        return true;
    }
}