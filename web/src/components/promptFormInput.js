import FormInput from "./formInput";

export default function PromptFormInput({label, inputName, placeholder, noticeText, invalidMessage, type}) {

    return (
        <>
            <label for={inputName} class="form-label">{label}</label>
            <FormInput type={type} className="form-control" name={inputName} placeholder={placeholder} />
            <small class="form-text text-muted">{noticeText}</small>
            {/* <div class="valid-feedback">Looks good!</div> */}
            <div class="invalid-feedback">{invalidMessage}</div>
        </>
    )
}