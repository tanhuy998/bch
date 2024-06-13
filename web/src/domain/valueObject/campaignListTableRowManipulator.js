import TableRowManipulator from "../../components/lib/tableRowDataAction";

const UI_PATH = '/admin/campaign';

export default class CampaignListTableRowManipulator extends TableRowManipulator {

    #endpointHost;

    constructor(endpointHost) {

        super();

        this.#endpointHost = endpointHost;
    }

    generateRowModificationPath(uuid) {

        return `${UI_PATH}/edit/${uuid}`;
    }

    generateRowDetailPath(uuid) {

        return `${UI_PATH}/${uuid}`;
    }

    generateRowDeletePath(uuid) {

        return `${this.#endpointHost}/campaigns/${uuid}`;
    }
}