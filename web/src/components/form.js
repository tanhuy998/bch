import { useState } from "react";
import FormContext, { defaultFormContextValue } from "../contexts/form.context";

/**
 * This component is wrapper for default html <form> element. This component
 * construct a context for it's chilldren <FormInput> to consumes especially
 * for instantly form input validating when input value changed.
 * 
 * if handleFormData is funciton, it will be the interceptor that prevent
 * the default behavior of the form element.
 * 
 * @returns 
 */
export default function Form({ handleFormData, children, validate, dataModel }) {

    
    const allProps = arguments[0];
    const handleSubmit = (function handleSubmit() {

        const event = arguments[0];

        if (typeof handleFormData === 'function') {

            event.preventDefault();

            handleFormData(event.target);

            return;
        }
    });

    return (
        <FormContext.Provider value={{
            ...defaultFormContextValue, 
            'dataModel': dataModel, 
            validate
        }}>
            <form method="post" onSubmit={handleSubmit} {...allProps}>
                {children}
            </form>
        </FormContext.Provider>
    )
} 

