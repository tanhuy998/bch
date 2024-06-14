import TabContext, { pillTabStyle } from "../contexts/tab.context";
import Tab from "./Tab";

export default function PillTab({ tabs,  initTabKey, initTabIndex  }) {

    return (
        <TabContext.Provider value={pillTabStyle}>
            <Tab tabs={tabs} initTabIndex={initTabIndex} initTabKey={initTabKey}/>
        </TabContext.Provider>
    )
}