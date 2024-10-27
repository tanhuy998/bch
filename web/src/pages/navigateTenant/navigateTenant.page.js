import { useEffect, useState } from "react";
import SwitchTenantUseCase from "./usecase";
import ErrorResponse from "../../backend/error/errorResponse";
import { useNavigate } from "react-router-dom";
import { useRedirectAuthReferer } from "../../hooks/authentication";
import LoadingCircle from "../../components/loadingCircle";
import { setAccessToken, setUserInfo } from "../../lib/authSignature.lib";


/**
 * 
 * @param {*} param0 
 * @returns 
 */
export default function NavigateTenantPage({ usecase }) {
    
    const isRotatingToken = useRedirectAuthReferer();

    const [tenantList, setTenantList] = useState(undefined);
    const [chosenTenantUUID, setChosenTenantUUID] = useState(undefined);

    const navigate = useNavigate();

    if (!(usecase instanceof SwitchTenantUseCase)) {

        throw new Error("invalid usecase passed to SitchTenantPage")
    }

    useEffect(() => {

        if (isRotatingToken) {

            return;
        }

        if (Array.isArray(tenantList)) {

            return;
        }

        usecase.fetchUserTenants()
            .then((val) => {

                setTenantList(val?.data);
            })
            .catch((err) => {

                if (err instanceof ErrorResponse) {

                    navigate("/login")
                    return;
                }

                alert(err?.message);
            });
        
    }, [isRotatingToken]);

    useEffect(() => {

        if (tenantList == undefined) {

            return;
        }

        if (!Array.isArray(tenantList) || tenantList?.length == 0) {

            navigate('/login');
            return;
        }

    }, [tenantList]);

    useEffect(() => {

        if (
            !chosenTenantUUID
            || typeof chosenTenantUUID !== "string"
        ) {

            return;
        }

        (async () => {

            try {

                const body = await usecase.switchToTenant(chosenTenantUUID);

                // if (data.status === 401) {

                //     return;
                // }

                setUserInfo(body?.data?.user);
                setAccessToken(body?.data?.accessToken);
            }
            catch (e) {

                alert(e?.message || e);
            }
            finally {

                setTenantList(undefined);
            }
        })()

        // usecase.switchToTenant(chosenTenantUUID)
        //     .then( data => {

        //             if (data.status === 401) {

        //                 return;
        //             }
                    
        //             setUserInfo(data?.data?.user)
        //             setAccessToken(data?.data?.accessToken)
        //             setTenantList(undefined);
        //     })
        //     .catch((err) => {

        //         setChosenTenantUUID(undefined);
        //         alert(err);
        //     });

    }, [chosenTenantUUID])

    if (!Array.isArray(tenantList)) {

        return <></>
    }

    return (
        <>

            <div class="wrapper">
                <div class="auth-content">
                    <div class="card">
                        <div class="card-body text-center">
                            <div class="mb-4">
                                <img class="brand" src="assets/img/bootstraper-logo.png" alt="bootstraper logo" />
                            </div>
                            <h5 class="mb-4 text-muted">Choose a tenant</h5>

                            <>
                                {
                                    /////////////////////////////////////////////////////////////////////////////
                                }
                            </>

                            <div id="t-tenant-display">
                                <table>
                                    <thead>
                                        <tr>
                                            <th>Tenant</th>
                                            <th>Tenant Role</th>
                                        </tr>
                                    </thead>
                                    <thead>
                                    </thead>
                                    <tbody>
                                        {
                                            tenantList.map(
                                                /**
                                                 * 
                                                 * @param {tenant_t} tenant 
                                                 * @returns 
                                                 */
                                                (tenant) => {
                                                    return (
                                                        <>

                                                            <tr>
                                                                <td>{tenant?.name}</td>
                                                                <td>{tenant?.isTenantAgent ? 'Tenant agent' : 'Member'}</td>
                                                                {
                                                                    (!chosenTenantUUID) ?
                                                                        <td>
                                                                            <button className="btn btn-outline-primary" onClick={
                                                                                () => {
                                                                                    
                                                                                    setChosenTenantUUID(tenant?.uuid)
                                                                                }
                                                                            }>
                                                                                CONNECT
                                                                            </button>
                                                                        </td>
                                                                        : chosenTenantUUID == tenant?.uuid ?
                                                                            <LoadingCircle />
                                                                            : <></>
                                                                }
                                                            </tr>
                                                        </>
                                                    )
                                                })
                                        }
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

        </>
    )
}