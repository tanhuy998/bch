import { createContext } from "react";

export function child_input_proxy_t() {

    /**
     * @type {function}
     */
    this.reset = null;

    /**
     * @type {any}
     */
    this.inputCurrentValue = undefined;

    /**
     * @type {boolean}
     */
    this.isValid;
}

export const defaultFormContextValue = {
    validate: null,
    dataModel: undefined,
    onValidInput: null, 
    onInvalidInput: null, 
    onAfterDebounce: null,
    delayingDebounces: null,
}

const FormContext = createContext(defaultFormContextValue);

export default FormContext;