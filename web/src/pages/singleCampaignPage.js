import { Link } from "react-router-dom";
import PaginationTable from "../components/paginationTable";

export default function SingleCampaignPage({title}) {

    return (
        <>
            <div class="row">
                <div class="col-md-12">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="card">
                                <h5 class="card-header">Campaign Detail</h5>
                                <div class="card-body">
                                    <h5 class="card-title">Some Title</h5>
                                    <p class="card-text">This is some text within a card body...</p>
                                    <Link to="#" class="btn btn-primary">Chỉnh sửa</Link>
                                    <></>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <PaginationTable title="Canidates"/>
            </div>
        </>
    )
}