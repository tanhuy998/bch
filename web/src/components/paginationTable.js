import { useEffect, useState } from "react";
import PaginationPanel from "./paginationPanel";

export default function PaginationTable({headers, title}) {

    const [navigation, setNavigation] = useState()

    /**@type {Array<string>} */
    headers = Array.isArray(headers) ? headers : [];
    title = typeof title === 'string' ? title : "";

    useEffect(state => {

        
    }, [navigation])

    return (
        <div class="container">
            <div class="page-title">
                <h3>{title}</h3>
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
                                        {/* <th>ID</th>
                                        <th>Name</th>
                                        <th>Salary</th>
                                        <th>Country</th>
                                        <th>City</th> */}
                                        {
                                            headers.map((header) => {
                                                <th>header</th>
                                            })
                                        }
                                    </tr>campaigns
                                </thead>
                                <tbody>
                                    <tr>
                                        <td>1</td>
                                        <td>Dakota Rice</td>
                                        <td>$36,738</td>
                                        <td>United States</td>
                                        <td>Oud-Turnhout</td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                        <PaginationPanel setNavigation={setNavigation}/>
                    </div>
                </div>
            </div>
        </div>
    )
}