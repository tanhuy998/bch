import { memo, useContext, useEffect, useRef, useState } from "react";
import TabContext from "../contexts/tab.context";
import { AnimatePresence } from "framer-motion";
import {motion} from 'framer-motion';
import useResizeObserver from "../hooks/resizeObserver.hook";

const tabContainerStyle = {
    transition: "height 0.5 ease-out"
}

function TabItem({ setCurrentTab, label, isActive }) {
    const tabContext = useContext(TabContext);
    const elementClass = `nav-link ${isActive ? 'active' : ''}`;

    return (
        <>
            <li {...(tabContext?.li || {})}>
                <button onClick={() => { !isActive && setCurrentTab(label) }} class={elementClass} {...tabContext?.liButton || {}}>{label}</button>
            </li>
        </>
    )
}

/**
 * 
 * @param {Array} tabKeys 
 * @param {int} initIndex 
 * @param {string} initKey 
 */
function resovleExactInitTabKey(tabKeys, initIndex, initKey) {

    if (!initKey && !initIndex) {

        return undefined;
    }

    if (
        typeof initKey === 'string'
        && tabKeys.includes(initKey)
    ) {

        return initKey;
    }

    if (initKey < tabKeys.length) {

        return tabKeys[initIndex];
    }

    throw new Error('tab initKey and initIndex are invalid');
}

export default memo(Tab)

function Tab({ tabs, initTabKey, initTabIndex }) {
    const tabContext = useContext(TabContext);
    const tabContainer = useRef();
    const [tabContainerHeight, setTabContainerHeight] = useState(0);
    const tabContainerRect = useResizeObserver(tabContainer);
    const [containerOpacity, setContainerOpacity] = useState(1);

    if (typeof tabs !== 'object') {

        throw new Error('There are no tab passed to tab list')
    }

    const tabKeys = Object.keys(tabs);
    const [currentTabKey, setCurrentTabKey] = useState();

    const initKey = resovleExactInitTabKey(tabKeys, initTabIndex, initTabKey);

    useEffect(() => {

        setContainerOpacity(1);

    }, [currentTabKey])

    useEffect(() => {

        if (!initKey) {

            return;
        }

        setCurrentTabKey(initKey);
    }, [])

    // useEffect(() => {

    //     if (tabContainerHeight === tabContainerRect.height) {

    //         setContainerOpacity(1);
    //     }

    //     setTabContainerHeight(tabContainerRect.height);

    // }, [tabContainerRect])

    // useEffect(() => {
    //     console.log('height', (tabContainer.current.getBoundingClientRect().height))
    //     setTabContainerHeight(tabContainer.current.getBoundingClientRect().height)
    // })

    return (
        <div>
            <ul {...(tabContext?.ul || {})}>
                {
                    Object.keys(tabs || {}).map((key) => {

                        return (
                            <TabItem setCurrentTab={setCurrentTabKey} label={key} isActive={key === currentTabKey} />
                        )
                    })
                }
            </ul>
            <AnimatePresence>
                {
                    (
                        <div 
                            
                            style={{
                                transition: "all 0.3s",
                                opacity: containerOpacity,
                                height: `${tabContainerRect.height}px`,
                                width: `${tabContainerRect.width}px`,
                                overflow: "hidden",
                            }}
                        >
                            <div {...(tabContext?.content || {})} ref={tabContainer}>
                                    {tabs?.[currentTabKey]}
                                </div>
                        </div>
                    )
                }
            </AnimatePresence>
        </div>
    )
}