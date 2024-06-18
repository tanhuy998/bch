import FormDelegator from "../../components/lib/formDelegator";

export default class ErrorTraceFormDelegator extends FormDelegator {

    #errorQueue = [];

    get traceError() {

        return this.#errorQueue;
    }

    pushErrors(...err) {

        return this.#errorQueue.push(...err);
    }
}