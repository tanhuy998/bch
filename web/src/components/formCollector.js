import { useContext, useEffect, useReducer, useState } from "react";
import FormCollectorContext from "../contexts/formCollector.context";
import FormCollectorDispatchContext from "../contexts/formCollectorDispatch.context";
import CollectableFormDelegator from "../domain/valueObject/collectableFormDelegator";

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

    const [formDelegatorList, addFormDelegatorToList] = useReducer(addFormRefReducer, []);
    const [newCollectedDelegator, setNewCollectedDelegator] = useState(null);
    const [endCollectSignal, setEndCollectSignal] = useState(false);
    const { signal, setFormCollectorResponse, setFormCollectorHandShake } = useContext(FormCollectorDispatchContext) || {};

    

    function register(formRef) {

        addFormDelegatorToList(formRef);
    }

    useEffect(() => {

        if (typeof setFormCollectorHandShake === 'function') {

            setFormCollectorHandShake(true);
        }

    }, [])

    useEffect(() => {
        
        if (!(newCollectedDelegator instanceof CollectableFormDelegator)) {

            return;
        }
        console.log(formDelegatorList)
        addFormDelegatorToList(newCollectedDelegator);

    }, [newCollectedDelegator])

    useEffect(() => {
        console.log('emit collector')

        if (!endCollectSignal) {

            return;
        }

        if (formDelegatorList.length === 0) {

            setFormCollectorResponse(true);
            return;
        }

        let res = true;

        for (const delegator of formDelegatorList) {

            if (!delegator.interceptSubmission()) {

                res = false;
            }
        }

        setFormCollectorResponse(res);

    }, [signal])

    return (
        <FormCollectorContext.Provider value={{register}}>
            {children}
            <EndCollector emit={setEndCollectSignal}/>
        </FormCollectorContext.Provider>
    )
}

function EndCollector({emit}) {

    return (
        <>
        {emit(true)}
        </>
    )
}