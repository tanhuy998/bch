import PaginationTable from "../components/paginationTable";

export default function ({usecase}) {

    return (
        <>
            <div class="page-title">

                <h3>
                    Campaign List
                    <a href="roles.html" class="btn btn-sm btn-outline-primary float-end"><i class="fas fa-plus-circle"></i> Thêm mới</a>
                    {/* <a href="users.html" class="btn btn-sm btn-outline-info float-end me-1"><i class="fas fa-angle-left"></i> <span class="btn-header">Return</span></a> */}
                </h3>
            </div>
            <div class="row">
                <div class="col-md-12 col-lg-12">
                    <div class="card">
                        <div class="card-header">Basic DataTables Table</div>
                        <div class="card-body">
                            <p class="card-title"></p>
                            <PaginationTable idField={"uuid"} endpoint={usecase} exposedFields={['title', 'issueTime', 'expire']} headers={['Campaign Name', 'Issue Time', 'Expires']} title="Campaigns" />
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}