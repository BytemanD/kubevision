<template>
    <v-dialog width="80%" scrollable v-model="display">
        <v-card>
            <v-card-title class="bg-primary">{{ title }}</v-card-title>
            <v-card-text>
                <div style="overflow-y: scroll; font-size: small;">
                    <tr>
                        <td class="bg-blue-grey text-white px-1 text-right">
                            <span v-for="index of lines" v-bind:key="index">{{ index }} <br /></span>
                        </td>
                        <td style="width: 100%;">
                            <pre id="describeResource"><code ref="codeEle" class="hljs"></code></pre>
                        </td>
                    </tr>
                </div>
            </v-card-text>
        </v-card>
    </v-dialog>
</template>

<script setup>
import { ref, watch, onUnmounted, onUpdated } from 'vue';

import 'highlight.js/styles/atom-one-light.css'
import hljs from 'highlight.js';

const props = defineProps({
    show: { type: Boolean, required: true, },
    title: { type: String, required: true, },
    code: { type: String, required: true, },
})
var display = ref(false)
var lines = ref(0)
const emits = defineEmits(['update:show'])

var wrapRegex = new RegExp('\\n', 'g');
var codeNode = null;

const codeEle = ref()

function getCodeLines() {
    let lines = 0;
    Array.from(props.code.matchAll(wrapRegex), () => lines++);
    return lines;
}
function updateCode() {
    if (!codeNode) {
        let preNode = document.getElementById('describeResource');
        codeNode = preNode.childNodes[0];
    }
    lines.value = getCodeLines()
    let hljsCode = hljs.highlightAuto(props.code, ['yaml']);
    codeNode.innerHTML = hljsCode.value;
}

watch(() => (props.show), (newValue, oldValue) => {
    display.value = newValue;
})
watch(() => (display.value), (newValue, oldValue) => {
    emits('update:show', newValue)
})
onUnmounted(() => {
    emits('update:show', display.value)
})

onUpdated(() => {
    console.log('on updated')
    if (display.value) {
        let preNode = document.getElementById('describeResource');
        if (preNode) {
            codeNode = preNode.childNodes[0];
            if (display.value) {
                updateCode()
            }
        } else {
            console.error('<pre> not found')
        }
    }
})

</script>
