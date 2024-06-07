import AdminTemplate from "../components/adminTemplate";
import {Component} from 'react';

export default class AdminHomePage extends Component {

    constructor(props) {

        super(props)
    }

    componentDidMount() {

        console.log(this.state)
    }

    render() {

        return (
            <AdminTemplate/>
        )
    }
}