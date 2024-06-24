import { createContext, useContext, useEffect, useRef, useState } from "react";
import FormCollectorDispatchContext from "../../contexts/formCollectorDispatch.context";
import TabEventContext from "../../contexts/tabEvent.contex";
import TabContext from "../../contexts/tab.context";
import { Tab } from "../../components/Tab";
import { pillTabStyle } from "../../contexts/tab.context";
import PageControllerContext from "./contexts/pageController.context";
import customTabContextvalue from "./constant";
import FormCollector from "../../components/formCollector";
import FormCollectorBus from "../../components/formCollectorBus";

const SUBMIT_PHASE = Infinity;

// function isFirstPagePhase(tabKey) {

//     return pagePhaseKeys.indexOf(tabKey) === 0;
// }

function NextPhaseButton({ resolveNextPhaseKey }) {

    const { currentTabKey, setCurrentTabKey, focusPoint } = useContext(PageControllerContext);
    const { formCollectorResponse, setEmitSignal, formCollectorHandShake, setFormCollectorHandShake } = useContext(FormCollectorDispatchContext) || {};

    const hasSignalEmit = typeof setEmitSignal === 'function';
    const hasHandShakeSetter = typeof setFormCollectorHandShake === 'function';

    const isHandShaked = formCollectorHandShake === true;

    function dispatchNextPhase() {

        const nextPhaseKey = resolveNextPhaseKey(currentTabKey);
        console.log('current phase', currentTabKey, 'next phase', nextPhaseKey)
        if (nextPhaseKey === SUBMIT_PHASE) {

            return;
        }

        focusPoint?.current?.scrollIntoView({ behavior: 'smooth', block: 'start' });
        setCurrentTabKey(nextPhaseKey);
    }

    function handleClick() {

        if (
            hasSignalEmit
            //&& isHandShaked
        ) {

            console.log('button emit');
            setEmitSignal(currentTabKey);
            return;
        }

        dispatchNextPhase();
    }

    useEffect(() => {

        if (formCollectorResponse === true) {

            dispatchNextPhase();
            setFormCollectorHandShake(false);
        }

    }, [formCollectorResponse]);

    return (
        <button type="button" onClick={handleClick} class="btn btn-outline-primary mb-2" value={currentTabKey}>Tiếp Theo</button>
    )
}

function PreviousPhaseButton() {

    return (
        <button></button>
    )
}

export default function PageController({ children, pagePhases }) {

    const pagePhaseKeys = Object.keys(pagePhases);
    const [currentTabKey, setCurrentTabKey] = useState(pagePhaseKeys[0]);
    const [formCollectorHandShake, setFormCollectorHandShake] = useState(false);
    const [formCollectorResponse, setFormCollectorResponse] = useState();
    const [emitSignal, setEmitSignal] = useState();
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

    console.log('phase', currentTabKey)
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
            <FormCollectorBus>
                <FormCollector sessionValue={currentTabKey}>
                    <PageControllerContext.Provider value={{ currentTabKey, setCurrentTabKey, focusPoint: pageMainTab, }}>

                        <div ref={pageMainTab}>
                            {/* <CustomTab initTabIndex={currentTab} tabs={tabContents} /> */}
                            <TabEventContext.Provider value={{ onTabSwitch: (oldTabKey, newTabKey) => { setCurrentTabKey(newTabKey) } }}>
                                <TabContext.Provider value={{ ...customTabContextvalue, currentTab: currentTabKey, }}>
                                    <Tab initTabKey={currentTabKey} tabs={pagePhases} />
                                    <br />

                                </TabContext.Provider>
                            </TabEventContext.Provider>
                            <NextPhaseButton resolveNextPhaseKey={resolveNextPhaseKey} />
                        </div>

                    </PageControllerContext.Provider>
                </FormCollector>
            </FormCollectorBus>
            {/* </FormCollectorDispatchContext.Provider> */}
        </>
    )
}