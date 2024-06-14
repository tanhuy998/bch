import { memo, useContext, useEffect, useState } from "react";
import TabContext from "../contexts/tab.context";

function TabItem({ setCurrentTab, label, isActive }) {
    const tabContext = useContext(TabContext);
    const elementClass = `nav-link ${isActive? 'active' : ''}`;

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
    if (typeof tabs !== 'object') {

        throw new Error('There are no tab passed to tab list')
    }

    const tabKeys = Object.keys(tabs);
    const [currentTabKey, setCurrentTabKey] = useState();

    const initKey = resovleExactInitTabKey(tabKeys, initTabIndex, initTabKey);

    useEffect(() => {


    }, [currentTabKey])

    useEffect(() => {

        if (!initKey) {

            return;
        }

        setCurrentTabKey(initKey);
    }, [])

    return (
        <>
            <ul {...(tabContext?.ul || {})}>
                {
                    Object.keys(tabs || {}).map((key) => {

                        return (
                            <TabItem setCurrentTab={setCurrentTabKey} label={key} isActive={key === currentTabKey} />
                        )
                    })
                }
            </ul>
            <div {...(tabContext?.content || {})}>
                {tabs?.[currentTabKey]}
            </div>

           
        </>
    )
}