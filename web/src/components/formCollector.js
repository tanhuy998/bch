import { useContext, useEffect, useReducer, useState } from "react";
import FormCollectorContext from "../contexts/formCollector.context";
import FormCollectorDispatchContext from "../contexts/formCollectorDispatch.context";

function addFormRefReducer(currentFormRefList, newForm) {

    return [...currentFormRefList || [], newForm];
}

/**
 * FormCollector collects wrapp it's children forms in order
 * to sumit forms that has delegator.
 * 
 * @param {*} param0 
 * @returns 
 */
export default function FormCollector({children}) {

    const [formRefList, addFormRef] = useReducer(addFormRefReducer, []);
    const { signal, setFormCollectorResponse, setFormCollectorHandShake } = useContext(FormCollectorDispatchContext) || {};

    

    function register(formRef) {

        addFormRef(formRef);
    }

    useEffect(() => {

        if (typeof setFormCollectorHandShake === 'function') {

            setFormCollectorHandShake(true);
        }

    }, [])

    useEffect(() => {

        console.log(formRefList)

    }, [formRefList])

    useEffect(() => {
        console.log('emit collector')
        if (formRefList.length === 0) {

            return;
        }

        for (const formRef of formRefList) {

            formRef?.current?.submit();
        }

        setFormCollectorResponse(true)

    }, [signal])

    return (
        <FormCollectorContext.Provider value={{register}}>
            {children}
        </FormCollectorContext.Provider>
    )
}