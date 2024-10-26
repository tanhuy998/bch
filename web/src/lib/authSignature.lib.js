import { jwtDecode } from "jwt-decode";

const ACCESS_TOKEN_KEY = "at";
const USER_INFO_KEY = "ui";
const ACCESS_TOKEN_EXP_KEY = "at-exp";
let userInfo;

let AT;

/**
 * @name getAccessToken
 * @function
 * @returns {string|null}
 */
export function getAccessToken() {

    return localStorage.getItem(ACCESS_TOKEN_KEY);
}

/**
 * @name setAccessToken
 * @function
 * @param {string} accessToken 
 * 
 * @throws {Error}
 */
export function setAccessToken(accessToken) {

    const payload = jwtDecode(accessToken, {
        header: false
    });
    
    const numericExp = payload.exp;
    console.log(numericExp, payload)

    if (!numericExp || typeof numericExp != "number") {

        throw new Error("invalid access token expiry");
    }

    localStorage.setItem(ACCESS_TOKEN_EXP_KEY, numericExp*1000);
    localStorage.setItem(ACCESS_TOKEN_KEY, accessToken);
}

export function removeAccessToken() {

    localStorage.removeItem(ACCESS_TOKEN_KEY);
    localStorage.removeItem(ACCESS_TOKEN_EXP_KEY);
}

/**
 * 
 * @returns {number|undefined}
 */
export function getAccessTokenExp() {

    const numeric = localStorage.getItem(ACCESS_TOKEN_EXP_KEY);

    if (!numeric) {

        return;
    }

    return Number(numeric);
}

/**
     *  @typedef {Object} UserInfo
     *  @property {string} uuid
     *  @property {string} name
     *  @property {string} username
     *  @property {string} tenantUUID
     */

/**
 * @name getUserInfo
 * @function
 * @returns {UserInfo|undefined}
 */
export function getUserInfo() {

    return userInfo;
}

/**
 * @name setUserInfo
 * @function
 * @param {UserInfo} data 
 */
export function setUserInfo(data) {

    userInfo = data;
}