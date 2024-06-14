export default class AuthenticateSessionExpiresError extends Error {

    constructor() {

        super("session expires, redirect to re-login");
    }
}