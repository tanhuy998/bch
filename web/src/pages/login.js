import '../assets/vendor/bootstrap/css/bootstrap.min.css'
import '../assets/css/error.css'
import { useState, useEffect } from 'react'

export default function Login() {

    const [submitValues, setSubmitValues] = useState({
        username: null,
        password: null
    });


    function handleSubmit(event) {

        console.log(event.target)        

        setSubmitValues(event.target)
    }

    useEffect(() => {

        console.log(submitValues)
    }, [submitValues]);

    return (
        <div class="wrapper">
            <div class="auth-content">
                <div class="card">
                    <div class="card-body text-center">
                        <div class="mb-4">
                            <img class="brand" src="../assets/img/bootstraper-logo.png" alt="bootstraper logo" />
                        </div>
                        <h6 class="mb-4 text-muted">Ban CHQS Phường 7 Quận 6</h6>
                        <form onSubmit={handleSubmit}>
                            <div class="mb-3 text-start">
                                <label for="username" class="form-label">Tài khoản</label>
                                <input type="text" class="form-control" placeholder="" required />
                            </div>
                            <div class="mb-3 text-start">
                                <label for="password" class="form-label">Mật khẩu</label>
                                <input type="password" class="form-control" placeholder="" required />
                            </div>
                            <br />
                            <button type="submit" class="btn btn-primary shadow-2 mb-4">Login</button>
                        </form> 
                    </div>
                </div>
            </div>
        </div>
    )
}