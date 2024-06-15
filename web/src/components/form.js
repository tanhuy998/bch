import { useState } from "react";
import FormContext, { defaultFormContextValue } from "../contexts/form.context";
import Validator from "./lib/validator.";
import FormDelegator from "./lib/formDelegator";

/**
 * This component is wrapper for default html <form> element. This component
 * construct a context for it's chilldren <FormInput> to consumes especially
 * for instantly form input validating when input value changed.
 * 
 */
export default function Form({ 
    /**
     * when delegate passed as prop, props corresponding to form handling will be skipped.
     * @type {FormDelegator}
     */
    delegate, 
    validateSubmit, 
    onValidationFailed, 
    /**
     * if interceptSubmit is function, it will neither prevents 
     * the default behavior of the form element nor ignores onSubmit listener
     */
    interceptSubmit, 
    /**
     * onSubmit is listener for the submission event of the form. If there exists interceptSubmit,
     * onSubmit will be ignored.
     * @type {function}
     */
    onSubmit, 
    validateField, 
    /**
     * dataModel is the object whose fields will be mapped with the form inputs value.
     * dataModel field are not assigned unless the validation on that field is failed.
     */
    dataModel,
    children, 
}) {

    const hasDelegator = delegate instanceof FormDelegator;

    if (
        delegate !== null && delegate !== undefined 
        && !hasDelegator
    ) {

        throw new Error('invalid delegator passed to form whose type is not instance of FormDelegator');
    }
    
    dataModel = hasDelegator ? delegate.dataModel : dataModel;
    interceptSubmit = hasDelegator ? delegate.interceptSubmission.bind(delegate) : interceptSubmit;

    const allProps = arguments[0];
    const hasDataModel = (hasDelegator ? delegate.dataModel : undefined) || typeof dataModel === 'object' || typeof dataModel === 'function';

    const hasInterceptor = hasDelegator || typeof interceptSubmit === 'function';
    const hasSubmissionListener = hasDelegator || typeof onSubmit === 'function';

    const isValidator = validateSubmit instanceof Validator;
    const hasSubmissionValidation = hasDelegator || typeof validateSubmit === 'function' || isValidator;

    const validateFunc = (hasDelegator ? delegate.validateModel.bind(delegate) : undefined) 
                        || (    isValidator ? 
                                function(dataObject) { return validateSubmit.validate(dataObject)}
                                : hasSubmissionValidation ? validateSubmit : undefined
                            );

    const emitValidationFailed = (hasDelegator ? delegate.onValidationFailed.bind(delegate) : undefined) 
                                    || (typeof onValidationFailed === 'function' ? onValidationFailed : () => {});
    
    const handleSubmit = !hasInterceptor ? hasSubmissionListener ? onSubmit : undefined
    : (function() {

        const event = arguments[0];

        if (
            typeof interceptSubmit !== 'function'
        ) {

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
            delegate: undefined,
            dataModel: dataModel, 
            validate: validateField
        }}>
            <form {...allProps} onSubmit={handleSubmit} >
                {children}
            </form>
        </FormContext.Provider>
    )
}