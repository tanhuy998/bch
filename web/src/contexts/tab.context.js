import { createContext } from "react";

export const pillTabStyle = {
    ul: {
        "class": "nav nav-pills mb-3",
        "id": "pills-tab",
        "role": "tablist"
    },
    li: {
        //"class": "nav-item",
        "role": "presentation"
    },
    liButton: {
        "id": "pills-home-tab",
        "data-bs-toggle": "pill",
        "href": "#pills-home",
        "role": "tab",
        "aria-controls": "pills-home",
        "aria-selected": "false",
        "tabindex": "-1",
    }
}

export const basicTabStyle =  {
    ul: {
        "class": "nav nav-tabs",
        "id": "myTab",
        "role": "tablist",
    },
    li: {
        "class": "nav-item",
        "role": "presentation"
    },
    liButton: {
        //"id": "pills-home-tab",
        "data-bs-toggle": "tab",
        "href": "#contact",
        "role": "tab",
        "aria-controls": "contact",
        "aria-selected": "false",
        "tabindex": "-1",
    }
}

const TabContext = createContext();

export default TabContext;