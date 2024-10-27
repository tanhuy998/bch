import { useEffect, useState } from "react"
import { useLocation, useNavigate } from "react-router-dom"
import { getAccessToken, getAccessTokenExp, removeAccessToken, setAccessToken, setUserInfo } from "../lib/authSignature.lib";
import AuthEndpoint from "../backend/autEndpoint";
import UserSessionExpireError from "../backend/error/userSessionExpireError";

const admin_path = "/admin";
const login_path = "/login";

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

export function useRedirectAuthReferer() {

    const isRotating = useAccessToken();
    const navigate = useNavigate();

    useEffect(() => {

        if (isRotating) {

            return;
        }

        const hasToken = typeof getAccessToken() === 'string';
        const isExpire = isTokenExpires();

        if (hasToken && !isExpire) {
            
            navigate(
                typeof referer_path === 'string' ? referer_path || admin_path : admin_path,
            );
            return;
        }
    })

    return isRotating;
}

/**
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
        
        referer_path = location.pathname;
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