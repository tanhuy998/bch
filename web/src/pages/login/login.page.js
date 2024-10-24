import '../../assets/vendor/bootstrap/css/bootstrap.min.css'
import '../../assets/css/error.css'
import './style.css'

import { useState, useEffect } from 'react'
import Form from '../../components/form';
import FormInput from '../../components/formInput';
import PromptFormInput from '../../components/promptFormInput';
import LoginUseCase from './usecase';
import { useNavigate } from 'react-router-dom';
import { useRedirectAdmin, useRedirectNavigateTenant } from '../../hooks/authentication';


function required(val) {

    return typeof val === "string" && val != ""
}

export default function Login({usecase}) {
 
    useRedirectAdmin()

    const [isLoggedIn, setIsLoggedIn] = useState(null);
    const navigate = useNavigate();

    if (!(usecase instanceof LoginUseCase)) {

        throw new Error("invalid usecase passed to Login page")
    }

    useEffect(() => {

        usecase.endpoint.isLoggedIn()
        .then((state) => {

            setIsLoggedIn(state)
        })
    }, []);

    useEffect(() => {

        if (isLoggedIn) {

            navigate("/auth/nav")
        }

    }, [isLoggedIn]);


    if (isLoggedIn || isLoggedIn == null) {

        return  <></>
    }

    return (
        <>
            <div class="container">
                <div class="row">
                    <div class="col-md-6 offset-md-3">

                        <div class="card my-5">
                            <Form className="card-body cardbody-color p-lg-5" delegate={usecase}>
                            {/* <form class="card-body cardbody-color p-lg-5"> */}

                                <div class="text-center">
                                    <img src="https://cdn.pixabay.com/photo/2016/03/31/19/56/avatar-1295397__340.png" class="img-fluid profile-image-pic img-thumbnail rounded-circle my-3"
                                        width="200px" alt="profile" />
                                </div>

                                <div class="mb-3">
                                    <PromptFormInput name="username" validate={required} invalidMessage="Username must not be empty" className="form-control" placeholder="User Name" />
                                    {/* <input type="text" class="form-control" id="Username" aria-describedby="emailHelp"
                                        placeholder="User Name" /> */}
                                </div>
                                <div class="mb-3">
                                    <PromptFormInput name="password" validate={required} invalidMessage="password must not be empty" type="password" className="form-control" placeholder="Password" />
                                    {/* <input type="password" class="form-control" id="password" placeholder="password" /> */}
                                </div>
                                <div class="text-center">           
                                    <button type="submit" class="btn btn-color px-5 mb-5 w-100">Login</button>
                                </div>
                                {/* <div id="emailHelp" class="form-text text-center mb-5 text-dark">Not
                                    Registered? <a href="#" class="text-dark fw-bold"> Create an
                                        Account</a>
                                </div> */}
                            {/* </form> */}
                            </Form>
                        </div>

                    </div>
                </div>
            </div>
        </>

    )
}