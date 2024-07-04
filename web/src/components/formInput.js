import { memo, useContext, useEffect, useState } from "react";
import FormContext, { child_input_proxy_t } from "../contexts/form.context";
import Validator from "./lib/validator.";
import "../assets/css/input.css";
import debug from "debug";

const formInputDebugger = debug('from-input');

const FIRST = 0;
const INPUT_DEBOUNCE_DELAY = 500;
const DATE_INPUT_DEBOUNCE_DELAY = 500;
const INIT_STATE = Symbol('init_state');

const NEED_DEBOUNCE_DELAY = new Set([
    'text', 'email', 'password', 'number', 'url', 
]);

export const IGNORE_VALIDATION = Symbol('input_ignore_validator');

const FormInput = memo(_FormInput);

export default FormInput;

export function _FormInput({defaulValue, validate, onValidInput, onInvalidInput, invalidMessage, onAfterDebounce, name, textArea, type, dataModelType}) {

    defaulValue = typeof defaulValue === 'string' || typeof defaulValue === 'number' ? defaulValue : undefined;
    const context = useContext(FormContext);
    const contextDelayingDebounces = context.delayingDebounces; 
    const htmlElementAttributes = prepareRenderAttributes(arguments[FIRST]);
    //const [dataModel, hasDataModel] = useDataModel();    
    const [inputCurrentValue, setInputCurrentValue] = useDataModelBinding(name, type, undefined, defaulValue);// useState(hasDataModel ? dataModel[name] : null);
    

    const inputProxy = useInputProxy(name, setInputCurrentValue);

    const isIgnoreValidation = validate === IGNORE_VALIDATION;

    validate = prepareValidateFunction(validate, context);

    onAfterDebounce = onAfterDebounce || context?.onAfterDebounce;
    const isValidInput = useInputFieldValidation(inputCurrentValue, validate);
    useInputProxyValidationListener(inputProxy, isValidInput);
    useValidationListener(isValidInput, inputCurrentValue, onValidInput, onInvalidInput, isIgnoreValidation);

    prepareElementClass(htmlElementAttributes, isValidInput);

    htmlElementAttributes.className += !isValidInput ? ' is-invalid' : ''; 

    const handleInputChange = (function () {

        const event = arguments[FIRST];

        setInputCurrentValue(event.target.value);

        const onChangeProp = htmlElementAttributes?.onChange;

        if (typeof onChangeProp === 'function') {

            onChangeProp(event);
        }
    });

    return (
        textArea === true ?
        <textarea {...{...htmlElementAttributes, onChange: handleInputChange, name}} >{inputCurrentValue || ''}</textarea> 
        : <input {...{...htmlElementAttributes, onChange: handleInputChange, name}} value={inputCurrentValue || ''}/>
    )
}

/**
 * 
 * @param {child_input_proxy_t} inputProxy 
 * @param {boolean} isValidInputState 
 */
function useInputProxyValidationListener(inputProxy, isValidInputState) {

    useEffect(() => {

        inputProxy.isValid = isValidInputState;

    }, [isValidInputState])
}

function useInputFieldValidation(inputCurrentValueState, validateFunc) {

    const [isValidInput, setIsValidInput] = useState(INIT_STATE);
    const hasValidation = typeof validateFunc === 'function';

    const emitEventAndValidate = () => {

        if (inputCurrentValueState === INIT_STATE) {

            return;
        }

        //formInputDebugger('end input', inputCurrentValue);

        // if (hasAfterDebounceListener) {

        //     onAfterDebounce(inputCurrentValue);
        // }
        console.log('timeout')
        if (!hasValidation || validateFunc(inputCurrentValueState)) {
            //formInputDebugger('not has validation or valid')
            //changValidInputState(true)
            setIsValidInput(true);
            return;
        }

        //changValidInputState(false);
        setIsValidInput(false);
    };

    useDebounce(inputCurrentValueState, emitEventAndValidate);

    return isValidInput;
}

function useValidationListener(isValidInputState, inputCurrentValueState, onValidInput, onInvalidInput, isIgnoreValidation) {

    const context = useContext(FormContext);

    onValidInput = (!isIgnoreValidation ?
        typeof onValidInput === 'function' ? onValidInput : context?.onValidInput
        : undefined);

    onInvalidInput = (!isIgnoreValidation ?
        typeof onInvalidInput === 'function' ? onInvalidInput : context?.onInvalidInput
        : undefined);

    //const hasAfterDebounceListener = typeof onAfterDebounce === 'function';
    const hasValidationSuccessListener = typeof onValidInput === 'function';
    const hasValidationFailedListener = typeof onInvalidInput === 'function';

    useEffect(() => {
        //formInputDebugger('isValidInput state change')
        if (isValidInputState === INIT_STATE) {
            //formInputDebugger('isValidInput init state')
            return;
        }
        //formInputDebugger('begin set data model field')
        //console.log('is input', [name], 'valid', isValidInput, onInvalidInput);
        (
            isValidInputState /* && !setDataModelFieldValue(inputCurrentValue) */?
            hasValidationSuccessListener && onValidInput(inputCurrentValueState)
            : hasValidationFailedListener && onInvalidInput(inputCurrentValueState)
        )

    }, [isValidInputState]);
}

function useDebounceListener(onAfterDebounce) {
    
    useEffect(() => {


    })
}

