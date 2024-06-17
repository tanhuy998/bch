export default class NewCampaignResponsePresenter {

    #endpointResponse;

    get message() {

        return this.#endpointResponse.message;
    }

    get createdUUID() {

        return this.#endpointResponse.data.createdUUID;
    }

    constructor(endpointResponse) {

        if (typeof endpointResponse !== 'object') {

            throw new Error('invalid endpoint response');
        }
        
        this.#endpointResponse = endpointResponse;
    }
}