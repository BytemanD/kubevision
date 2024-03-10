<template>
    <v-dialog width="80%" scrollable v-model="display">
        <v-card>
            <v-card-title>
                <v-row>
                    <v-col cols="4">
                        <v-select hide-details density="compact" :items="pod.containers" item-title="name"
                            v-model="data.container">
                            <template v-slot:prepend>容器</template>
                        </v-select>
                    </v-col>
                    <v-col>
                        <v-text-field hide-details density="compact" placeholder="请输入命令" v-model="data.command">
                            <template v-slot:prepend>命令</template>
                        </v-text-field>
                    </v-col>
                    <v-col cols="2" class="text-center my-auto">
                        <v-btn color="warning" :disabled="!data.command || data.command.length == 0"
                            @click="execute()" :loading="loading">执行</v-btn>
                    </v-col>
                </v-row>
            </v-card-title>
            <v-card-text style="height: 600px;" class="pt-4">
                <v-alert density="compact" color="red" v-if="data.error"> {{ data.error }}</v-alert>
                <pre class="bg-black" v-else >{{ data.content }} </pre>
            </v-card-text>
        </v-card>
    </v-dialog>
</template>

<script setup>
import { ref, watch, onUnmounted, onUpdated } from 'vue';

import API from '@/assets/app/api'

const props = defineProps({
    show: { type: Boolean, required: true, },
    pod: { type: Object, required: true, },
})
const emits = defineEmits(['update:show'])

var data = ref({
    container: null, command: 'ls -l', content: '',
    error: ''
})

var display = ref(false)
var loading = ref(false)

watch(() => (props.show), (newValue, oldValue) => {
    display.value = newValue;
})
watch(() => (display.value), (newValue, oldValue) => {
    emits('update:show', newValue)
})
onUnmounted(() => {
    emits('update:show', display.value)
})

async function execute() {
    loading.value = true
    try {
        data.value.content = ''
        data.value.error = ''
        let result = await API.pods.execute(props.pod.name, data.value.container, data.value.command)
        data.value.content = result.stdout || result.stderr
    } catch(e){
        console.error('execute failed: ', e)
        data.value.error = e
    } finally{
        loading.value = false
    }
}
onUpdated(() => {
    if (display.value) {
        if (props.pod.containers.length > 0) {
            data.value.container = props.pod.containers[0].name
        }
    }
})

</script>
