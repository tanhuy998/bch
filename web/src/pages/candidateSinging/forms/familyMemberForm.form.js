import BinaryCheckBoxFormInput from "../../../components/binaryCheckBoxFormInput";
import Form from "../../../components/form";
import PromptFormInput from "../../../components/promptFormInput";
import { validateFormalName } from "../../../lib/validator";

export default function FamilyMemberForm({ delegator, who, disableDeadInput }) {

    return (
        <Form delegate={delegator}>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput
                        label={`Họ Tên ${who}`}
                        name="name"
                        type="text"
                        invalidMessage="Tên chỉ được phép chứa ký tự"
                        validate={validateFormalName}
                    />
                </div>
                <div className="col-md-3">
                    <PromptFormInput
                        label="Năm Sinh"
                        type="date"
                        name="dateOfBirth"
                    />
                    {
                        !disableDeadInput &&
                        <>
                            <div className="form-check form-switch">
                                <BinaryCheckBoxFormInput className="form-check-input" name="dead" />
                                <label for="dead">đã chết?</label>
                            </div>
                        </>
                    }
                </div>
            </div>
            <br />
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput
                        label="Nghề Nghiệp"
                        type="text"
                        name="job"
                    />
                </div>
            </div>
        </Form>
    )
}