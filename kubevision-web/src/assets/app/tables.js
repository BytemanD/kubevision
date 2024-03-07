import { Utils } from './lib.js'
import API from '@/assets/app/api'
import Notify from '@/assets/app/notify'

class DataTable {
    constructor(headers, api, bodyKey = null, name = '') {
        this.headers = headers || [];
        this.api = api;
        this.bodyKey = bodyKey;
        this.name = name;
        this.itemsPerPage = 20;
        this.search = '';
        this.items = [];
        this.statistics = {};
        this.selected = []
        this.extendItems = []
        this.newItemDialog = null;
        this.refreshing = false
        this.columns = this.headers.map((header) => { return header.key });
    }
    async openNewItemDialog() {
        if (this.newItemDialog) {
            this.newItemDialog.open();
        }
    }
    async createNewItem() {
        if (this.newItemDialog) {
            await this.newItemDialog.commit();
            this.refresh();
        }
    }
    async deleteSelected() {
        if (this.selected.length == 0) {
            return;
        }
        Notify.info(`开始删除`);
        let deleting = [];
        for (let i in this.selected) {
            let item = this.selected[i];
            try {
                await this.api.delete(item);
                deleting.push(item);
                this.watchDeleting(item)
            } catch {
                Notify.error(`删除 ${item} 失败`)
            }
        }
        this.resetSelected()
    }
    async watchDeleting(itemId) {
        do {
            try {
                let item = await (this.api.show(itemId))
                this.updateItem(item);
                Utils.sleep(2)
            } catch (e) {
                if (e.response.status == 404) {
                    console.error(e)
                    Notify.success(`${this.name} ${itemId} 已删除`)
                    this.removeItem(itemId)
                    break;
                }
            }
        } while (true)
    }
    resetSelected() {
        this.selected = [];
    }
    updateItem(newItem) {
        for (var i = 0; i < this.items.length; i++) {
            if (this.items[i].id != newItem.id) {
                continue;
            }
            for (var key in newItem) {
                if (this.items[i][key] == newItem[key]) {
                    continue
                }
                this.items[i][key] = newItem[key];
            }
            break
        }
    }
    removeItem(id) {
        let index = -1;
        for (let i in this.items) {
            if (this.items[i].id == id) {
                index = i
                break;
            }
        }
        if (index >= 0) {
            this.items.splice(index, 1)
        }
    }
    async refresh(filters = {}) {
        this.refreshing = true;
        let result = null
        try {
            if (this.api.detail) {
                result = await this.api.detail(filters);
            } else {
                result = await this.api.list(filters)
            }
        } catch (e) {
            console.error(e)
            Notify.error(`${this.name || '资源'} 查询失败`)
            return;
        } finally {
            this.refreshing = false;
        }
        this.items = this.bodyKey ? result[this.bodyKey] : result;
        return result;
    }
    getSelectedItems() {
        let items = [];
        for (let i in this.items) {
            if (this.selected.indexOf(this.items[i].id) < 0) {
                continue
            }
            items.push(this.items[i])
        }
        return items;
    }
}

export class NamespaceTable extends DataTable {
    constructor() {
        super([
            { title: '名字', key: 'name' },
            { title: '状态', key: 'status' },
            { title: '标签', key: 'labels' },
            { title: '创建时间', key: 'creation' },
            { title: '操作', key: 'actions' },
        ], API.namespaces, 'namespaces', '命名空间');
    }
}
export class NodeTable extends DataTable {
    constructor() {
        super([
            { title: '名字', key: 'name' },
            { title: 'Ready', key: 'ready' },
            { title: '内网IP', key: 'internal_ip' },
            { title: 'CPU', key: 'cpu' },
            { title: '内存', key: 'memory' },
            { title: '系统', key: 'os_image' },
            { title: '创建时间', key: 'creation' },
            { title: '操作', key: 'actions' },
        ], API.nodes, 'nodes', '节点');
        this.extendItems = [
            { title: '内核版本', key: 'kernel_version' },
            { title: '容器运行时版本', key: 'container_runtime_version' },
        ];
        this.hideLabels = [
            'beta.kubernetes.io/arch', 'beta.kubernetes.io/os',
            'kubernetes.io/arch', 'kubernetes.io/os', 'kubernetes.io/hostname',
            'node.kubernetes.io/exclude-from-external-load-balancers',
            'node-role.kubernetes.io/control-plane',
        ]
    }
}
export class DaemonsetTable extends DataTable {
    constructor() {
        super([{ title: '名字', key: 'name' },
        { title: '可用', key: 'number_available' },
        { title: '当前', key: 'current_number_scheduled' },
        { title: '准备', key: 'number_ready' },
        { title: '期望', key: 'desired_number_scheduled' },
        { title: '更新', key: 'updated_number_scheduled' },
        { title: 'node_selector', key: 'node_selector' },
        // { title: 'selector', key: 'selector' },
        // { title: 'containers', key: 'containers' },
        { title: '创建时间', key: 'creation' },
        { title: '操作', key: 'actions' },
        ], API.daemonsets, 'daemonsets', '服务守护进程');
        this.extendItems = [];
    }
}
export class DeploymentTable extends DataTable {
    constructor() {
        super([{ title: '名字', key: 'name' },
        { title: '准备', key: 'ready_replicas' },
        // { title: '副本', key: 'replicas' },
        { title: '更新', key: 'updated_replicas' },
        { title: '可用', key: 'available_replicas' },
        // { title: 'containers', key: 'containers' },
        { title: '创建时间', key: 'creation' },
        { title: '操作', key: 'actions' },
        ], API.deployments, 'deployments', '服务守护进程');
        this.extendItems = [
            // { title: 'images', key: 'images' },
        ];
    }
}
export class PodTable extends DataTable {
    constructor() {
        super([
            { title: '名字', key: 'name' },
            { title: '状态', key: 'status' },
            { title: '就绪', key: 'ready' },
            { title: '阶段', key: 'phase' },
            { title: '所在节点', key: 'node_name' },
            { title: 'IP地址', key: 'pod_id' },
            // { title: 'containers', key: 'containers' },
            { title: '创建时间', key: 'creation' },
            { title: '操作', key: 'actions' },
        ], API.pods, 'pods', '容器组');
        this.extendItems = [
            { title: 'labels', key: 'labels' },
            { title: 'node_selector', key: 'node_selector' },
            { title: 'containers', key: 'containers' },
        ];
        this.waiting = {};
    }
    updateWaiting(pod) {
        this.waiting[pod.name] = Utils.getPodWaiting(pod);
    }
}

