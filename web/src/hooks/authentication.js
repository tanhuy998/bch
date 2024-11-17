import { useEffect, useState } from "react"
import { useLocation, useNavigate } from "react-router-dom"
import { getAccessToken, getAccessTokenExp, removeAccessToken, setAccessToken, setUserInfo } from "../lib/authSignature.lib";
import AuthEndpoint from "../backend/autEndpoint";
import UserSessionExpireError from "../backend/error/userSessionExpireError";

const admin_path = "/admin";
const login_path = "/login";
const regex_admin_path = /^\/admin(\/(.)+)*/

let referer_path;

/**
 * @returns {boolean}
 */
export function isTokenExpires() {

    const accessToken = getAccessToken();
    const tokenExp = getAccessTokenExp();

    return typeof accessToken === "string" 
            && typeof Number(tokenExp) === "number"
            && (Date.now() - Number(tokenExp)) >= 0;
}

export function hasToken() {

    return typeof getAccessToken() === "string";
}

/**
 * redirect to the last page that used authentication,
 * if there is no access token, redirect back to login page.
 * 
 * 
 * @returns 
 */
export function useRedirectAdmin() {

    const isRotating = useAccessToken();
    const navigate = useNavigate();

    useEffect(() => {

        if (isRotating) {

            return;
        }

        const hasToken = typeof getAccessToken() === 'string';
        const isExpire = isTokenExpires();

        if (hasToken && !isExpire) {
            
            navigate(admin_path);
            return;
        }
    })

    return isRotating;
}

/**
 *  read current access token, if no access token, redirect to login page
 *  and vice versa if the whole user session expired.
 *  If access token expired, rotate the signatures.
 * 
 * @returns {boolean}
 */
export default function useAuthentication() {
    
    const location = useLocation();
    const navigate = useNavigate();
    const isRotatingToken = useAccessToken()
    referer_path = undefined;

    useEffect(() => {

        if (isRotatingToken) {

            return;
        }

        if (hasToken()) {
            
            return
        }

        if (location.pathname === login_path) {

            return;
        }

        navigate(login_path);

    }, [isRotatingToken]);

    return isRotatingToken;
}

/**
 * 
 * @param {boolean} waitForAnotherHook 
 * @returns {boolean}
 */
export function useAccessToken(waitForAnotherHook) {

    const [isPending, setIsPending] = useState(!waitForAnotherHook);

    useEffect(() => {

        if (waitForAnotherHook) {

            return;
        }

        setIsPending(true);

    }, [waitForAnotherHook]);

    useEffect(() => {

        if (!isPending) {

            return;
        }

        const hasToken = typeof getAccessToken() === 'string';
        const isExpire = isTokenExpires();

        if (!hasToken) {

            setIsPending(false);
            return
        }

        if (hasToken && !isExpire) {

            setIsPending(false);
            return;
        }

        (async () => {

            try {

                await AuthEndpoint.rotateSignatures()
            }
            catch (e) {

                if (e instanceof UserSessionExpireError) {

                    removeAccessToken();
                }
                else {

                    alert(e?.message ?? e);
                }
            }
            finally {

                setIsPending(false);
            }
        })();     

    }, [isPending]);

    return waitForAnotherHook || isPending;
}