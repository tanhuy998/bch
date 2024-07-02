import { useEffect, useState } from "react";
import { useDataModelBinding } from "./formInput";

const FIRST = 0;

export default function SelectInput({ children, defaultValue, name, onUnSelected, onSelected, castedType, notNull}) {

    const allProps = arguments[FIRST];
    
    const [selectedValue, setSelectedValue] = useDataModelBinding(name, 'select', castedType, defaultValue);
    const hasUnSelectedListener = typeof onUnSelected === 'function';
    const hasSelectedListener = typeof onSelected === 'function';

    const handleChange = (event) => {

        setSelectedValue(event.target.value);
    }
    
    useEffect(() => {
        
        if (
            selectedValue === ""
            && hasUnSelectedListener
        ) {

            onUnSelected();
            return;
        }

        if (
            selectedValue != ""
            && hasSelectedListener
        ) {

            onSelected(selectedValue);
            return;
        }

    }, [selectedValue]) 

    return (
        <select defaultValue={defaultValue} {...{...allProps, children: undefined}} value={selectedValue} onInput={handleChange} >
            {!notNull && <option selected></option>}
            {children}
        </select>
    )
}


