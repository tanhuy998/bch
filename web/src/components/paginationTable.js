import { useEffect, useState, Component, memo } from "react";
import PaginationPanel from "./paginationPanel";
import TableRow from "./tableRow";

/**
 * 
 * @param {Function} navigator 
 * @returns {Function}
 */
function dispatchNavigate(navigator) {

    return
}

function fetchData(setDataFunc, endpoint, pageCursor, pageLimit, isPrevDir) {

    return async function() {


    }
}

class PaginationTable extends Component  {

    #headers;
    #title;
    /**@type {HttpEndpoint} */
    #endpoint;
    #pagination;

    #paginationNav;

    #tableData = [];

    constructor(props) {

        super(props)

        this.#title = props.title;
        this.#headers = (Array.isArray(props.headers) ? props.headers : []).map(header => <th>{header}</th>)
        this.#endpoint = props.endpoint;
        //this.#pagination = props.pagination;

        this.fetchData = this.fetchData.bind(this);
    }

    async componentDidMount() {

        //await this.fetchData()
    }

    async componentDidUpdate(prevProps, prevState) {

        await this.fetchData()        
    }

    async fetchData() {

        try {

            console.log("table fetch");

            const fetchedData = await this.#endpoint.fetch();
        }
        catch (e) {
            console.log("fetch table err", e)
            this.#tableData = e   
        }
    }

    #prepareTableRows() {

        return this.#tableData.map((row) => <TableRow dataObject={row} />)        
    }

    render() {

        return (
            <>
                <div class="page-title">
                   
                    <h3>{this.#title}
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
                                <table class="table table-hover" id="dataTables-example" width="100%">
                                    <thead>
                                        <tr>
                                            {this.#headers}
                                        </tr>
                                    </thead>
                                    <tbody>
                                        {this.#prepareTableRows()}
                                    </tbody>
                                </table>
                            </div>
                            {typeof this.#paginationNav === 'object' && <PaginationPanel setNavigation={this.#paginationNav} />}
                        </div>
                    </div>
                </div>
            </>
        )
    }
}

const MemoPaginationTable = memo(PaginationTable);

export default MemoPaginationTable;