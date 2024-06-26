import { useContext } from "react";
import FormCollectorBusContext from "../contexts/formCollectorBus.context";

export default function useCollectedForms() {

    const busContext = useContext(FormCollectorBusContext);



    return typeof busContext === 'object' ? [busContext.collectedEndSignal, busContext.collectedDelegator] : false;
}