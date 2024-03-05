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
                    </v-row>
                </template>

                <template v-slot:[`item.data_nums`]="{ item }">
                    {{ Object.keys(item.data).length }}
                </template>
            </v-data-table>
        </v-col>
    </v-row>
</template>

<script setup>
import { ref } from 'vue';

import { ConfigMapTable } from '@/assets/app/tables';

var table = ref(new ConfigMapTable())

table.value.refresh()

</script>