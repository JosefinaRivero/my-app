
import CONFIG from "../app.config.json"

export const getHotels = () => {
    return invokeApi(CONFIG.BASEURL+"/hotels")
}

export const invokeApi = async (path, method, data) => {

    const url = path;
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    let body = null;
    if (method !== "GET") {
        body = JSON.stringify(data);
    }
    const reqInfo = {
        method: method,
        body: body,
        headers: headers
    };
    const response = await fetch(url, reqInfo);

    if (response.ok) {
            try {
                return await response.json();
            }
            catch (ex) {
                console.error(ex);
                throw new Error("Null Object response, Status :" + response.status);
            }
    } else {
        if (response.status === 404) {
            throw new Error("Requested data not found");
        }
        if (response.status === 403) {
            throw new Error("Forbidden action, not authorized");
        }
        if (response.status === 401) {
            throw new Error("Unauthorized action, not authorized");
        }
        throw new Error("Error Status: " + response.status)
        
    }
}