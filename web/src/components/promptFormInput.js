import { useEffect, useState } from "react";
import FormInput, {_FormInput} from "./formInput";
import SelectInput from "./selectInput";


export default function PromptFormInput({label, validate, name, placeholder, noticeText, invalidMessage, type, textArea, defaultValue, ...rest}) {

    const [isValidInput, setIsValidInput] = useState();
    
    useEffect(() => {

        console.log('prompt', isValidInput, invalidMessage)

    }, [isValidInput])

    return (
        <>
            {label && <label for={name} class="form-label">{label}</label>}
            <FormInput {...rest} defaulValue={defaultValue} validate={validate} type={type} onValidInput={() => {setIsValidInput(true)}} onInvalidInput={() => {setIsValidInput(false)}} className="form-control" name={name} placeholder={placeholder} textArea={textArea}/>
            <small class="form-text text-muted">{noticeText}</small>

            {/* <div class="valid-feedback">Looks good!</div> */}
            {
                !isValidInput && <div class="invalid-feedback is-invalid">{invalidMessage}</div>
            }
        </>
    )
}

export function PromptSelectInput({defaultValue, label, name, placeholder, noticeText, invalidMessage, children, required, className, castedType , notNull, ...rest }) {

    const [isSelected, setIsSelected] = useState(false);

    return (
        <>
            {label && <label for={name} class="form-label">{label}</label>}
            <SelectInput 
                {...rest}
                className={className + (required && !isSelected ? 'is-invalid' : '' )} 
                name={name} 
                onUnSelected={() => {setIsSelected(false)}} 
                onSelected={() => {setIsSelected(true)}}
                castedType={castedType}
                notNull={notNull}
                defaultValue={defaultValue}
            >
                {children}
            </SelectInput>
            <small class="form-text text-muted">{noticeText}</small>
            {
                required && !isSelected && <div class="invalid-feedback is-invalid">{invalidMessage}</div>
            }
            <br/>
        </>
    )
}

export function PrompEthinicitySelectInput({className, ...rest}) {

    return (
        <PromptSelectInput
            {...rest}
            label="Dân Tộc"
            name="ethnicity"
            invalidMessage="Dân tộc không được bỏ trống"
            required={true}
            className={className}
        >
            <option value="Kinh">Kinh</option>
            <option value="Hoa">Hoa</option>
            <option value="Chăm">Chăm</option>
            <option value="Khmer">Khmer</option>
        </PromptSelectInput>
    )
}