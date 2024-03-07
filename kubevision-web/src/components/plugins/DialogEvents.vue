<template>
    <v-dialog v-model="display" scrollable style="font-size: small;">
        <template v-slot:activator="{ props }">
            <v-btn v-bind="props" icon="mdi-bell" fab color="warning">
            </v-btn>
        </template>
        <v-card title="事件">
            <template v-slot:append>
                <v-btn variant="tonal" color="red" size="small" @click="table.refresh()" icon="mdi-refresh">
                </v-btn>
            </template>
            <v-card-text>
                <v-data-table density="compact" :headers="table.headers" :items="table.items" item-key="name"
                    items-per-page="10" :search="table.search" v-model="table.selected">
                    <template v-slot:[`item.involved_object`]="{ item }">
                        {{ item.involved_object && item.involved_object.kind }}/{{ item.involved_object.name }}
                    </template>
                    <template v-slot:[`item.type`]="{ item }">
                        <v-icon color="warning" v-if="item.type.toLowerCase() == 'warning'">mdi-alert-circle</v-icon>
                        <v-icon color="info" v-else-if="item.type.toLowerCase() == 'normal'">mdi-information</v-icon>
                        <v-chip x-small v-else>{{ item.type }}</v-chip>
                    </template>
                    <template v-slot:[`item.message`]="{ item }">
                        <v-tooltip top v-if="item.message.length > 70">
                            <template v-slot:activator="{ props }">
                                <span class="text-info" v-bind="props">{{ item.message.slice(0, 70) }} ...</span>
                            </template>
                            {{ item.message }}
                        </v-tooltip>
                        <span v-else>{{ item.message }}</span>
                    </template>
                </v-data-table>
            </v-card-text>
        </v-card>

    </v-dialog>
</template>

<script setup>

import { ref } from 'vue';

import { EventTable } from '@/assets/app/tables';

var display = ref(false);
var table = ref(new EventTable())

table.value.refresh()
</script>