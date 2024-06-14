import { Link } from "react-router-dom";


export default function SideBar({ sideBarActive }) {

    return (
        <nav id="sidebar" class={sideBarActive}>
            <div class="sidebar-header">
                <img src="assets/img/bootstraper-logo.png" alt="bootraper logo" class="app-logo" />
            </div>
            <ul class="list-unstyled components text-secondary">
                <li>
                    <Link to="/admin"><i class="fas fa-home"></i>Trang Chủ</Link>
                </li>
                <li>
                    <Link to="#"><i class="fas fa-home"></i>Tuổi 17</Link>
                </li>
                <li>
                    <Link to="/admin/campaigns"><i class="fas fa-layer-group"></i>Campaigns</Link>
                </li>
            </ul>
        </nav>
    )
}