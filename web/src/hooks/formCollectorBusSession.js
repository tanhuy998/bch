import { useContext } from "react";
import FormCollectorBusContext from "../contexts/formCollectorBus.context";

export default function useFormCollectorBusSession() {

    const ctx = useContext(FormCollectorBusContext);

    return typeof ctx === 'object' ? ctx.busSession : false;
}