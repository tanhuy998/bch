import { useContext, useEffect, useState } from "react";
import FormContext from "../contexts/form.context";

const INPUT_DEBOUNCE_DURATION = 1500;
const INIT_STATE = Symbol('init_state');

export const IGNORE_VALIDATION = Symbol('input_ignore_validator');

export default function FormInput({validate, onValidInput, onInvalidInput, onAfterDebounce, name}) {

    const context = useContext(FormContext);

    const htmlElementAttributes = {
        ...arguments[0], 
        validate:undefined, 
        onValidInput:undefined, 
        onInvalidInput:undefined, 
        onAfterDebounce:undefined,
    };

    const [debounceTimeout, setDebounceTimeout] = useState(null);
    const [inputCurrentValue, setInputCurrentValue] = useState(null);
    const [isValidInput, setIsValidInput] = useState(INIT_STATE);
    const [dataModelFieldValue, setDataModelFieldValue] = useState(null);

    const dataModel = context?.dataModel;
    const hasDataModel = typeof dataModel !== 'object' && typeof dataModel !== 'function';

    const isIgnoreValidation = (validate === IGNORE_VALIDATION);

    validate = (!isIgnoreValidation ? 
                typeof validate === 'function' ? validate : context?.validate 
                : undefined);

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

        const event = arguments[0];

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

        (
            isValidInput ? 
            hasValidationSuccessListener ? !setDataModelFieldValue(inputCurrentValue) && onValidInput(inputCurrentValue) : undefined
            : hasValidationFailedListener ? onValidInput(inputCurrentValue) : undefined
        )

    }, [isValidInput]);

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

        }, INPUT_DEBOUNCE_DURATION));

    }, [inputCurrentValue])

    return (
        <input onChange={handleInputChange} {...htmlElementAttributes}/>
    )
}