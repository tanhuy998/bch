import { createContext } from "react";



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