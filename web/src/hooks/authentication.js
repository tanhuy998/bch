import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import { ACCESS_TOKEN_KEY } from "../const"
import HttpEndpoint from "../backend/endpoint"
import { fetch_options } from "../domain/models/fetchOption.model"

const adminPath = "/admin"
const loginPath = "/login"
const switchTenantPath = "/auth/switch"

const endpoint = new HttpEndpoint({})

export function useRedirectAdmin() {

    const navigate = useNavigate()

    useEffect(() => {
        console.log("auth 1")
        if (typeof localStorage.getItem(ACCESS_TOKEN_KEY) === "string") {

            navigate(adminPath)
        }
        
    }, [])
}

/**
 * 
 * @returns {boolean|null}
 */
export function useRedirectSwitchTenant() {

    const [isWaiting, setIsWaiting] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        console.log("auth 2")

        const opts = new fetch_options

        opts.method = "HEAD"

        endpoint.fetchRaw(
            opts,
            undefined,
            '/auth/login',
        ).then((res) => {

            switch (res.status) {
                case 204:
                    setIsWaiting(false);
            }
        }).catch(err => {

            alert(err?.message);
        })

    }, [])

    useEffect(() => {

        if (isWaiting == false) {

            navigate(switchTenantPath);
        }

    }, [isWaiting])

    return isWaiting;
}

export default function useAthentication() {

    const [state] = useState(null);
    const navigate = useNavigate();

    useRedirectAdmin();
    const isWaiting =  useRedirectSwitchTenant();

    useEffect(() => {
        console.log("auth 3")

        if (isWaiting == true) {

            return
        }

        navigate(loginPath)

    }, [isWaiting])

    return state;
}