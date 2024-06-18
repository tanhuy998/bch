import ErrorResponse from "../../backend/error/errorResponse";
import AdvanceValidationFormDelegator from "./advancedValidationFormdelegator";

export class EndpointFormDelegatorAction {

    /**
     * @type {string}
     */
    static get CREATE() {

        return 'create';
    }

    /**
     * @type {string}
     */
    static get UPDATE() {

        return 'update';
    }
}

export default class EndpointFormDelegator extends AdvanceValidationFormDelegator {

    /**@type {import("../../backend/httpRequestBuilder").HttpEndpoint} */
    get endpoint() {


    }

    get action() {

        return EndpointFormDelegatorAction.CREATE;
    }

    /**
     * @param {any} 
     * @returns {string}
     */
    shouldNavigate(endpointResponse) {

        return '';
    }

    async interceptSubmission() {
        console.log('submit', this.dataModel)
        try {

            const action = this.action;
            const res = await this.endpoint[action](this.dataModel);

            const navigatePath = this.shouldNavigate(res);

            if (!navigatePath || typeof navigatePath !== 'string') {

                return
            }

            this.navigate(navigatePath);
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
}