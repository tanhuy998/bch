import Form from "../../../components/form";
import PromptFormInput from "../../../components/promptFormInput";

export default function FamilyMemberHistory({}) {

    return (
        <>
            <Form>
                <div className="row">
                    <div className="col-md-6">
                        <PromptFormInput label="Trước 1975" textArea={true}/>
                    </div>
                    <div className="col-md-6">
                        <PromptFormInput label="Sau 1975" textArea={true}/>
                    </div>
                </div>
            </Form>
        </>
    )
}