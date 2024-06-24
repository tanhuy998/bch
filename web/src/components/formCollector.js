import { useContext, useEffect, useReducer, useState } from "react";
import FormCollectorContext from "../contexts/formCollector.context";
import FormCollectorDispatchContext from "../contexts/formCollectorDispatch.context";
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

    const [currentSessionValue, setCurrentSessionValue] = useState(null);
    const [formDelegatorList, addFormDelegatorToList] = useReducer(addFormRefReducer, []);
    const [newCollectedDelegator, setNewCollectedDelegator] = useState(null);
    const [endCollectSignal, setEndCollectSignal] = useState(false);
    const { signal, setFormCollectorResponse, setFormCollectorHandShake } = useContext(FormCollectorDispatchContext) || {};


    function register(formRef) {

        addFormDelegatorToList(formRef);
    }

    // useEffect(() => {

    //     if (typeof setFormCollectorHandShake === 'function') {

    //         setFormCollectorHandShake(true);
    //     }

    // }, [])

    useEffect(() => {
        console.log('form collector session', sessionValue)    
        if (sessionValue !== currentSessionValue) {
            console.log('new session', sessionValue)
            setCurrentSessionValue(sessionValue);
            setEndCollectSignal(false);
            addFormDelegatorToList(NEW_SESSION);
        }
    })

    
    useEffect(() => {

        
        if (sessionValue !== currentSessionValue) {

            setCurrentSessionValue(sessionValue);
            setEndCollectSignal(false);
            addFormDelegatorToList(NEW_SESSION);
        }

    }, []);

    useEffect(() => {

        if (endCollectSignal) {
            
            console.log(`session ${sessionValue} collected forms`, formDelegatorList)
        }

    }, [endCollectSignal]);

    useEffect(() => {
        
        if (!(newCollectedDelegator instanceof CollectableFormDelegator)) {

            return;
        }

        if (endCollectSignal) {

            return;
        }

        console.log('add delegator')
        console.log(formDelegatorList)
        addFormDelegatorToList(newCollectedDelegator);

    }, [newCollectedDelegator])

    // useEffect(() => {

    //     if (!endCollectSignal) {

    //         return;
    //     }
    //     console.log('====================')
    //     if (formDelegatorList.length === 0) {

    //         setFormCollectorResponse(true);
    //         return;
    //     }
    //     console.log('collector receive emit signal')
    //     let res = true;

    //     for (const delegator of formDelegatorList) {
    //         console.log('do delegator interception')
    //         if (!delegator.interceptSubmission()) {

    //             res = false;
    //         }
    //     }

    //     setFormCollectorResponse(res);

    // }, [signal])

    return (
        <FormCollectorContext.Provider value={{register}}>
            {children}
            <EndCollector emit={setEndCollectSignal} sessionValue={sessionValue}/>
        </FormCollectorContext.Provider>
    )
}

function EndCollector({emit, sessionValue}) {

    useEffect(() => {

        emit(true);

    }, [sessionValue])

    return (
        <>
        </>
    )
}