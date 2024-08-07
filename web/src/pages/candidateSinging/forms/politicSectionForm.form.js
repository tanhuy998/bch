import Form from "../../../components/form";
import PromptFormInput from "../../../components/promptFormInput";

export default function PoliticSectionForm({ delegator }) {

    return (
        <Form delegate={delegator}>
            <div className="row">
                <div className="col-md-4">
                    <PromptFormInput label="Ngày Vào Đoàn TNCS Hồ Chí Minh" name="unionJoinDate" type="date" />
                </div>
                <div className="col-md-4">
                    <PromptFormInput label="Ngày Vào Đảng" name="partyJoinDate" type="date" />
                </div>
            </div>
            <br />
            {/* <div className="row">
                <h4>Quá trình </h4>
                <br />
                <PromptFormInput label="Chi tiết" name="historyDetail" textArea={true} />
            </div> */}
        </Form>
    )
}