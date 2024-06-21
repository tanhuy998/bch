import { memo, useContext, useEffect, useState } from "react";
import FormContext from "../contexts/form.context";
import Validator from "./lib/validator.";
import "../assets/css/input.css";

const FIRST = 0;
const INPUT_DEBOUNCE_DELAY = 1500;
const DATE_INPUT_DEBOUNCE_DELAY = 500;
const INIT_STATE = Symbol('init_state');

const NEED_DEBOUNCE_DELAY = new Set([
    'text', 'email', 'password', 'number', 'url', 
]);

export const IGNORE_VALIDATION = Symbol('input_ignore_validator');

const FormInput = memo(_FormInput);

export default FormInput;

function _FormInput({validate, onValidInput, onInvalidInput, invalidMessage, onAfterDebounce, name, textArea, type}) {

    const context = useContext(FormContext);
    const contextDelayingDebounces = context.delayingDebounces; 
    const htmlElementAttributes = prepareRenderAttributes(arguments[FIRST]);
    const [dataModel, hasDataModel] = prepareDataModel(context);    

    const [debounceTimeout, setDebounceTimeout] = useState(null);
    const [inputCurrentValue, setInputCurrentValue] = useState(hasDataModel ? dataModel[name] : null);
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

    if (hasDataModel) {

        dataModel[name] = inputCurrentValue;
    }

    const handleInputChange = (function() {

        const event = arguments[FIRST];

        setInputCurrentValue(event.target.value);
        
        const onChangeProp = htmlElementAttributes?.onChange;

        if (typeof onChangeProp === 'function') {

            onChangeProp(event);
        }
    });
    
    const changValidInputState = (state) => {
        console.log('change isValidInput')
        setIsValidInput(INIT_STATE);
        setIsValidInput(state);
    }

    const emitEventAndValidate = () => {

        console.log('end input', inputCurrentValue);

        if (hasAfterDebounceListener) {

            onAfterDebounce(inputCurrentValue);
        }

        if (!hasValidation || validate(inputCurrentValue)) {
            console.log('not has validation or valid')
            changValidInputState(true)
            //setIsValidInput(true);
            return;
        }

        changValidInputState(false);
        //setIsValidInput(false);
    };
    
    // useEffect(() => {

    //     if (!hasDataModel) {

    //         return;
    //     }

    //     dataModel[name] = type === 'date' ? convertToDate(dataModelFieldValue) :  dataModelFieldValue;

    // }, [dataModelFieldValue]);

    useEffect(() => {
       console.log('isValidInput state change')
        if (isValidInput === INIT_STATE) {
            console.log('isValidInput init state')
            return;
        }
        console.log('begin set data model field')
        console.log('is input', [name], 'valid', isValidInput);
        (
            isValidInput /* && !setDataModelFieldValue(inputCurrentValue) */? 
            hasValidationSuccessListener ?  onValidInput(inputCurrentValue) : undefined
            : hasValidationFailedListener ? onInvalidInput(inputCurrentValue) : undefined
        )

    }, [isValidInput]);

    useEffect(() => {

        if (
            !hasDataModel 
            || typeof dataModelFieldValue !== 'string'
        ) {

            return;
        }

        dataModel[name] = type === 'date' ? convertToDate(dataModelFieldValue) : dataModelFieldValue;

        console.log('model field assign', dataModelFieldValue, dataModel)
    }, [dataModelFieldValue])

    useEffect(() => {

        if (inputCurrentValue === null) {

            return;
        }
        console.log(debounceTimeout)
        setDataModelFieldValue(inputCurrentValue);

        if (
            typeof debounceTimeout === 'number'
        ) {

            contextDelayingDebounces.delete(debounceTimeout);
            clearTimeout(debounceTimeout);
        }

        if (!textArea && !NEED_DEBOUNCE_DELAY.has(type)) {

            emitEventAndValidate();
            return;
        }

        const newDebounceTimeout = setTimeout(emitEventAndValidate, INPUT_DEBOUNCE_DELAY)

        contextDelayingDebounces.add(newDebounceTimeout);
        setDebounceTimeout(newDebounceTimeout);

    }, [inputCurrentValue])

    prepareElementClass(htmlElementAttributes, isValidInput);

    return (
        textArea === true ?
        <textarea {...{...htmlElementAttributes, onChange: handleInputChange, name}} >{inputCurrentValue || ''}</textarea> 
        : <input {...{...htmlElementAttributes, onChange: handleInputChange, name}} value={inputCurrentValue || ''}/>
    )
}

function prepareElementClass(htmlElementAttributes, isValidInputState) {

    /**@type {string} */
    let propClass = (htmlElementAttributes.className || '');

    if (!isValidInputState) {

        propClass.replace('invalid', '');
        propClass += ' invalid';
    }

    htmlElementAttributes.className = propClass;
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

function convertToDate(value) {

    return new Date(value);
}