import Form from "../../../components/form";
import PromptFormInput from "../../../components/promptFormInput";

export default function EducationSectionForm({ delegator }) {

    return (
        <Form delegate={delegator}>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput label="Trình Độ Học Vấn" name="highestGrade" type="text" />
                </div>
            </div>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput label="Trình Độ Chuyên Môn" name="expertise" type="text" />
                </div>
                <div className="col-md-6">
                    <PromptFormInput label="Chuyên Ngành Đào tạo" name="" type='text' />
                </div>
            </div>
        </Form>
    )
}
