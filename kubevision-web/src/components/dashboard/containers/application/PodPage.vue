<template>
    <v-row>
        <v-col cols="12">
            <v-data-table density="compact" :headers="table.headers" :loading="table.refreshing" :items="table.items"
                item-value="name" :items-per-page="table.itemsPerPage" :search="table.search" show-select
                v-model="table.selected" show-expand single-expand>
                <template v-slot:top>
                    <v-row>
                        <v-col cols="6">
                            <v-toolbar density="compact">
                                <delete-confirm-dialog :disabled="table.selected.length == 0" title="确定删除Pod?"
                                    @click:comfirm="deleteSelected()" :items="table.selected" />
                            </v-toolbar>
                        </v-col>
                        <v-col>
                            <v-text-field density="compact" hide-details v-model="table.search"
                                append-icon="mdi-magnify" placeholder="搜索"></v-text-field>
                        </v-col>
                        <v-col cols="2">
                            <v-btn variant="text" icon="mdi-refresh" color="info" v-on:click="table.refresh()"></v-btn>
                        </v-col>
                    </v-row>
                </template>

                <template v-slot:[`item.ready`]="{ item }">
                    / {{ item.containers.length }}
                    <!-- {{ Utils.getPodReadyNum(item) }} /{{ item.container_statuses.length }} -->
                </template>

                <template v-slot:[`item.containers`]="{ item }">
                    <v-chip size="x-small" class="mr-1 mb-1" v-for="c in item.containers" v-bind:key="c.name">
                        {{ c.name }}</v-chip>
                </template>
                <template v-slot:[`item.actions`]="{ item }">
                    <v-menu offset-y>
                        <template v-slot:activator="{ props }">
                            <v-btn color="purple" size="small" variant="text" v-bind="props" icon="mdi-dots-vertical"></v-btn>
                        </template>
                        <v-list density="compact">
                            <v-list-item @click="openDialogPod(item)">
                                <v-list-item-title>查看</v-list-item-title>
                            </v-list-item>
                            <v-list-item @click="describeResource(item)">
                                <v-list-item-title>描述</v-list-item-title>
                            </v-list-item>
                            <v-list-item @click="openDialogPodLogs(item)" :disabled="item.status != 'Running'">
                                <v-list-item-title>日志</v-list-item-title>
                            </v-list-item>
                            <v-list-item @click="openDialogPodExecute(item)" :disabled="item.status != 'Running'">
                                <v-list-item-title>执行命令</v-list-item-title>
                            </v-list-item>
                        </v-list>
                    </v-menu>
                </template>
                <template v-slot:expanded-row="{ columns, item }">
                    <td></td>
                    <td :colspan="columns.length - 2">
                        <table>
                            <tr v-for="extendItem in table.extendItems" v-bind:key="extendItem.text">
                                <td><strong>{{ extendItem.title }}:</strong> </td>
                                <td>{{ item[extendItem.key] }}</td>
                            </tr>
                        </table>
                        <br>
                    </td>
                </template>
            </v-data-table>
            <dialog-code :title="'Pod: ' + selected.name" :show="displayCode" :code="code"
                @update:show="(i) => { displayCode = i }" />
            <dialog-pod-execute :pod="selected" :show="displayPodExecute"
                @update:show="(i) => { displayPodExecute = i }" />
            <dialog-pod-logs :pod="selected" :show="displayPodLogs" @update:show="(i) => { displayPodLogs = i }" />
            <dialog-pod :pod="selected" :show="displayPodDialog" @update:show="(i) => { displayPodDialog = i }"/>
        </v-col>
    </v-row>
</template>

<script setup>
import { ref } from 'vue';
import notify from '@/assets/app/notify';
import API from '@/assets/app/api'
import { PodTable } from '@/assets/app/tables';

import DialogCode from '@/components/plugins/DialogCode.vue';
import DialogPodExecute from '@/components/plugins/DialogPodExecute.vue';
import DialogPodLogs from '@/components/plugins/DialogPodLogs.vue';
import DeleteConfirmDialog from '@/components/plugins/dialogs/DeleteComfirmDialog.vue';
import DialogPod from '@/components/plugins/DialogPod.vue';

var table = ref(new PodTable())

var selected = ref({})
var code = ref('')
var displayCode = ref(false)
var displayPodExecute = ref(false)
var displayPodLogs = ref(false)
var displayPodDialog = ref(false)

table.value.refresh()

async function openDialogPodExecute(item) {
    selected.value = item;
    displayPodExecute.value = true;
}
async function openDialogPodLogs(item) {
    selected.value = item;
    displayPodLogs.value = true;
}
async function openDialogPod(item) {
    selected.value = item;
    displayPodDialog.value = true;
}
async function describeResource(item) {
    selected.value = item;
    displayCode.value = true;
    code.value = await API.pods.describe(item.name)
}

async function deleteSelected(){
    for (let item of table.value.selected) {
        await API.pods.delete(item)
        notify.success(`删除pod ${item} 成功`)
    }
    table.value.resetSelected()
}

</script>