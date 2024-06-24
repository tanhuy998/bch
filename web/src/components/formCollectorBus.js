import { useState } from "react";
import FormCollectorDispatchContext from "../contexts/formCollectorDispatch.context";

export default function FormCollectorBus({children}) {

    const [formCollectorHandShake, setFormCollectorHandShake] = useState(false);
    const [formCollectorResponse, setFormCollectorResponse] = useState();
    const [emitSignal, setEmitSignal] = useState();

    return (
        <FormCollectorDispatchContext.Provider
            value={{
                emitSignal: null,
                setEmitSignal,
                formCollectorResponse,
                setFormCollectorResponse,
                formCollectorHandShake,
                setFormCollectorHandShake,
            }}
        >
            {children}
        </FormCollectorDispatchContext.Provider>
    )
}