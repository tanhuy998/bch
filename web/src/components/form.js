import FormContext from "../contexts/form.context";

/**
 * This component is wrapper for default html <form> element. This component
 * construct a context for it's chilldren <FormInput> to consumes especially
 * for instantly form input validating when input value changed.
 * 
 * if handleFormData is funciton, it will be the interceptor that prevent
 * the default behavior of the form element.
 * 
 * @returns 
 */
export default function Form({ handleFormData, children, validateFunc }) {

    const allProps = arguments[0];

    function handleSubmit(e) {

        if (typeof handleFormData === 'function') {

            e.preventDefault();

            handleFormData(e.target);

            return;
        }
    }

    return (
        <FormContext.Provider>
            <form method="post" onSubmit={handleSubmit} {...allProps}>
                {children}
            </form>
        </FormContext.Provider>
    )
} 