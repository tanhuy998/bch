import TabContext, { basicTabStyle } from "../contexts/tab.context";
import Tab from "./Tab";

export default function BasicTab({ tabs,  initTabKey, initTabIndex  }) {

    return (
        <TabContext.Provider value={basicTabStyle}>
            <Tab tabs={tabs} initTabIndex={initTabIndex} initTabKey={initTabKey}/>
        </TabContext.Provider>
    )
}