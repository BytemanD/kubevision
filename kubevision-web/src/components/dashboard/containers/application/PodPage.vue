<template>
    <v-row>
        <v-col cols="12">
            <v-data-table density="compact" :headers="table.headers" :loading="table.refreshing" :items="table.items"
                item-value="name" :items-per-page="table.itemsPerPage" :search="table.search" show-select
                v-model="table.selected" show-expand single-expand>
                <template v-slot:top>
                    <v-row>
                        <v-col cols="6"></v-col>
                        <v-col>
                            <v-text-field density="compact" hide-details v-model="table.search"
                                append-icon="mdi-magnify" placeholder="搜索"></v-text-field>
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
            {{ table.items[0] }}
        </v-col>
    </v-row>
</template>

<script setup>
import { ref } from 'vue';

import { PodTable } from '@/assets/app/tables';

var table = ref(new PodTable())
table.value.refresh()

</script>