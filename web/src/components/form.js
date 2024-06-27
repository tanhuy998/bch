import { useContext, useEffect, useReducer, useRef, useState } from "react";
import FormContext, { child_input_proxy_t, defaultFormContextValue } from "../contexts/form.context";
import Validator from "./lib/validator.";
import FormDelegator from "./lib/formDelegator";
import {useNavigate} from "react-router-dom";
import FormCollectorContext from "../contexts/formCollector.context";

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
    shouldNavigate,
}) {

    const hasDelegator = delegate instanceof FormDelegator;
    /**@type {FormDelegator} */
    const delegator = hasDelegator ? delegate : null;
    const navigate = useNavigate();
    const [childrenInputs] = useState(new Map());
    
    const formCollectorRef = useRefForFormCollector(hasDelegator, delegator);

    const resetFormContext = useReducer(() => {

        if (hasDelegator) {

            delegator.reset();
        }

        resetInputProxies(childrenInputs);

        return null;
    }, null)[1];

    if (
        typeof shouldNavigate !== 'string'
        && shouldNavigate !== undefined
        && shouldNavigate !== null
    ) {

        throw new Error('form shouldNavigate attribute must be a string');
    } 

    if (
        delegate !== null && delegate !== undefined 
        && !hasDelegator
    ) {

        throw new Error('invalid delegator passed to form whose type is not instance of FormDelegator');
    }

    dataModel = hasDelegator ? delegator.dataModel : dataModel;
    interceptSubmit = hasDelegator ? delegator.interceptSubmission.bind(delegator) : interceptSubmit;

    const allProps = arguments[0];
    const hasDataModel = (hasDelegator ? delegator.dataModel : undefined) || typeof dataModel === 'object' || typeof dataModel === 'function';

    const hasInterceptor = hasDelegator || typeof interceptSubmit === 'function';
    const hasSubmissionListener = hasDelegator || typeof onSubmit === 'function';

    const isValidator = validateSubmit instanceof Validator;
    const hasSubmissionValidation = hasDelegator || typeof validateSubmit === 'function' || isValidator;

    const validateFunc = (hasDelegator ? delegator.validateModel.bind(delegator) : undefined) 
                        || (    isValidator ? 
                                function(dataObject) { return validateSubmit.validate(dataObject)}
                                : hasSubmissionValidation ? validateSubmit : undefined
                            );

    const emitValidationFailed = (hasDelegator ? delegator.onValidationFailed.bind(delegator) : undefined) 
                                    || (typeof onValidationFailed === 'function' ? onValidationFailed : () => {});
    
    const delayingDebounceTimeouts = new Set();

    

    if (hasInterceptor) {

        delegator.setNavigator(navigate);
    }

    const handleSubmit = !hasInterceptor ? hasSubmissionListener ? onSubmit : undefined
    : (function() {

        //clearDebounces(delayingDebounceTimeouts);

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
            console.log('form validation failed')
            emitValidationFailed(targetDataObj);
            return;
        }
        
        interceptSubmit(targetDataObj);
        resetFormContext();
    });
    
    return (
        <FormContext.Provider value={{
            ...defaultFormContextValue, 
            delegate: undefined,
            dataModel: dataModel, 
            validate: validateField,
            delayingDebounces: delayingDebounceTimeouts,
            childrenInputs: childrenInputs,
        }}>
            <form {...allProps} ref={formCollectorRef} onSubmit={handleSubmit} >
                {children}
            </form>
        </FormContext.Provider>
    )
}

function useRefForFormCollector(hasDelegator, delegator) {

    const collectorContext = useContext(FormCollectorContext);
    const [registered, setRegistered] = useState(false);

    useEffect(() => {

        if (
            hasDelegator &&
            typeof collectorContext?.register === 'function'
            && !registered
        ) {

            setRegistered(true);
            collectorContext.register(delegator);
        }

    }, [])
}

/**
 * 
 * @param {Set<number>} list 
 */
function clearDebounces(list) {

    for (const timeout of list.values()) {

        clearTimeout(timeout);
    }
}

/**
 * 
 * @param {Map<string, child_input_proxy_t>} list 
 */
function resetInputProxies(list) {

    for (const proxy of list.values()) {

        if (typeof proxy.reset !== 'function') {

            continue;
        }

        proxy.reset();
    }
}