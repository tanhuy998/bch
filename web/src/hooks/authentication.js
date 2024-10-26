import { useEffect, useState } from "react"
import { useLocation, useNavigate } from "react-router-dom"
import HttpEndpoint from "../backend/endpoint"
import { fetch_options } from "../domain/models/fetchOption.model"
import { getAccessToken, getAccessTokenExp, removeAccessToken, setAccessToken, setUserInfo } from "../lib/authSignature.lib";
import AuthEndpoint from "../backend/autEndpoint";
import UserSessionExpireError from "../backend/error/userSessionExpireError";

const adminPath = "/admin";
const loginPath = "/login";
const switchTenantPath = "/auth/switch";

const endpoint = new HttpEndpoint({});


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

            navigate(adminPath);
            return;
        }

    }, [isRotating])

    return isRotating;
}

// /**
//  * 
//  * @param {boolean} waitForAnotherHook
//  * @returns {boolean|null}
//  */
// export function useRedirectNavigateTenant(waitForAnotherHook) {

//     const [isWaiting, setIsWaiting] = useState(true);
//     const navigate = useNavigate();

//     useEffect(() => {

//         if (waitForAnotherHook === true) {

//             return
//         }

//         const opts = new fetch_options;
//         opts.method = 'HEAD';

//         endpoint.fetchRaw(
//             opts,
//             undefined,
//             '/auth/login',
//         ).then((res) => {

//             switch (res.status) {
//                 case 204:
//                     //setIsWaiting(false);
//                     navigate('/auth/nav');
//                 case 401:
//                     setIsWaiting(false);
//             }
//         }).catch(err => {

//             alert(err?.message);
//         })

//     }, [waitForAnotherHook]);

//     return isWaiting;
// }

// /**
//  *  @returns {[getUserInfo, setUserInfo]}
//  */
// export function useUserInfo() {

//     return [useUserInfo, setUserInfo];
// }

/**
 * 
 * @returns {boolean}
 */
export default function useAuthentication() {

    const location = useLocation();
    const navigate = useNavigate();
    const isRotatingToken = useAccessToken()

    useEffect(() => {

        if (isRotatingToken) {

            return;
        }

        if (hasToken()) {

            return
        }

        if (location.pathname == loginPath) {

            return;
        }
        
        navigate(loginPath);

    }, [isRotatingToken]);

    return isRotatingToken;
}

export function useAccessToken(waitForAnotherHook) {

    const [isPending, setIsPending] = useState(true);

    useEffect(() => {

        if (waitForAnotherHook) {

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

                AuthEndpoint.rotateSignatures()                
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

        // AuthEndpoint.rotateSignatures()
        //     .then(() => {

        //         setIsPending(false);
        //     })
        //     .catch((e) => {

        //         if (e instanceof UserSessionExpireError) {

        //             removeAccessToken();
        //             setIsPending(false);
        //             return;
        //         }

        //         setIsPending(false);
        //         alert(e?.message ?? e);
        //     });

    }, [waitForAnotherHook]);

    return isPending;
}