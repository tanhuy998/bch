import TableRowManipulator from "../../components/lib/tableRowDataAction";

const UI_PATH = '/admin/campaign';

export default class CampaignListTableRowManipulator extends TableRowManipulator {

    #endpointHost;

    constructor(endpointHost) {

        super();

        this.#endpointHost = endpointHost;
    }

    generateRowDeletePath(uuid) {

        return `${UI_PATH}/${uuid}`;
    }

    generateRowDetailPath(uuid) {

        return `${UI_PATH}/${uuid}`;
    }

    generateRowModificationPath(uuid) {

        return `${this.#endpointHost}/campaigns/${uuid}`;
    }
}