import { useEffect, useState } from "react";
import SwitchTenantUseCase from "./usecase";
import ErrorResponse from "../../backend/error/errorResponse";
import { useNavigate } from "react-router-dom";
import { useRedirectAdmin, useRedirectLogin } from "../../hooks/authentication";
import "./style.css"

export default function SwitchTenantPage({ usecase }) {

    useRedirectAdmin()

    const [tenantList, setTenantList] = useState(undefined);
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

    if (!Array.isArray(tenantList)) {

        return <></>
    }

    return (
        <>
        {
            tenantList.map((tenant) => {
                return (
                    <>
                        <button>{tenant?.name}</button>
                    </>
                )
            })
        }
        </>
    )
}