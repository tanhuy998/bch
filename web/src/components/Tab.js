import { memo, useEffect, useState } from "react";

function TabItem({ setCurrentTab, label, isActive }) {

    const elementClass = `nav-link ${isActive? 'active' : ''}`;

    return (
        <>
            <li class="nav-item" role="presentation">
                <button onClick={() => { !isActive && setCurrentTab(label) }} class={elementClass} id="pills-home-tab" data-bs-toggle="pill" href="#pills-home" role="tab" aria-controls="pills-home" aria-selected="false" tabindex="-1">{label}</button>
            </li>
        </>
    )
}

export default memo(Tab)

function Tab({ tabs }) {

    if (typeof tabs !== 'object') {

        throw new Error('There are no tab passed to tab list')
    }

    const tabKeys = Object.keys(tabs);
    const [currentTabKey, setCurrentTabKey] = useState(tabs[tabKeys[0]]);

    useEffect(() => {


    }, [currentTabKey])

    return (
        <>
            <ul class="nav nav-pills mb-3" id="pills-tab" role="tablist">
                {
                    Object.keys(tabs || {}).map((key) => {

                        return (
                            <TabItem setCurrentTab={setCurrentTabKey} label={key} isActive={key === currentTabKey} />
                        )
                    })
                }
            </ul>
            <div class="tab-content" id="pills-tabContent">
                {tabs?.[currentTabKey]}
            </div>

           
        </>
    )
}