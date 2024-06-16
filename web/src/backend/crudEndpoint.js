import AuthEndpoint from "./autEndpoint";

export default class CRUDEndpoint extends AuthEndpoint {

    /**
     * 
     * @param {campaign_model_t} model 
     */
    async create(model) {

        return super.fetch(
            {
                method: 'POST',
                body: JSON.stringify({
                    data: model
                })
            }
        )
    }

    async read() {


    }

    async update(uuid, model) {

        this.#assertString(uuid);

        return super.fetch(
            {
                method: 'PATCH',
                body: JSON.stringify({
                    data: model
                })
            },
            undefined,
            `/${uuid}`
        )
    }

    async delete(uuid) {

       this.#assertString(uuid)

        return super.fetch(
            {
                method: 'DELETE',
            },
            undefined,
            `/${uuid}`
        )
    }

    fetch() {

        throw new Error('method fetch() is not allowed on crud endpoints');
    }

    #assertString(unknown) {

        if (typeof unknown !== 'string' || unknown === '') {

            throw new Error('uuid argument of CRUDEndpoint.delete() must be at least typeof string');
        }
    }
}