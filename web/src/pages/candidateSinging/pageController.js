import { createContext, useContext, useEffect, useRef, useState } from "react";
import TabEventContext from "../../contexts/tabEvent.contex";
import TabContext from "../../contexts/tab.context";
import { Tab } from "../../components/Tab";
import { pillTabStyle } from "../../contexts/tab.context";
import PageControllerContext from "./contexts/pageController.context";
import customTabContextvalue from "./constant";
import FormCollector from "../../components/formCollector";
import FormCollectorBus from "../../components/formCollectorBus";
import useCollectedForms from "../../hooks/formCollectorBus";
import debug from 'debug';
import useFormCollectorBusSession from "../../hooks/formCollectorBusSession";
import CollectableFormDelegator from "../../domain/valueObject/collectableFormDelegator";
import { useParams } from "react-router-dom";


const debugButton = debug('page-controller:button');
const debugPhase = debug('page-controller:phase')
const SUBMIT_PHASE = Infinity;

// function isFirstPagePhase(tabKey) {

//     return pagePhaseKeys.indexOf(tabKey) === 0;
// }

function NextPhaseButton({ resolveNextPhaseKey, pageUsecase }) {
    const { campaignUUID, candidateUUID } = useParams();
    const { currentTabKey, setCurrentTabKey, focusPoint, pageFormDelegators } = useContext(PageControllerContext);

    function dispatchNextPhase() {

        const nextPhaseKey = resolveNextPhaseKey(currentTabKey);
        debugButton('current phase', currentTabKey, 'next phase', nextPhaseKey)
        /**@type {Array<CollectableFormDelegator>} */
        const delegators = pageFormDelegators?.[currentTabKey];

        if (collectDelegatorValidationErrors(delegators).length > 0) {

            return;
        }

        if (nextPhaseKey === SUBMIT_PHASE) {

            pageUsecase.submit(campaignUUID, candidateUUID)
            return;
        }

        focusPoint?.current?.scrollIntoView({ behavior: 'smooth', block: 'start' });
        setCurrentTabKey(nextPhaseKey);
    }

    function handleClick() {

        dispatchNextPhase();
    }


    return (
        <div 
            style={{
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
            }}
        >
            <button type="button" onClick={handleClick} class="btn btn-warning mb-2" value={currentTabKey}>Tiếp Theo</button>
        </div>
    )
}

/**
 * 
 * @param {Array<CollectableFormDelegator>} delegators 
 * @returns {Array<any>}
 */
function collectDelegatorValidationErrors(delegators) {

    return (Array.isArray(delegators) ? delegators : [])
        .filter(d => d?.notPassValidation())
        .map(d => d.validationFailedFootPrint);
}

function PreviousPhaseButton() {

    return (
        <button></button>
    )
}

export default function PageController({ children, pagePhases, pageFormDelegators, pageUsecase}) {

    const pagePhaseKeys = Object.keys(pagePhases);
    const [currentTabKey, setCurrentTabKey] = useState(pagePhaseKeys[0]);
    // const [formCollectorHandShake, setFormCollectorHandShake] = useState(false);
    // const [formCollectorResponse, setFormCollectorResponse] = useState();
    // const [emitSignal, setEmitSignal] = useState();
    //const [enumerateTabKey, nextPhase] = useReducer(nextPagePhase, pagePhaseKeys[0]);

    const pageMainTab = useRef();

    function resolveNextPhaseKey(currentTabKey) {


        const pagePhasesCount = pagePhaseKeys.length;

        const currentPhaseIndex = pagePhaseKeys.indexOf(currentTabKey);

        if (currentPhaseIndex < 0) {

            throw new Error('unknown page phase');
        }

        if (currentPhaseIndex === pagePhasesCount - 1) {

            return SUBMIT_PHASE;
        }

        return pagePhaseKeys[currentPhaseIndex + 1];
    }

    debugPhase('phase', currentTabKey)
    return (
        <>
            {/* <FormCollectorDispatchContext.Provider
                value={{
                    emitSignal: null,
                    setEmitSignal,
                    formCollectorResponse,
                    setFormCollectorResponse,
                    formCollectorHandShake,
                    setFormCollectorHandShake,
                }}
            > */}
            {/* <FormCollectorBus> */}
                {/* <FormCollector sessionValue={currentTabKey}> */}
                    <PageControllerContext.Provider value={{ currentTabKey, setCurrentTabKey, focusPoint: pageMainTab, pageFormDelegators}}>

                        <div ref={pageMainTab}>
                            {/* <CustomTab initTabIndex={currentTab} tabs={tabContents} /> */}
                            <TabEventContext.Provider value={{ onTabSwitch: (oldTabKey, newTabKey) => { setCurrentTabKey(newTabKey) } }}>
                                <TabContext.Provider value={{ ...customTabContextvalue, currentTab: currentTabKey, }}>
                                    <Tab initTabKey={currentTabKey} tabs={pagePhases} />
                                    <br />

                                </TabContext.Provider>
                            </TabEventContext.Provider>
                            <NextPhaseButton pageUsecase={pageUsecase} resolveNextPhaseKey={resolveNextPhaseKey}/>
                        </div>

                    </PageControllerContext.Provider>
                {/* </FormCollector> */}
            {/* </FormCollectorBus> */}
            {/* </FormCollectorDispatchContext.Provider> */}
        </>
    )
}