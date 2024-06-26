import Form from "../../../components/form";
import PromptFormInput from "../../../components/promptFormInput";

export default function FamilyMemberForm({delegator, who, disableDeadInput}) {

    return (
        <Form delegate={delegator}>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput label={`Họ Tên ${who}`} inputName="name" type="text" invalidMessage="Tên chỉ được phép chứa ký tự"/>
                </div>
                {!disableDeadInput && <div className="col-md-6">
                    <PromptFormInput label="Tình Trạng: " inputName="dead" type="text" />
                </div>}
            </div>
            <br />
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput label="Năm Sinh" type="date" inputName="dateOfBirth" />
                </div>
                <div className="col-md-6">
                    <PromptFormInput label="Nghề Nghiệp" type="text" inputName="job" />
                </div>
            </div>
        </Form>
    )
}