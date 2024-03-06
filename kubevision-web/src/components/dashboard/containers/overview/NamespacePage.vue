<template>
    <v-row>
        <v-col cols="12">
            <v-data-table density="compact" :headers="table.headers" :loading="table.refreshing" :items="table.items"
                item-value="name" :items-per-page="table.itemsPerPage" :search="table.search" show-select
                v-model="table.selected">
                <template v-slot:top>
                    <v-row>
                        <v-col cols="6">
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

                <template v-slot:[`item.labels`]="{ item }">
                    <v-chip size="small" label v-bind:key="key" v-for="value, key in item.labels" class="mr-2">{{
                key }}={{ value }}</v-chip>
                </template>
            </v-data-table>
        </v-col>
    </v-row>
</template>

<script setup>
import { ref } from 'vue';

import { NamespaceTable } from '@/assets/app/tables';

var table = ref(new NamespaceTable())

table.value.refresh()

</script>