function useDebounce(inputCurrentValueState, cb, delayDuration = INPUT_DEBOUNCE_DELAY) {

    const [debounceTimeout, setDebounceTimeout] = useState(null);

    useEffect(() => {

        setDebounceTimeout(setTimeout(cb, delayDuration));
        console.log(debounceTimeout)

        return () => {
            console.log('clear timeout')
            clearTimeout(debounceTimeout);
        }

    }, [inputCurrentValueState]);
}

export function useDataModelBinding(inputName, inputType, transformDataFunc, initValue) {

    const [dataModel, hasDataModel] = useDataModel();
    const [inputCurrentValue, setInputCurrentValue] = useState(hasDataModel ? dataModel[inputName] : initValue);
    //const hasDataModel = typeof dataModel === 'object';
    const hasForceCastType = typeof transformDataFunc === 'function';

    useEffect(() => {

        if (!hasDataModel) {

            return;
        }
        
        dataModel[inputName] = hasForceCastType ? transformDataFunc(inputCurrentValue) : convertDataModelInputType(inputType, inputCurrentValue);
        console.log(inputCurrentValue, dataModel)
    }, [inputCurrentValue])

    return [inputCurrentValue, setInputCurrentValue];
}

function useInputProxy(name, setInputCurrentValue) {

    const {childrenInputs} = useContext(FormContext) || {};
    const [inputProxy] = useState(new child_input_proxy_t())

    useEffect(() => {

        if (!(childrenInputs instanceof Map)) {

            return;
        }

        if (childrenInputs?.has(name)) {

            return;
        }

        inputProxy.reset = () => {setInputCurrentValue(null)};

        childrenInputs?.set(name, inputProxy);

    }, []);

    return inputProxy;
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

function useDataModel() {

    const context = useContext(FormContext);
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

function convertDataModelInputType(ipnutType, dataModelFieldValue,isTextArea) {

    switch (ipnutType) {

        case 'date':
            return convertToDate(dataModelFieldValue);
        case 'number':
            return Number(ipnutType);
        default:
            return dataModelFieldValue;
    }
}

export function OldFormInput({ validate, onValidInput, onInvalidInput, invalidMessage, onAfterDebounce, name, textArea, type }) {

    const context = useContext(FormContext);
    const contextDelayingDebounces = context.delayingDebounces;
    const htmlElementAttributes = prepareRenderAttributes(arguments[FIRST]);
    const [dataModel, hasDataModel] = useDataModel(context);

    const [debounceTimeout, setDebounceTimeout] = useState(null);
    const [inputCurrentValue, setInputCurrentValue] = useState(hasDataModel ? dataModel[name] : null);
    const [isValidInput, setIsValidInput] = useState(INIT_STATE);
    const [dataModelFieldValue, setDataModelFieldValue] = useState(dataModel?.[name]);

    const inputProxy = useInputProxy(name, setInputCurrentValue);

    const isIgnoreValidation = validate === IGNORE_VALIDATION;

    validate = prepareValidateFunction(validate, context);
    const hasValidation = typeof validate === 'function';


    onAfterDebounce = onAfterDebounce || context?.onAfterDebounce;

    onValidInput = (!isIgnoreValidation ?
        typeof onValidInput === 'function' ? onValidInput : context?.onValidInput
        : undefined);

    onInvalidInput = (!isIgnoreValidation ?
        typeof onInvalidInput === 'function' ? onInvalidInput : context?.onInvalidInput
        : undefined);

    const hasAfterDebounceListener = typeof onAfterDebounce === 'function';
    const hasValidationSuccessListener = typeof onValidInput === 'function';
    const hasValidationFailedListener = typeof onInvalidInput === 'function';

    if (hasDataModel) {

        dataModel[name] = inputCurrentValue;
    }

    const handleInputChange = (function () {

        const event = arguments[FIRST];

        setInputCurrentValue(event.target.value);

        const onChangeProp = htmlElementAttributes?.onChange;

        if (typeof onChangeProp === 'function') {

            onChangeProp(event);
        }
    });

    const changValidInputState = (state) => {
        formInputDebugger('change isValidInput')
        setIsValidInput(INIT_STATE);
        setIsValidInput(state);
    }

    const emitEventAndValidate = () => {

        formInputDebugger('end input', inputCurrentValue);

        if (hasAfterDebounceListener) {

            onAfterDebounce(inputCurrentValue);
        }

        if (!hasValidation || validate(inputCurrentValue)) {
            formInputDebugger('not has validation or valid')
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
        formInputDebugger('isValidInput state change')
        if (isValidInput === INIT_STATE) {
            formInputDebugger('isValidInput init state')
            return;
        }
        formInputDebugger('begin set data model field')
        console.log('is input', [name], 'valid', isValidInput, onInvalidInput);
        (
            isValidInput /* && !setDataModelFieldValue(inputCurrentValue) */ ?
                hasValidationSuccessListener && onValidInput(inputCurrentValue)
                : hasValidationFailedListener && onInvalidInput(inputCurrentValue)
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
        formInputDebugger(debounceTimeout)
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
        console.log(newDebounceTimeout)
        contextDelayingDebounces.add(newDebounceTimeout);
        setDebounceTimeout(newDebounceTimeout);

    }, [inputCurrentValue])

    prepareElementClass(htmlElementAttributes, isValidInput);

    htmlElementAttributes.className += !isValidInput ? ' is-invalid' : '';

    return (
        textArea === true ?
            <textarea {...{ ...htmlElementAttributes, onChange: handleInputChange, name }} >{inputCurrentValue || ''}</textarea>
            : <input {...{ ...htmlElementAttributes, onChange: handleInputChange, name }} value={inputCurrentValue || ''} />
    )
}