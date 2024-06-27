import Form from "../../../components/form";
import PromptFormInput from "../../../components/promptFormInput";

export default function JobSectionFormm({ delegator }) {

    return (
        <Form delegate={delegator}>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput label="Nghề nghiệp" name="job" textArea={true} />
                </div>
                <div className="col-md-6">
                    <PromptFormInput label="Nơi Làm Việc" name="jobPlace" textArea={true} />
                </div>
            </div>
        </Form>
    )
}