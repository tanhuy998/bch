import AuthEndpoint from "../../backend/autEndpoint"
import CRUDEndpoint from "../../backend/crudEndpoint"

export default class TenantMainPageUseCase extends CRUDEndpoint {

    constructor() {

        super({
            uri: '/auth/man',
        })
    }

    async getTenantUsers() {


    }

    async getTenantCommandGroup() {


    }

    async createUsers() {


    }

    async deleteUser() {


    }
}