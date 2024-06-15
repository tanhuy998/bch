import { useContext, useEffect, useState } from "react";
import FormContext from "../contexts/form.context";
import Validator from "./lib/validator.";

const FIRST = 0;
const INPUT_DEBOUNCE_DELAY = 1500;
const INIT_STATE = Symbol('init_state');

export const IGNORE_VALIDATION = Symbol('input_ignore_validator');

export default function FormInput({validate, onValidInput, onInvalidInput, invalidMessage, onAfterDebounce, name, textArea}) {

    const context = useContext(FormContext);
    const htmlElementAttributes = prepareRenderAttributes(arguments[FIRST]);
    const [dataModel, hasDataModel] = prepareDataModel(context);    

    const [debounceTimeout, setDebounceTimeout] = useState(null);
    const [inputCurrentValue, setInputCurrentValue] = useState(null);
    const [isValidInput, setIsValidInput] = useState(INIT_STATE);
    const [dataModelFieldValue, setDataModelFieldValue] = useState(dataModel?.[name]);
    
    const isIgnoreValidation = validate === IGNORE_VALIDATION;

    validate = prepareValidateFunction(validate, context);
    const hasValidation = typeof validate === 'function';
    

    onAfterDebounce = onAfterDebounce || context?.onAfterDebounce;
    
    onValidInput = (!isIgnoreValidation ? 
                    typeof onValidInput === 'function'?  onValidInput : context?.onValidInput 
                     : undefined);

    onInvalidInput = (isIgnoreValidation ? 
                    typeof onInvalidInput === 'function' ? onInvalidInput : context?.onInvalidInput
                    : undefined);

    const hasAfterDebounceListener = typeof onAfterDebounce === 'function';
    const hasValidationSuccessListener = typeof onValidInput === 'function';
    const hasValidationFailedListener = typeof onInvalidInput === 'function';

    const handleInputChange = (function() {

        const event = arguments[FIRST];

        setInputCurrentValue(event.target.value);
        
        const onChangeProp = htmlElementAttributes?.onChange;

        if (typeof onChangeProp === 'function') {

            onChangeProp(event);
        }
    })
    
    useEffect(() => {

        if (!hasDataModel) {

            return;
        }

        dataModel[name] = dataModelFieldValue;

    }, [dataModelFieldValue]);

    useEffect(() => {
       
        if (isValidInput === INIT_STATE) {

            return;
        }
        
        console.log('is input', [name], 'valid', isValidInput);
        (
            isValidInput && !setDataModelFieldValue(inputCurrentValue) ? 
            hasValidationSuccessListener ?  onValidInput(inputCurrentValue) : undefined
            : hasValidationFailedListener ? onValidInput(inputCurrentValue) : undefined
        )

    }, [isValidInput]);

    useEffect(() => {

        console.log('model field assign', dataModelFieldValue)
    }, [dataModelFieldValue])

    useEffect(() => {

        if (inputCurrentValue === null) {

            return;
        }
        console.log(debounceTimeout)
        if (
            typeof debounceTimeout === 'number'
        ) {

            clearTimeout(debounceTimeout);
        }

        setDebounceTimeout(setTimeout(() => {

            console.log('end input');

            if (hasAfterDebounceListener) {

                onAfterDebounce(inputCurrentValue);
            }

            if (!hasValidation || validate(inputCurrentValue)) {

                setIsValidInput(true);
                return;
            }

            setIsValidInput(false);

        }, INPUT_DEBOUNCE_DELAY));

    }, [inputCurrentValue])

    return (
        textArea === true ?
        <textarea {...{...htmlElementAttributes, onChange: handleInputChange, name}} >{inputCurrentValue}</textarea> 
        : <input {...{...htmlElementAttributes, onChange: handleInputChange, name}} value={inputCurrentValue}/>
    )
}

/**
 * 
 * @param {*} validate 
 * @param {*} context 
 * @returns {function name(params): Boolean {}} 
 */
function prepareValidateFunction(validate, context) {

    const isIgnoreValidation = (validate === IGNORE_VALIDATION);

    let isValidator = validate instanceof Validator;
    validate = (!isIgnoreValidation ? 
                typeof validate === 'function' || isValidator ? validate : context?.validate 
                : undefined);
    isValidator = validate instanceof Validator;

    validate = isValidator ? (val) => validate.validate(val) : validate;

    return validate;
}

function prepareDataModel(context) {

    const dataModel = context?.dataModel;
    const hasDataModel = typeof dataModel === 'object' || typeof dataModel === 'function';
    
    return [hasDataModel ? dataModel : undefined, hasDataModel];
}

function prepareRenderAttributes(props = {}) {

    return {
        ...props, 
        validate:undefined, 
        onValidInput:undefined, 
        onInvalidInput:undefined, 
        onAfterDebounce:undefined,
    };
}

function formatDateInput(val) {

    const d = new Date(val);

    return `${d.getDate()}-${d.getMonth()}-${d.getFullYear}`;
}