import CRUDEndpoint from "../../backend/crudEndpoint";
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

    /**@type {CRUDEndpoint} */
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

        if (!(this.endpoint instanceof CRUDEndpoint)) {


        }

        try {

            const action = this.action;
            const res = await this.endpoint[action](this.dataModel);

            this.navigateAfterInterceptionSuccess(res);
        }
        catch (e) {
            console.log('endpoint throw error')
            this.#handleError(e);
        }
    }

    navigateAfterInterceptionSuccess(res) {

        const navigatePath = this.shouldNavigate(res);

        if (!navigatePath || typeof navigatePath !== 'string') {

            return
        }

        this.navigate(navigatePath);
    }

    _handleError(err) {

        this.#handleError(err);
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