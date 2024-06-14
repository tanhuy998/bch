import { useContext } from "react";
import FormContext from "../contexts/form.context";

export default function FormInput({validator}) {

    const context = useContext(FormContext);
    const allProps = arguments[0];

    return (
        <input {...allProps} />
    )
}