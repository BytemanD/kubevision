<template>
    <v-row>
        <v-col cols="12">
            <v-data-table density="compact" :headers="table.headers" :loading="table.refreshing" :items="table.items"
                item-value="name" :items-per-page="table.itemsPerPage" :search="table.search" show-select
                v-model="table.selected">
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
                <template v-slot:[`item.node_selector`]="{ item }">
                    <v-chip size="x-small" class="mr-1 mb-1" v-for="(v, k) in item.node_selector || {}" v-bind:key="k">
                        {{ k }}={{ v }}</v-chip>
                </template>
            </v-data-table>
        </v-col>
    </v-row>
</template>

<script setup>
import { ref } from 'vue';

import { JobTable } from '@/assets/app/tables';

var table = ref(new JobTable())

table.value.refresh()

</script>