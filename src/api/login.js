
import CONFIG from "../app.config.json"
import {invokeApi} from "./common" 


export const registerUser = (user) => {
    return invokeApi(CONFIG.BASEURL+"/user/register", "PUT", user)
}

export const loginUser = (user) => {
    return invokeApi(CONFIG.BASEURL+"/user/login", "PUT", user)
}