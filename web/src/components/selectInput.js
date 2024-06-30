import { useState } from "react";
import { useDataModelBinding } from "./formInput";

const FIRST = 0;

export default function SelectInput({children, defaultValue, name}) {

    const allProps = arguments[FIRST];
    
    const [selectedValue, setSelectedValue] = useDataModelBinding(name, 'select');

    const handleChange = (event) => {

        setSelectedValue(event.target.value);
    }

    return (
        <select {...{...allProps, children: undefined}} onChange={handleChange} >
            {children}
        </select>
    )
}
