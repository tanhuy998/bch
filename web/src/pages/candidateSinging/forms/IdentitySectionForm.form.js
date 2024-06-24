import Form from "../../../components/form";
import FormInput from "../../../components/formInput";

export default function IdentittySectionForm({ name }) {

    return (
        <Form name={name}>
            <div className="row">
                <div className="col">
                    <label >Họ Và Tên</label>
                    <FormInput className="form-control" name="name" />
                </div>
                <div className="col"></div>
            </div>
            <div class="row g-2">
                <div class="mb-3 col-md-6">
                    <label for="email" class="form-label">Ngày Sinh</label>
                    <FormInput type="date" className="form-control" name="dateOfBirth" placeholder="Email" required="" />
                    <small class="form-text text-muted">Enter a valid email address.</small>
                    <div class="valid-feedback">Looks good!</div>
                    <div class="invalid-feedback">Please enter your email address.</div>
                </div>
                <div class="mb-3 col-md-6">
                    <label for="password" class="form-label">Giới Tính</label>
                    <FormInput type="select" className="form-control" name="male" required="" />
                    <small class="form-text text-muted">Your password must be 8-20 characters long, contain letters and numbers only.</small>
                    <div class="valid-feedback">Looks good!</div>
                    <div class="invcurrentTab, setCurrentTabalid-feedback">Please enter your password.</div>
                </div>
            </div>
            <div className="row g-2">
                <div className="col-6">
                    <label >Số Căn Cước Công DÂn</label>
                    <FormInput className="form-control" name="name" />
                </div>
            </div>
            <div class="row g-2">
                <div class="mb-3 col-md-6">
                    <label for="city" class="form-label">Nơi Đăng Ký Khai Sinh</label>
                    <input type="text" class="form-control" name="city" placeholder="City" required="" />
                    <div class="valid-feedback">Looks good!</div>
                    <div class="invalid-feedback">Please enter your City.</div>
                </div>
                <div class="mb-3 col-md-4">
                    <label for="state" class="form-label">Quê Quán</label>
                    <select name="state" class="form-select" required="">
                        <option value="" selected="">Choose...</option>
                        <option value="1">New York</option>
                        <option value="2">Los Angeles</option>
                    </select>
                    <div class="valid-feedback">Looks good!</div>
                    <div class="invalid-feedback">Please select a State.</div>
                </div>
                <div class="mb-3 col-md-2">
                    <label for="zip" class="form-label">Dân Tộc</label>
                    <input type="text" class="form-control" name="zip" placeholder="00000" required="" />
                    <div class="valid-feedback">Looks good!</div>
                    <div class="invalid-feedback">Please enter a Zip code.</div>
                </div>
            </div>
            <div className="line"><br /></div>
            <br />
            <div className="row">
                <div className="col">
                    <label >Nơi Thường Trú</label>
                    <FormInput className="form-control" name="name" />
                </div>
                <div className="col"></div>
            </div>
            <div className="row">
                <div className="col">
                    <label >Nơi Ở Hiện Tại</label>
                    <FormInput className="form-control" name="name" />
                </div>
                <div className="col"></div>
            </div>
        </Form>
    )
}