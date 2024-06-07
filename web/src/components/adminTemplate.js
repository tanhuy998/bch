
import '../assets/vendor/fontawesome/css/fontawesome.min.css';
import '../assets/vendor/fontawesome/css/solid.min.css';
import '../assets/vendor/fontawesome/css/brands.min.css';
import '../assets/vendor/bootstrap/css/bootstrap.min.css';
import '../assets/css/master.css';
import { Component } from "react";
import SideBar from "../components/sidebar";
import NavBar from "../components/navbar";
import { Outlet } from 'react-router-dom';



export default class AdminTemplate extends Component {

    
    constructor(props) {

        super(props)

        this.renderMainContent = this.renderMainContent.bind(this)
    }

    renderMainContent() {

        const renderContent = this.props.renderContent

        return typeof renderContent === 'function' ? renderContent : () => {}
    }

    render() {

        return (
            <div class="wrapper">
                <SideBar />
                <div id="body" class="active">
                    <NavBar />
                    <div class="content">
                        {this.renderMainContent()()}
                        <Outlet />
                    </div>
                </div>
            </div>
        )
    }
}