import { configureStore } from "@reduxjs/toolkit";
import toogleSideBarReducer from "./states/toggleSideBar"

export default configureStore({
    reducer: {
        toggleSidebar: toogleSideBarReducer
    }
})