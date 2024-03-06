import axios from 'axios';

import notify from './notify.js';
import { LOG, Utils, MESSAGE } from "./lib.js";
import { RegExp } from 'core-js';

async function waitDeletedByList(api, bodyKey, item) {
    let items = [];
    do {
        console.debug(new Date().toLocaleString(), `wait ${item.name || item.id} to be deleted`);
        items = (await api.list({ id: item.id }))[bodyKey];
        if (items.length != 0) {
            await Utils.sleep(5);
        }
    } while (items.length != 0);
}
async function waitDeletedByGet(api, bodyKey, item) {
    let itemBody = {};
    let deleted = false;
    do {
        try {
            console.debug(new Date().toLocaleString(), `wait ${item.name || item.id} to be deleted`);
            itemBody = (await api.get(item.id))[bodyKey];
            await Utils.sleep(3);
        } catch (error) {
            MESSAGE.success(`${item.id || item.name} 删除成功`);
            deleted = true;
        }
    } while (!deleted);
    return itemBody
}

class Restfulclient {
    constructor(baseUrl, async = false) {
        this.async = async;
        this.baseUrl = baseUrl;
    }
    getHeaders() {
        return null;
    }
    _get_url(...args) {
        var urls = [this.baseUrl]
        args.map(x => { if (!!x) { urls.push(x) } })
        return urls.join('/')
    }
    _parseToQueryString(filters) {
        if (!filters) { return '' }
        let queryParams = [];
        for (var key in filters) {
            if (Array.isArray(filters[key])) {
                LOG.debug(`filters: ${filters[key]}`)
                filters[key].forEach(value => {
                    queryParams.push(`${key}=${value}`)
                })
            } else {
                queryParams.push(`${key}=${filters[key]}`)
            }
        }
        return queryParams.join('&');
    }
    async waitDeleted() {

    }
    _getErrorMsg(response) {
        let errorData = response ? response.data : {};
        if (errorData.badRequest && errorData.badRequest.message) {
            return errorData.badRequest.message
        } else {
            return JSON.stringify(errorData)
        }
    }
    async get(url = null) {
        let reqUrl = this.baseUrl;
        if (url) {
            if (url.startsWith('/')) {
                reqUrl = url;
            } else {
                reqUrl = this._get_url(url)
            }
        }
        let resp = await axios.get(reqUrl, { headers: this.getHeaders() });
        return resp.data
    }
    async delete(id) {
        let resp = await axios.delete(
            this._get_url(id),
            { headers: this.getHeaders() });
        return resp.data
    }
    async doPost(body, url = null) {
        try {
            let reqUrl = this.baseUrl;
            if (url) {
                if (url.startsWith('/')) {
                    reqUrl = url;
                } else {
                    reqUrl = this._get_url(url);
                }
            }
            let resp = await axios.post(
                reqUrl, body, { headers: this.getHeaders() });
            return resp
        } catch (e) {
            console.error(this._getErrorMsg(e.response));
            throw Error(this._getErrorMsg(e.response))
        }
    }
    async post(body, url = null) {
        let resp = await this.doPost(body, url)
        return resp.data
    }
    async put(id, body) {
        let resp = await axios.put(this._get_url(id), body, { headers: this.getHeaders() });
        return resp.data
    }
    async show(id, filters = null) {
        let url = filters ? `${id}?${this._parseToQueryString(filters)}` : id;
        let data = await this.get(`${url}`, { headers: this.getHeaders() });
        return data
    }
    async list(filters = {}) {
        let queryString = this._parseToQueryString(filters);
        let url = this._get_url()
        if (queryString) { url += `?${queryString}` }
        let resp = await axios.get(url, { headers: this.getHeaders() });
        return resp.data;
    }
    async patch(id, body, headers = {}) {
        let config = { headers: this.getHeaders() };
        for (let key in headers) {
            config.headers[key] = headers[key];
        }
        let resp = await axios.patch(this._get_url(id), body, config);
        return resp.data
    }
    async postAction(id, action, data) {
        let body = {};
        body[action] = data;
        return (await axios.post(
            this._get_url(id, action), body, { headers: this.getHeaders() })).data;
    }

    async listActive() {
        return (await this.list({ status: 'active' }))
    }
}

class System extends Restfulclient {
    constructor() {
        super('/login');
    }
    async login(username, password) {
        let auth = {
            username: username,
            password: password
        }
        return await this.doPost({ auth: auth })
    }
    async isLogin() {
        return await this.get('/login')
    }
    async logout() {
        await this.delete('login')
        notify.success('成功退出')
    }
}
class Cluster extends Restfulclient {
    constructor() { super('/cluster') }
}
class Namespaces extends Restfulclient {
    constructor() { super('/namespaces') }
}
class Nodes extends Restfulclient {
    constructor() { super('/nodes') }
}
class Deployments extends Restfulclient {
    constructor() { super('/deployments') }
}
class Daemonsets extends Restfulclient {
    constructor() { super('/daemonsets') }
}
class Pods extends Restfulclient {
    constructor() { super('/pods') }
}
class Services extends Restfulclient {
    constructor() { super('/services') }
}

class Jobs extends Restfulclient {
    constructor() { super('/jobs') }
}
class CronJobs extends Restfulclient {
    constructor() { super('/cronjobs') }
}
class Configmaps extends Restfulclient {
    constructor() { super('/configmaps') }
}
class Secrets extends Restfulclient {
    constructor() { super('/secrets') }
}
class Events extends Restfulclient {
    constructor() { super('/events') }
}
class Statefulsets extends Restfulclient {
    constructor() { super('/statefulsets') }
}

class AuthInfo extends Restfulclient {
    constructor() { super('/auth_info') }
    async get() {
        return (await this.list()).auth_info
    }
}
class Task extends Restfulclient {
    constructor() { super('/tasks') }
}
class Actions extends Restfulclient {
    constructor() { super('/actions') }
    async checkLastVersion() {
        let body = await this.post({ 'checkLastVersion': {} });
        return body.checkLastVersion
    }
}
class Version extends Restfulclient {
    constructor() { super('/version') }
    async get() {
        return (await super.get()).version;
    }
}


export class KubevisionAPI {
    constructor() {
        this.system = new System();
        this.cluster = new Cluster();

        this.namespaces = new Namespaces()
        this.nodes = new Nodes()
        this.daemonsets = new Daemonsets()
        this.deployments = new Deployments()
        this.pods = new Pods()
        this.services = new Services()
        this.jobs = new Jobs()
        this.cronjobs = new CronJobs()
        this.configmaps = new Configmaps()
        this.events = new Events()
        this.secrets = new Secrets()
        this.statefulsets = new Statefulsets()

        // this.task = new Task();

        this.actions = new Actions();
        this.version = new Version();
    }
}


Restfulclient.prototype.getHeaders = function () {
    let headers = {
        'X-Auth-Token': localStorage.getItem('X-Auth-Token'),
    };
    if (sessionStorage.getItem('namespace')) {
        headers['X-Namespace'] = sessionStorage.getItem('namespace');
    }
    return headers;
}

const API = new KubevisionAPI();

export default API;
