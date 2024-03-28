<template>
    <v-dialog v-model="display" scrollable width="80%">
        <v-card class="primary darken-2 white--text">
            <v-card-title class="bg-primary">容器组 {{ pod.name }}</v-card-title>
            <v-card-subtitle>创建时间: {{ pod.creation }}</v-card-subtitle>
            <v-card-text style="height: 600px;" class="pt-4">
                <v-row>
                    <v-col>
                        <h4>节点</h4>{{ pod.node_name }}
                    </v-col>
                    <v-col>
                        <h4>容器组IP</h4> {{ pod.pod_ip }}
                    </v-col>
                    <v-col>
                        <h4>阶段</h4> {{ pod.phase }}
                    </v-col>
                </v-row>
                <v-row>
                    <v-col cols="12" v-for="container in pod.containers" v-bind:key="container.name">
                        <v-sheet elevation="2" class="pa-4">
                            <v-row>
                                <v-col cols="4">
                                    <v-chip label color="primary darken-1">容器:</v-chip> {{ container.name }}
                                </v-col>
                                <v-col cols="4">
                                    <template v-for="port in container.ports" v-bind:key="port.containerPort">
                                        <pill-chip :name="port.protocol" :value="port.containerPort" small class="mt-1" />
                                    </template>
                                </v-col>
                                <v-col cols="4">
                                    <!-- <v-btn x-small class="pa-1" color="grey"><v-icon small>mdi-console-line</v-icon></v-btn>
                                        <v-btn x-small class="pa-1 ml-2" color="grey"><v-icon>mdi-cards-variant</v-icon></v-btn>
                                    <v-btn x-small class="pa-1 ml-2" color="primary"><v-icon>mdi-database</v-icon></v-btn>
                                    <v-btn x-small class="pa-1 ml-2" color="error"><v-icon>mdi-information</v-icon></v-btn> -->
                                </v-col>
                                <v-col cols="12">
                                    <span class="grey--text">镜像抓取策略: {{ container.image_pull_policy }}</span>
                                    <span class="grey--text ma-6">镜像: {{ container.image }}</span>
                                </v-col>
                            </v-row>
                        </v-sheet>
                    </v-col>
                </v-row>
            </v-card-text>
        </v-card>
    </v-dialog>
</template>

<script setup>
import { ref, watch, onUnmounted, onUpdated } from 'vue';

import PillChip from '@/components/plugins/chips/PillChip.vue';

const props = defineProps({
    show: { type: Boolean, required: true, },
    pod: { type: Object, required: true, },
})

const emits = defineEmits(['update:show'])
var display = ref(false)
var loading = ref(false)

watch(() => (props.show), (newValue, oldValue) => {
    display.value = newValue;
})
watch(() => (display.value), (newValue, oldValue) => {
    emits('update:show', newValue)
})

</script>
