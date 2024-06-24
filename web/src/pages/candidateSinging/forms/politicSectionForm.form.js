import Form from "../../../components/form";
import PromptFormInput from "../../../components/promptFormInput";

export default function PoliticSectionForm({ delegator }) {

    return (
        <Form delegate={delegator}>
            <div className="row">
                <div className="col-md-6">
                    <PromptFormInput label="Ngày Vào Đoàn TNCS Hồ Chí Minh" inputName="unionJoinDate" type="text" />
                </div>
                <div className="col-md-6">
                    <PromptFormInput label="Ngày Vào Đảng" inputName="partyJoinDate" type="text" />
                </div>
            </div>
            <br />
            <div className="row">
                <h4>Quá trình </h4>
                <br />
                <PromptFormInput label="Chi tiết" inputName="historyDetail" textArea={true} />
            </div>
        </Form>
    )
}