import Form from "../../../components/form";
import PromptFormInput from "../../../components/promptFormInput";

export default function EducationSectionForm({ delegator }) {

    return (
        <Form delegate={delegator}>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput label="Trình Độ Học Vấn" inputName="highestGrade" type="text" />
                </div>
            </div>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput label="Trình Độ Chuyên Môn" inputName="expertise" type="text" />
                </div>
                <div className="col-md-6">
                    <PromptFormInput label="Chuyên Ngành Đào tạo" inputName="" type='text' />
                </div>
            </div>
        </Form>
    )
}
