
import '../assets/vendor/fontawesome/css/fontawesome.min.css';
import '../assets/vendor/fontawesome/css/solid.min.css';
import '../assets/vendor/fontawesome/css/brands.min.css';
import '../assets/vendor/bootstrap/css/bootstrap.min.css';
import '../assets/css/master.css';
import React, { Component, useEffect, useState } from "react";
import SideBar from "../components/sidebar";
import NavBar from "../components/navbar";
import { Outlet, useLocation, useOutlet } from 'react-router-dom';
import { Provider } from 'react-redux';
import { AnimatePresence } from 'framer-motion';


export default function AdminTemplate(props) {

    const [sideBarActive, setSideBarActive] = useState("");
    const location = useLocation();
    const element = useOutlet();

    return (
        <div class="wrapper">
            <SideBar sideBarActive={sideBarActive} />
            <div id="body" class="active">
                <NavBar toggleSideBar={() => { setSideBarActive(sideBarActive === "active" ? "" : "active") }} />
                <div class="content">
                    <div class="container">
                        <AnimatePresence mode='wait' initial={true}>
                            {/* <Outlet key={location.pathname}/> */}
                            {element && React.cloneElement(element, {key: location.pathname})}
                        </AnimatePresence>
                    </div>
                </div>
            </div>
        </div>
    )
}