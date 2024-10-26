
import '../assets/vendor/fontawesome/css/fontawesome.min.css';
import '../assets/vendor/fontawesome/css/solid.min.css';
import '../assets/vendor/fontawesome/css/brands.min.css';
import '../assets/vendor/bootstrap/css/bootstrap.min.css';
import '../assets/css/master.css';
import React, { Component, useEffect, useState } from "react";
import SideBar from "../components/sidebar";
import NavBar from "../components/navbar";
import { Link, Outlet, useLocation, useNavigate, useOutlet } from 'react-router-dom';
import { Provider } from 'react-redux';
import { AnimatePresence } from 'framer-motion';
import useAthentication from '../hooks/authentication';


export default function AdminTemplate(props) {

    const isWaitingAuthentication = useAthentication();

    const [sideBarActive, setSideBarActive] = useState("");
    const location = useLocation();
    const element = useOutlet();
    const navigate = useNavigate();


    if (isWaitingAuthentication) {

        return <></>
    }

    return (
        <div class="wrapper">
            <SideBar sideBarActive={sideBarActive} />
            <div id="body" class="active">
                <NavBar toggleSideBar={() => { setSideBarActive(sideBarActive === "active" ? "" : "active") }} />
                <div class="content">
                    <div class="container">
                        <AnimatePresence mode='wait' initial={true}>
                            <button onClick={() => { navigate(-1) }} style={{marginLeft: 5}} class="btn btn-sm btn-outline-secondary float-end"><i class="fas fa-arrow-left"></i> Quay lại</button>
                            {/* <Outlet key={location.pathname}/> */}
                            {element && React.cloneElement(element, {key: location.pathname})}
                        </AnimatePresence>
                    </div>
                </div>
            </div>
        </div>
    )
}