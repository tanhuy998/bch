import {createSlice, configureStore} from '@reduxjs/toolkit';

const toggleSideBarSlice = createSlice({
    name: 'toggleSideBar',
    initialState: {
        value: ""
    },
    reducers: {
        toggleSideBar: state => {

            state.value = state.value == "active" ? "" : "active"
        }
    }
})

export const {toggleSideBar} = toggleSideBarSlice.actions

export default toggleSideBarSlice.reducer