export class ConfigMapTable extends DataTable {
    constructor() {
        super([
            { title: '名字', key: 'name' },
            { title: '数据个数', key: 'data_nums' },
            { title: '创建时间', key: 'creation' },
            { title: '操作', key: 'actions' },
        ], API.configmaps, 'configmaps', '配置字典');
        this.extendItems = [
            { title: '数据', key: 'data_list' },
        ]
    }
}
export class SecretTable extends DataTable {
    constructor() {
        super([
            { title: '名字', key: 'name' },
            { title: '数据个数', key: 'data_nums' },
            { title: '创建时间', key: 'creation' },
            { title: '操作', key: 'actions' },
        ], API.secrets, 'secrets', '加密数据');
        this.extendItems = [
            { title: '数据', key: 'data_list' },
        ]
    }
}
export class ServiceTable extends DataTable {
    constructor() {
        super([
            { title: '名字', key: 'name' },
            { title: '类型', key: 'type' },
            { title: 'IP', key: 'cluster_ip' },
            //    { title: 'IP列表', key: 'cluster_i_ps' },
            { title: 'IP栈', key: 'ip_families' },
            { title: '端口', key: 'ports' },
            { title: '创建时间', key: 'creation' },
            { title: '操作', key: 'actions' },
        ], API.services, 'services', '服务');
        this.extendItems = [
            { title: '数据', key: 'data_list' },
        ]
    }
}
export class CronjobTable extends DataTable {
    constructor() {
        super([
            { title: '名字', key: 'name' },
            { title: '计划', key: 'schedule' },
            { title: 'node_selector', key: 'node_selector' },
            // { title: 'containers', key: 'containers' },
            { title: '创建时间', key: 'creation' },
        ], API.cronjobs, 'cronjobs', '定时任务');
        this.extendItems = [
            { title: '数据', key: 'data_list' },
        ]
    }
}
export class JobTable extends DataTable {
    constructor() {
        super([
            { title: '名字', key: 'name' },
            { title: 'node_selector', key: 'node_selector' },
            { title: 'containers', key: 'containers' },
            { title: '创建时间', key: 'creation' },
        ], API.jobs, 'jobs', '任务');
        this.extendItems = [
            { title: '数据', key: 'data_list' },
        ]
    }
}
export class StatefulsetsTable extends DataTable {
    constructor() {
        super([
            { title: '名字', key: 'name' },
            { title: 'node_selector', key: 'node_selector' },
            { title: 'selector', key: 'selector' },
            { title: 'containers', key: 'containers' },
            { title: '创建时间', key: 'creation' },
        ], API.events, 'events', '事件');
    }
}
export class EventTable extends DataTable {
    constructor() {
        super([
            // { title: '名字', key: 'name' },
            { title: '时间', key: 'event_time' },
            { title: '类型', key: 'type' },
            // { title: '操作', key: 'action' },
            { title: '对象', key: 'involved_object' },
            { title: '重复次数', key: 'count' },
            { title: '原因', key: 'reason' },
            { title: '消息', key: 'message' },
            // { title: '来源', key: 'source' },
        ], API.events, 'events', '事件');
    }
}

export default DataTable;
