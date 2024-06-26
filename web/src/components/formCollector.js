import { useContext, useEffect, useReducer, useState } from "react";
import FormCollectorContext from "../contexts/formCollector.context";
import FormCollectorBusContext from "../contexts/formCollectorBus.context";
import CollectableFormDelegator from "../domain/valueObject/collectableFormDelegator";

const NEW_SESSION = Symbol('new_session');

function addFormRefReducer(currentFormRefList, newForm) {

    if (newForm === NEW_SESSION) {

        return [];
    }

    return [...currentFormRefList || [], newForm];
}

/**
 * FormCollector collects wrapp it's children forms in order
 * to sumit forms that has delegator.
 * 
 * @param {*} param0 
 * @returns 
 */
export default function FormCollector({sessionValue, children}) {

    const [formDelegatorList, addFormDelegatorToList] = useReducer(addFormRefReducer, []);
    const [endCollected, setEndCollected] = useState(false);
    const { setCollectedDelegator,  setCollectedEndSignal, setBusSession} = useContext(FormCollectorBusContext) || {};

    const insideCollectorBus = typeof setCollectedDelegator === 'function' && typeof setCollectedEndSignal === 'function'

    function register(formRef) {

        addFormDelegatorToList(formRef);
    }


    useEffect(() => {

        if (!endCollected) {

            return;
        }

        setCollectedDelegator(formDelegatorList);

    }, [endCollected])

    return (
        <FormCollectorContext.Provider value={{register}}>
            {children}
            <EndCollector emit={setEndCollected} sessionValue={sessionValue}/>
        </FormCollectorContext.Provider>
    )
}

function EndCollector({ emit, sessionValue, propSessionValue}) {

    useEffect(() => {

        emit(true);

    })

    return (
        <>
        </>
    )
}