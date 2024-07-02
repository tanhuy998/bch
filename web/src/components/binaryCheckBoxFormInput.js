import { useDataModelBinding } from "./formInput"

export default function BinaryCheckBoxFormInput({name, className}) {

    const [checked, setChecked] = useDataModelBinding(name, "checkbox", undefined, false);

    return (
        <input className={className} type="checkbox" value={checked} onChange={() => {setChecked(!checked)}}></input>
    )
}