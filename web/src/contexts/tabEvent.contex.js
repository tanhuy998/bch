import { createContext } from "react";

const TabEventContext = createContext({
    onTabSwitch: (oldTabKey, newTabKey) => { },
});

export default TabEventContext;