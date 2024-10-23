import CRUDEndpoint from "../../backend/crudEndpoint";
import HttpEndpoint from "../../backend/endpoint";
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

    get endpointAction() {

        return this.action
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

        this.#checkEnpoint();

        try {

            const action = this.endpointAction;
            const res = await this.endpoint[action](this.dataModel);

            this.navigateAfterInterceptionSuccess(res);
        }
        catch (e) {
            console.log('endpoint throw error')
            this.#handleError(e);
        }
    }

    #checkEnpoint() {

        if (!(this.endpoint instanceof HttpEndpoint)) {

            throw new Error('invalid endpoint to sumit a form which is not an instance of CRUDEndpoint');
        }

        if (typeof this.endpoint[this.endpointAction] != 'function') {

            throw new Error('the endpoint action of FormDelegator is not a function');
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