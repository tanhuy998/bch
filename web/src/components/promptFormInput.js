import { useEffect, useState } from "react";
import FormInput, {_FormInput} from "./formInput";


export default function PromptFormInput({label, validate, name, placeholder, noticeText, invalidMessage, type, textArea}) {

    const [isValidInput, setIsValidInput] = useState();
    
    useEffect(() => {

        console.log('prompt', isValidInput, invalidMessage)

    }, [isValidInput])

    return (
        <>
            <label for={name} class="form-label">{label}</label>
            <FormInput validate={validate} type={type} onValidInput={() => {setIsValidInput(true)}} onInvalidInput={() => {setIsValidInput(false)}} className="form-control" name={name} placeholder={placeholder} textArea={textArea}/>
            <small class="form-text text-muted">{noticeText}</small>

            {/* <div class="valid-feedback">Looks good!</div> */}
            {
                !isValidInput && <div class="invalid-feedback is-invalid">{invalidMessage}</div>
            }
        </>
    )
}