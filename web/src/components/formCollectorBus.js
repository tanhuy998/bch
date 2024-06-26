import { useState } from "react";
import FormCollectorBusContext from "../contexts/formCollectorBus.context";

export default function FormCollectorBus({children}) {

    // const [formCollectorHandShake, setFormCollectorHandShake] = useState(false);
    // const [formCollectorResponse, setFormCollectorResponse] = useState();
    // const [emitSignal, setEmitSignal] = useState();

    const [collectedDelegator, setCollectedDelegator] = useState();
    const [collectedEndSignal, setCollectedEndSignal] = useState(null);
    const [busSession, setBusSession] = useState(null);

    return (
        <FormCollectorBusContext.Provider
            value={{
                // emitSignal: null,
                // setEmitSignal,
                // formCollectorResponse,
                // setFormCollectorResponse,
                // formCollectorHandShake,
                // setFormCollectorHandShake,
                collectedDelegator, setCollectedDelegator,
                collectedEndSignal, setCollectedEndSignal,
                busSession, setBusSession,
            }}
        >
            {children}
        </FormCollectorBusContext.Provider>
    )
}