import Form from "../../../components/form";
import FormInput from "../../../components/formInput";
import PromptFormInput from "../../../components/promptFormInput";
import { validatePeopleName } from "../../../lib/validator";

export default function IdentittySectionForm({ name, delegator }) {

    return (
        <Form delegate={delegator} name={name}>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput label='Họ Và Tên' type='text' name="name" validate={validatePeopleName} invalidMessage="Tên chỉ chứa ký tự"/>
                </div>
                <div className="col"></div>
            </div>
            <div class="row g-2">
                <div class="mb-3 col-md-6">
                    <PromptFormInput label="Ngày Sinh" type="date" name="dateOfBirth"/>
                </div>
                <div class="mb-3 col-md-6">
                    <PromptFormInput label="Giới Tính" name="male"/>
                </div>
            </div>
            <div className="row g-2">
                <div className="mb-3 col-md-6">
                    <PromptFormInput label="Số Căn Cước Công Dân" name="idNumber" type="text"/>
                </div>
            </div>
            <div class="row g-2">
                <div class="mb-3 col-md-6">
                    <PromptFormInput label="Nơi Đăng Ký Khai Sinh" name="birthPlace" type="text"/>
                </div>
            </div>
            <div class="row g-2">
                <div class="mb-3 col-md-4">
                    <PromptFormInput label="Quê quán" name="birthPlace" type="text" />
                </div>
                <div class="mb-3 col-md-4">
                    <PromptFormInput label="Tôn Giáo" name="placeOfOrigin" type="text" />
                </div>
                <div class="mb-3 col-md-4">
                    <PromptFormInput label="Dân Tộc" name="ethnicity" type="text" />
                </div>
            </div>
            <div className="line"><br /></div>
            <br />
            <div className="row">
                <div className="mb-3 col-md-6">
                    <PromptFormInput label="Nơi Thường Trú" name="permanentResident" type="text"/>
                </div>
                <div className="mb-3 col-md-6">
                    <PromptFormInput label="Nơi Ở Hiện Tại" name="currentResident" type="text" />
                </div>
                {/* <div className="col"></div> */}
            </div>
            <div className="row">
                
                <div className="col"></div>
            </div>
        </Form>
    )
}