import { useState } from "react";
import FormContext, { defaultFormContextValue } from "../contexts/form.context";
import Validator from "./lib/validator.";

/**
 * This component is wrapper for default html <form> element. This component
 * construct a context for it's chilldren <FormInput> to consumes especially
 * for instantly form input validating when input value changed.
 * 
 * dataModel is the object whose fields will be mapped with the form inputs value.
 * dataModel field are not assigned unless the validation on that field is failed.
 * 
 * onSubmit is listener for the submission event of the form. If there exists interceptSubmit,
 * onSubmit will be ignored.
 * 
 * if interceptSubmit is function, it will neither prevents 
 * the default behavior of the form element nor ignores onSubmit listener
 * 
 * 
 * @returns 
 */
export default function Form({ validateSubmit, onValidationFailed, interceptSubmit, onSubmit, children, validateField, dataModel }) {

    
    const allProps = arguments[0];
    const hasDataModel = typeof dataModel === 'object' || typeof dataModel === 'function';

    const hasInterceptor = typeof interceptSubmit === 'function';
    const hasSubmissionListener = typeof onSubmit === 'function';

    const isValidator = validateSubmit instanceof Validator;
    const hasSubmissionValidation = typeof validateSubmit === 'function' || isValidator;

    const validateFunc = isValidator ? 
                        function(dataObject) { return validateSubmit.validate(dataObject)}
                        : hasSubmissionValidation ? validateSubmit : undefined; 

    const emitValidationFailed = typeof onValidationFailed === 'function' ? onValidationFailed : () => {};
    
    const handleSubmit = !hasInterceptor ? hasSubmissionListener ? onSubmit : undefined
    : (function() {

        const event = arguments[0];

        if (typeof interceptSubmit !== 'function') {

            return;
        }

        event.preventDefault();

        const targetDataObj = hasDataModel ? dataModel : event.target;
        const validationOK = hasSubmissionValidation && validateFunc(targetDataObj) || !hasSubmissionValidation;

        if (
            !validationOK
        ) {

            emitValidationFailed(targetDataObj);
            return;
        }
        
        interceptSubmit(targetDataObj);
    });
    
    return (
        <FormContext.Provider value={{
            ...defaultFormContextValue, 
            dataModel: dataModel, 
            validate: validateField
        }}>
            <form {...allProps} onSubmit={handleSubmit} >
                {children}
            </form>
        </FormContext.Provider>
    )
}