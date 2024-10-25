import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import HttpEndpoint from "../backend/endpoint"
import { fetch_options } from "../domain/models/fetchOption.model"

const adminPath = "/admin";
const loginPath = "/login";
const switchTenantPath = "/auth/switch";

const endpoint = new HttpEndpoint({});
const ACCESS_TOKEN_KEY = "at";
const USER_INFO_KEY = "ui";

export function useRedirectAdmin() {

    const navigate = useNavigate()

    useEffect(() => {
        console.log("auth 1");
        if (typeof localStorage.getItem(ACCESS_TOKEN_KEY) === "string") {

            navigate(adminPath);
        }
    })
}

/**
 * 
 * @returns {boolean|null}
 */
export function useRedirectNavigateTenant() {

    const [isWaiting, setIsWaiting] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        console.log("auth 2");

        const opts = new fetch_options;

        opts.method = "HEAD";

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

/**
 *  @returns {[getAccessToken, setAccessToken]}
 */
export function useAccessToken() {

    const [state, setState] = useState();

    /**
     * @name getAccessToken
     * @function
     * @returns {string|null}
     */
    function getAccessToken() {

        return localStorage.getItem(ACCESS_TOKEN_KEY);
    }

    /**
     * @name setAccessToken
     * @function
     * @param {string} accessToken 
     */
    function setAccessToken(accessToken) {

        localStorage.setItem(ACCESS_TOKEN_KEY, accessToken);
        setState(accessToken);
    }

    return [getAccessToken, setAccessToken];
}



/**
 *  @returns {[getUserInfo, setUserInfo]}
 */
export function useUserInfo() {

    /**
     *  @typedef {Object} UserInfo
     *  @property {string} uuid
     *  @property {string} name
     *  @property {string} username
     *  @property {string} tenantUUID
     */

    /**
     * @name getUserInfo
     * @function
     * @returns {UserInfo|undefined}
     */
    function getUserInfo() {

        const ret = localStorage.getItem(USER_INFO_KEY);

        return typeof ret === "object" && ret !== null ? ret : undefined
    }

    /**
     * @name setUserInfo
     * @function
     * @param {UserInfo} data 
     */
    function setUserInfo(data) {

        localStorage.setItem(USER_INFO_KEY, data);
    }

    return [getUserInfo, setUserInfo];
}

export default function useAthentication() {

    const [state] = useState(null);
    const navigate = useNavigate();

    useRedirectAdmin();
    const isWaiting = useRedirectNavigateTenant();

    useEffect(() => {
        console.log("auth 3")

        if (isWaiting == true) {

            return;
        }

        navigate(loginPath);

    }, [isWaiting]);

    return state;
}