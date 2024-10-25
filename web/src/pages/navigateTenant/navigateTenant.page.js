import { useEffect, useState } from "react";
import SwitchTenantUseCase from "./usecase";
import ErrorResponse from "../../backend/error/errorResponse";
import { useNavigate } from "react-router-dom";
import { useAccessToken, useRedirectAdmin, useRedirectLogin } from "../../hooks/authentication";
import LoadingCircle from "../../components/loadingCircle";


/**
 * 
 * @param {*} param0 
 * @returns 
 */
export default function NavigateTenantPage({ usecase }) {

    useRedirectAdmin()

    const [tenantList, setTenantList] = useState(undefined);
    const [chosenTenantUUID, setChosenTenantUUID] = useState(undefined);
    const [getAccessToken, setAccessToken] = useAccessToken()
    const navigate = useNavigate()

    if (!(usecase instanceof SwitchTenantUseCase)) {

        throw new Error("invalid usecase passed to SitchTenantPage")
    }

    useEffect(() => {

        usecase.fetchUserTenants()
            .then((val) => {

                setTenantList(val?.data);
            })
            .catch((err) => {

                if (err instanceof ErrorResponse) {

                    navigate("/login")
                    return
                }

                alert(err?.message)
            });

    }, [])

    useEffect(() => {

        if (tenantList == undefined) {

            return
        }

        if (!Array.isArray(tenantList) || tenantList?.length == 0) {

            navigate('/login')
            return
        }

    }, [tenantList])

    useEffect(() => {

        if (
            !chosenTenantUUID
            || typeof chosenTenantUUID !== "string"
        ) {

            return;
        }

        usecase.switchToTenant(chosenTenantUUID)
            .then( data => {

                    if (data.status === 401) {

                        return
                    }

                    setAccessToken(data?.data?.accessToken)
            })
            .catch((err) => {

                setChosenTenantUUID(undefined)
                alert(err)
            });

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
                                                            {/* <div className="mb-3">
                                                <button class="btn btn-outline-primary shadow-2 mb-4">{tenant?.name}</button>
                                            </div> */}
                                                            <tr>
                                                                <td>{tenant?.name}</td>
                                                                <td>{tenant?.isTenantAgent ? 'Tenant agent' : 'Member'}</td>
                                                                {
                                                                    (!chosenTenantUUID) ?
                                                                        <td>
                                                                            <button className="btn btn-outline-primary" onClick={
                                                                                () => {
                                                                                    console.log(tenant?.name)
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