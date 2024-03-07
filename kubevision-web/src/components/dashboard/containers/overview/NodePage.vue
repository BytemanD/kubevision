<template>
    <v-row>
        <v-col cols="12">
            <v-data-table density="compact" :headers="table.headers" :loading="table.refreshing" :items="table.items"
                item-value="name" :items-per-page="table.itemsPerPage" :search="table.search" show-select
                v-model="table.selected" show-expand single-expand hover>
                <template v-slot:top>
                    <v-row>
                        <v-col cols="6"></v-col>
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
                    <v-icon color="green" v-if="item.ready == 'True'">mdi-check-circle</v-icon>
                    <v-icon color="red" v-else>mdi-close-circle</v-icon>
                </template>

                <template v-slot:[`item.labels`]="{ item }">
                    <v-chip size="small" label v-bind:key="key" v-for="value, key in item.labels" class="mr-2">
                        {{ key }}={{ value }}</v-chip>
                </template>

                <template v-slot:[`item.cpu`]="{ item }">{{ item.capacity.cpu }}</template>

                <template v-slot:[`item.memory`]="{ item }">{{ item.capacity.memory }}</template>

                <template v-slot:expanded-row="{ columns, item }">
                    <td></td>
                    <td :colspan="columns.length - 1">
                        <table>
                            <tr v-for="extendItem in table.extendItems" v-bind:key="extendItem.text">
                                <td><strong>{{ extendItem.title }}:</strong> </td>
                                <td>{{ item[extendItem.text] }}</td>
                            </tr>
                            <tr>
                                <td><strong>标签</strong> </td>
                                <td>
                                    <template v-for="value, key in item.labels">
                                        <v-chip size="small" label close class="my-1 mr-1"
                                            v-if="table.hideLabels.indexOf(key) < 0" v-bind:key="key"
                                            @click:close="table.deleteLabel(item, key)">
                                            {{ key }}={{ value }}
                                        </v-chip>
                                    </template>
                                </td>
                            </tr>
                        </table>
                    </td>

                </template>

            </v-data-table>
        </v-col>
    </v-row>
</template>

<script setup>
import { ref } from 'vue';

import { NodeTable } from '@/assets/app/tables';

var table = ref(new NodeTable())

table.value.refresh()

</script>